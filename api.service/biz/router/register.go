// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	core "api.service/biz/router/core"
	first "api.service/biz/router/first"
	second "api.service/biz/router/second"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	second.Register(r)

	first.Register(r)

	core.Register(r)
}