// Code generated by Kitex v0.4.4. DO NOT EDIT.

package interactionservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	first "interaction.rpc/kitex_gen/douyin/extra/first"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddVideoFavorite(ctx context.Context, Req *first.AddVideoFavoriteRequest, callOptions ...callopt.Option) (r *first.AddVideoFavoriteResponse, err error)
	CancelVideoFavorite(ctx context.Context, Req *first.CancelVideoFavoriteRequest, callOptions ...callopt.Option) (r *first.CancelVideoFavoriteResponse, err error)
	GetFavoriteList(ctx context.Context, Req *first.GetFavoriteListRequest, callOptions ...callopt.Option) (r *first.GetFavoriteListResponse, err error)
	AddComment(ctx context.Context, Req *first.AddCommentRequest, callOptions ...callopt.Option) (r *first.AddCommentResponse, err error)
	DeleteComment(ctx context.Context, Req *first.DeleteCommentRequest, callOptions ...callopt.Option) (r *first.DeleteCommentResponse, err error)
	CommentList(ctx context.Context, Req *first.CommentListRequest, callOptions ...callopt.Option) (r *first.CommentListResponse, err error)
	GetVideoFavoriteCount(ctx context.Context, Req *first.GetVideoFavoriteCountRequest, callOptions ...callopt.Option) (r *first.GetVideoFavoriteCountResponse, err error)
	GetVideoCommentCount(ctx context.Context, Req *first.GetVideoCommentCountRequest, callOptions ...callopt.Option) (r *first.GetVideoCommentCountResponse, err error)
	IsFavorite(ctx context.Context, Req *first.IsFavoriteRequest, callOptions ...callopt.Option) (r *first.IsFavoriteResponse, err error)
	GetCommentById(ctx context.Context, Req *first.GetCommentByIdRequest, callOptions ...callopt.Option) (r *first.GetCommentByIdResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kInteractionServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kInteractionServiceClient struct {
	*kClient
}

func (p *kInteractionServiceClient) AddVideoFavorite(ctx context.Context, Req *first.AddVideoFavoriteRequest, callOptions ...callopt.Option) (r *first.AddVideoFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddVideoFavorite(ctx, Req)
}

func (p *kInteractionServiceClient) CancelVideoFavorite(ctx context.Context, Req *first.CancelVideoFavoriteRequest, callOptions ...callopt.Option) (r *first.CancelVideoFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelVideoFavorite(ctx, Req)
}

func (p *kInteractionServiceClient) GetFavoriteList(ctx context.Context, Req *first.GetFavoriteListRequest, callOptions ...callopt.Option) (r *first.GetFavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavoriteList(ctx, Req)
}

func (p *kInteractionServiceClient) AddComment(ctx context.Context, Req *first.AddCommentRequest, callOptions ...callopt.Option) (r *first.AddCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddComment(ctx, Req)
}

func (p *kInteractionServiceClient) DeleteComment(ctx context.Context, Req *first.DeleteCommentRequest, callOptions ...callopt.Option) (r *first.DeleteCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, Req)
}

func (p *kInteractionServiceClient) CommentList(ctx context.Context, Req *first.CommentListRequest, callOptions ...callopt.Option) (r *first.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, Req)
}

func (p *kInteractionServiceClient) GetVideoFavoriteCount(ctx context.Context, Req *first.GetVideoFavoriteCountRequest, callOptions ...callopt.Option) (r *first.GetVideoFavoriteCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoFavoriteCount(ctx, Req)
}

func (p *kInteractionServiceClient) GetVideoCommentCount(ctx context.Context, Req *first.GetVideoCommentCountRequest, callOptions ...callopt.Option) (r *first.GetVideoCommentCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoCommentCount(ctx, Req)
}

func (p *kInteractionServiceClient) IsFavorite(ctx context.Context, Req *first.IsFavoriteRequest, callOptions ...callopt.Option) (r *first.IsFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFavorite(ctx, Req)
}

func (p *kInteractionServiceClient) GetCommentById(ctx context.Context, Req *first.GetCommentByIdRequest, callOptions ...callopt.Option) (r *first.GetCommentByIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentById(ctx, Req)
}