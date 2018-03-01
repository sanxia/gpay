package alipay

/* ================================================================================
 * AppPayResponse数据域结构
 * api: alipay.open.auth.token.app
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type AppOauthTokenResponse struct {
	UserId          string `form:"user_id" json:"user_id"`                     //授权商户的user_id
	AuthAppId       string `form:"auth_app_id" json:"auth_app_id"`             //授权商户的appid
	AppAuthToken    string `form:"app_auth_token" json:"app_auth_token"`       //应用授权令牌
	AppRefreshToken string `form:"app_refresh_token" json:"app_refresh_token"` //刷新令牌
	ExpiresIn       string `form:"expires_in" json:"expires_in"`               //应用授权令牌的有效时间（从接口调用时间作为起始时间），单位到秒
	ReExpiresIn     string `form:"re_expires_in" json:"re_expires_in"`         //刷新令牌的有效时间（从接口调用时间作为起始时间），单位到秒
}
