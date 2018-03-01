package alipay

/* ================================================================================
 * AppOauthTokenQueryResponse数据域结构
 * api: alipay.open.auth.token.app.query
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type AppOauthTokenQueryResponse struct {
	UserId      string   `form:"user_id" json:"user_id"`           //授权商户的user_id
	AuthAppId   string   `form:"auth_app_id" json:"auth_app_id"`   //授权商户的appid
	AuthMethods []string `form:"auth_methods" json:"auth_methods"` //当前app_auth_token的授权接口列表
	AuthStart   string   `form:"auth_start" json:"auth_start"`     //授权生效时间
	AuthEnd     string   `form:"auth_end" json:"auth_end"`         //授权失效时间
	ExpiresIn   string   `form:"expires_in" json:"expires_in"`     //应用授权令牌失效时间，单位到秒
	Status      string   `form:"status" json:"status"`             //valid：有效状态；invalid：无效状态
}
