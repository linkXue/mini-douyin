// Code generated by hertz generator.

package first

import (
	"api.service/biz/model/api/douyin/core"
	first "api.service/biz/model/api/douyin/extra/first"
	"api.service/biz/rpc"
	basics "basics.rpc/kitex_gen/douyin/core"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	interaction "interaction.rpc/kitex_gen/douyin/extra/first"
	society "society.rpc/kitex_gen/douyin/extra/second"
	"strings"
)

// FavoriteAction .
// @router /douyin/favorite/action [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req first.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	resp := new(first.DouyinFavoriteActionResponse)

	videoId := req.VideoId
	actionType := req.ActionType
	myId := c.GetInt64("myId")
	if videoId == 0 || myId == 0 {
		hlog.Infof("videoId or myId is null")
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "videoId or myId is null"))
		return
	}
	_, err = rpc.BasicsService.GetUserInfoById(ctx, &basics.GetUserRequest{UserId: myId})
	if err != nil {
		hlog.Infof("BasicsService failed err:%v\n", err)
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	_, err = rpc.BasicsService.GetVideoInfoById(ctx, &basics.GetVideoByIdRequest{VideoId: videoId})
	if err != nil {
		hlog.Infof("BasicsService failed err:%v\n", err)
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	if actionType == 1 {
		_, err := rpc.InteractionService.AddVideoFavorite(ctx, &interaction.AddVideoFavoriteRequest{VideoId: videoId, UserId: myId})
		if err != nil {
			hlog.Infof("InteractionService failed err:%v\n", err)
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
			return
		}
	} else if actionType == 2 {

		_, err := rpc.InteractionService.CancelVideoFavorite(ctx, &interaction.CancelVideoFavoriteRequest{VideoId: videoId, UserId: myId})
		if err != nil {
			hlog.Infof("InteractionService failed err:%v\n", err)
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
			return
		}
	} else {
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "action type valied"))
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = new(string)
	*resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req first.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	resp := new(first.DouyinFavoriteListResponse)

	userId := req.UserId
	myId := c.GetInt64("myId")
	if userId != myId {
		resp.StatusCode = 400
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "No permissions"
		resp.VideoList = make([]*first.Video, 0, 0)
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	if userId == 0 {
		hlog.Infof("userId is null")
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "userId is null"))
		return
	}
	favoriteListResponse, err := rpc.InteractionService.GetFavoriteList(ctx, &interaction.GetFavoriteListRequest{UserId: userId})
	if err != nil {
		hlog.Infof("InteractionService failed err:%v\n", err)
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	videoList := favoriteListResponse.VideoList
	videos := make([]*first.Video, 0, len(videoList))
	for _, video := range videoList {
		if myId == 0 || video.Author.Id == 0 {
			hlog.Infof("myId or userId is null")
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "myId or userId is null"))
			return
		}
		societyInfoResponse, err := rpc.SocietyService.SocietyInfo(ctx, &society.SocietyInfoRequest{MyId: myId, UserId: video.Author.Id})
		if err != nil {
			hlog.Infof("SocietyService failed err:%v\n", err)
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
			return
		}
		videos = append(videos, &first.Video{
			Id: video.Id,
			Author: &first.User{
				Id:            video.Author.Id,
				Name:          video.Author.Name,
				FollowCount:   &societyInfoResponse.FollowCount,
				FollowerCount: &societyInfoResponse.FollowerCount,
				IsFollow:      societyInfoResponse.IsFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}
	resp.StatusCode = 0
	resp.StatusMsg = new(string)
	*resp.StatusMsg = "success"
	resp.VideoList = videos
	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req first.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	myId := c.GetInt64("myId")
	actionType := req.ActionType
	videoId := req.VideoId
	commentText := req.CommentText
	commentId := req.CommentId
	resp := new(first.DouyinCommentActionResponse)
	_, err = rpc.BasicsService.GetUserInfoById(ctx, &basics.GetUserRequest{UserId: myId})
	if err != nil {
		hlog.Infof("BasicsService failed err:%v\n", err)
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	_, err = rpc.BasicsService.GetVideoInfoById(ctx, &basics.GetVideoByIdRequest{VideoId: videoId})
	if err != nil {
		hlog.Infof("BasicsService failed err:%v\n", err)
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	if actionType == 1 {
		if myId == 0 || videoId == 0 || *commentText == "" {
			hlog.Infof("myId or videoId or commentText is null")
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "myId or videoId or commentText is null"))
			return
		}
		//????????????
		addCommentResponse, err := rpc.InteractionService.AddComment(ctx, &interaction.AddCommentRequest{UserId: myId, VideoId: videoId, CommentText: *commentText})
		if err != nil {
			hlog.Infof("InteractionService failed err:%v\n", err)
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
			return
		}
		comment := addCommentResponse.Comment
		if myId == 0 || comment.User.Id == 0 {
			hlog.Infof("myId or comment.User.Id is null")
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "myId or comment.User.Id is null"))
			return
		}
		societyInfoResponse, err := rpc.SocietyService.SocietyInfo(ctx, &society.SocietyInfoRequest{MyId: myId, UserId: comment.User.Id})
		if err != nil {
			hlog.Infof("SocietyService failed err:%v\n", err)
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
			return
		}
		resp.StatusCode = 0
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "success"
		resp.Comment = &first.Comment{
			Id: comment.Id,
			User: &first.User{
				Id:            comment.User.Id,
				Name:          comment.User.Name,
				FollowCount:   &societyInfoResponse.FollowCount,
				FollowerCount: &societyInfoResponse.FollowerCount,
				IsFollow:      societyInfoResponse.IsFollow,
			},
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		}
		c.JSON(consts.StatusOK, resp)
	} else if actionType == 2 {
		if *commentId == 0 {
			hlog.Infof("commentId is null")
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "commentId is null"))
			return
		}
		//????????????
		commentByIdResponse, err := rpc.InteractionService.GetCommentById(ctx, &interaction.GetCommentByIdRequest{CommentId: *commentId})
		if err != nil {
			hlog.Infof("InteractionService failed err:%v\n", err)
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
			return
		}
		if commentByIdResponse.UserId != myId {
			hlog.Infof("auth failed")
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "auth failed."))
			return
		}
		if *commentId == 0 {
			hlog.Infof("commentId is null")
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "commentId is null"))
			return
		}
		_, err = rpc.InteractionService.DeleteComment(ctx, &interaction.DeleteCommentRequest{CommentId: *commentId})
		if err != nil {
			hlog.Infof("InteractionService failed err:%v\n", err)
			c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
			return
		}
		resp.StatusCode = 0
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "success"
		c.JSON(consts.StatusOK, resp)
	} else {
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "action type valied"))
		return
	}
}

