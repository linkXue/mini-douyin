// Code generated by hertz generator.

package Core

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	hertzSentinel "github.com/hertz-contrib/opensergo/sentinel/adapter"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		hertzSentinel.SentinelServerMiddleware(
			hertzSentinel.WithServerResourceExtractor(func(ctx context.Context, c *app.RequestContext) string {
				return "core_service"
			}),
			hertzSentinel.WithServerBlockFallback(func(ctx context.Context, c *app.RequestContext) {
				c.AbortWithStatusJSON(400, utils.H{
					"status_code": 10222,
					"status_msg":  "too many request; the quota used up",
				})
			}),
		),
		func(ctx context.Context, c *app.RequestContext) {
			hlog.Infof("entry %s", c.FullPath())
		},
	}
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedrequestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userrequestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publish_ctionrequestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistrequestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginrequestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerrequestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}