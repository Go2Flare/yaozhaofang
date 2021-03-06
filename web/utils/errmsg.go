package utils

const (
	RECODE_OK        = "0"    //"成功"
	RECODE_DBERR     = "4001" //"数据库查询错误"
	RECODE_NODATA    = "4002" //"无数据"
	RECODE_DATAEXIST = "4003" //"数据已存在"
	RECODE_DATAERR   = "4004" //"数据错误"

	RECODE_SESSIONERR = "4101" //"用户未登录"
	RECODE_LOGINERR   = "4102" //"用户登录失败"
	RECODE_PARAMERR   = "4103" //"参数错误"
	RECODE_USERONERR  = "4104" //"用户已经注册"
	RECODE_ROLEERR    = "4105" //"用户身份错误"
	RECODE_PWDERR     = "4106" //"密码错误"
	RECODE_USERERR    = "4107" //"用户不存在或未激活"
	RECODE_IMGSERR    = "4108" //"图片验证码错误"
	RECODE_SMSERR     = "4109" //"短信验证码错误"
	RECODE_MOBILEERR  = "4110" //"手机号错误"

	RECODE_REQERR    = "4201" //"非法请求或请求次数受限"
	RECODE_IPERR     = "4202" //"IP受限"
	RECODE_THIRDERR  = "4301" //"第三方系统错误"
	RECODE_IOERR     = "4302" //"文件读写错误"
	RECODE_SERVERERR = "4500" //"内部错误"
	RECODE_UNKNOWERR = "4501" //"未知错误"
)

var recodeText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_USERONERR:  "用户已经注册",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
	RECODE_IMGSERR:    "图片验证码错误",
	RECODE_SMSERR:     "短信验证码错误",
	RECODE_MOBILEERR:  "手机号错误",
}

//RecodeText 根据errmsg 获取具体信息
func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
