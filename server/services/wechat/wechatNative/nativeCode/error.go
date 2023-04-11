package nativeCode

const (
	TRADE_ERROR           = 403 //	交易错误	因业务原因交易失败，请查看接口返回的详细信息
	SYSTEMERROR           = 500 //	系统错误	系统异常，请用相同参数重新调用
	SIGN_ERROR            = 401 //	签名错误	请检查签名参数和方法是否都符合签名算法要求
	RULELIMIT             = 403 //	业务规则限制	因业务规则限制请求频率，请查看接口返回的详细信息
	PARAM_ERROR           = 400 //	参数错误	请根据接口返回的详细信息检查请求参数
	OUT_TRADE_NO_USED     = 403 //	商户订单号重复	请核实商户订单号是否重复提交
	ORDERNOTEXIST         = 404 //	订单不存在	请检查订单是否发起过交易
	ORDER_CLOSED          = 400 //	订单已关闭	当前订单已关闭，请重新下单
	OPENID_MISMATCH       = 500 //	openid和appid不匹配	请确认openid和appid是否匹配
	NOAUTH                = 403 //	商户无权限	请商户前往申请此接口相关权限
	MCH_NOT_EXISTS        = 400 //	商户号不存在	请检查商户号是否正确
	INVALID_TRANSACTIONID = 500 //	订单号非法	请检查微信支付订单号是否正确
	INVALID_REQUEST       = 400 //	无效请求	请根据接口返回的详细信息检查
	FREQUENCY_LIMITED     = 429 //	频率超限	请降低请求接口频率
	BANKERROR             = 500 //	银行系统异常	银行系统异常，请用相同参数重新调用
	APPID_MCHID_NOT_MATCH = 400 //	appid和mch_id不匹配	请确认appid和mch_id是否匹配
	ACCOUNTERROR          = 403 //	账号异常	用户账号异常，无需更多操作
)
