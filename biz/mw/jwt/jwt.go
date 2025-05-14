package jwt

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"github.com/qingyggg/aufer/biz/model/cmd/common"
	"github.com/qingyggg/aufer/biz/model/http/basic/user"
	"github.com/qingyggg/aufer/biz/rpc"
	user_rpc "github.com/qingyggg/aufer/cmd/user/rpc"
	"github.com/qingyggg/aufer/pkg/constants"
	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	identity      = "uid"
)

func Init() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte(constants.SecretKey),
		TokenLookup: "cookie:token",
		Timeout:     24 * time.Hour,
		MaxRefresh:  time.Hour * 6,
		IdentityKey: identity,
		// Verify password at login
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			//var loginRequest
			req := new(user.LoginOrRegRequest)
			c.Set("hasErr", true)
			if err := c.BindAndValidate(req); err != nil {
				return nil, err
			}
			uid, puid, err := user_rpc.UserLogin(rpc.Clients.UserClient, ctx, req.Email, req.Password)
			if err != nil {
				return nil, err
			}
			c.Set("hasErr", false)
			c.Set("puid", puid)
			c.Set("uid", uid)
			return puid, nil
		},
		// Set the payload in the token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					identity: v,
				}
			}
			return jwt.MapClaims{}
		},
		// build login response if verify password successfully
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			hlog.CtxInfof(ctx, "Login success ，token is issued clientIP: "+c.ClientIP())
			c.SetCookie("token", token, int(24*time.Hour), "/", "api.mols.site", protocol.CookieSameSiteNoneMode, true, true)
		},
		// Verify token and get the id of logged-in user
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(float64); ok {
				current_user_id := int64(v)
				c.Set("current_user_id", current_user_id)
				hlog.CtxInfof(ctx, "Token is verified clientIP: "+c.ClientIP())
				return true
			}
			return false
		},
		// Validation failed, build the message
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.StatusUnauthorized, &common.BaseResponse{
				StatusCode: errno.AuthorizationFailedErrCode,
				StatusMsg:  message,
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			resp := utils.BuildBaseResp(e)
			return resp.StatusMsg
		},
	})
}