// CommentList .
// @router /douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req first.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	//?????????????????????
	videoId := req.VideoId
	myId := c.GetInt64("myId")
	_, err = rpc.BasicsService.GetUserInfoById(ctx, &basics.GetUserRequest{UserId: myId})
	if err != nil {
		hlog.Infof("BasicsService failed err:%v\n", err)
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	_, err = rpc.BasicsService.GetVideoInfoById(ctx, &basics.GetVideoByIdRequest{VideoId: videoId})
	if err != nil {
		hlog.Infof("BasicsService failed err:%v\n", err)
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}

	isLogin := false
	if err != nil {
		hlog.Infof("token parse failed???not logined")
		isLogin = false
	} else {
		hlog.Infof("token parse success???loginning")
		isLogin = true
	}
	resp := new(first.DouyinCommentListResponse)
	if videoId == 0 {
		hlog.Infof("videoId is null")
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "videoId is null"))
		return
	}
	commentListResponse, err := rpc.InteractionService.CommentList(ctx, &interaction.CommentListRequest{VideoId: videoId})
	if err != nil {
		hlog.Infof("InteractionService failed err:%v", err)
		c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
		return
	}
	comments := commentListResponse.CommentList

	commentList := make([]*first.Comment, 0, len(comments))
	for _, comment := range comments {
		var societyInfoResponse *society.SocietyInfoResponse
		if isLogin {
			if myId == 0 || comment.User.Id == 0 {
				hlog.Infof("myId or comment.User.Id is null")
				c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "myId or comment.User.Id is null"))
				return
			}
			societyInfoResponse, err = rpc.SocietyService.SocietyInfo(ctx, &society.SocietyInfoRequest{MyId: myId, UserId: comment.User.Id})
			if err != nil {
				hlog.Infof("InteractionService failed err:%v", err)
				c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
				return
			}
		} else {
			if comment.User.Id == 0 {
				hlog.Infof("comment.User.Id is null")
				c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, "comment.User.Id is null"))
				return
			}
			societyInfoResponse, err = rpc.SocietyService.SocietyInfo(ctx, &society.SocietyInfoRequest{MyId: comment.User.Id, UserId: comment.User.Id})
			if err != nil {
				hlog.Infof("InteractionService failed err:%v", err)
				c.JSON(consts.StatusBadRequest, returnErrorResponse(consts.StatusBadRequest, err.Error()))
				return
			}
		}
		commentList = append(commentList, &first.Comment{
			Id: comment.Id,
			User: &first.User{
				Id:            comment.User.Id,
				Name:          comment.User.Name,
				FollowCount:   &societyInfoResponse.FollowCount,
				FollowerCount: &societyInfoResponse.FollowerCount,
				IsFollow:      societyInfoResponse.IsFollow,
			},
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		})
	}
	resp.StatusCode = 0
	resp.StatusMsg = new(string)
	*resp.StatusMsg = "success"
	resp.CommentList = commentList
	c.JSON(consts.StatusOK, resp)
}

func returnErrorResponse(code int32, msg string) core.CommonResponse {
	if strings.Contains(msg, "remote or network error[remote]: biz error: ") {
		msg = strings.Replace(msg, "remote or network error[remote]: biz error: ", "", 1)
	}
	return core.CommonResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
}
