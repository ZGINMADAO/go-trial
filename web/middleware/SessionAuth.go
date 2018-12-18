package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

func SessionAuth(ctx iris.Context, session *sessions.Session) {

	whiteList := []string{"/admin/auth/login"}
	profile := session.Get("admin_session_profile")
	for _, val := range whiteList {
		if val == ctx.Path() {
			ctx.Next()
		}
	}

	if profile == nil {
		ctx.Redirect("/admin/auth/login")
		return
	}

	ctx.Next()
}
