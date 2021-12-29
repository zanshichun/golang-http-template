package constx

const (
	RESPONSE_STATUS_SUCCESS    int    = 200                                    // status：处理成功
	RESPONSE_STATUS_FAILED     int    = 15000                                  // status: 处理失败
	ERROR_VALUES_IS_NOT_BASE64 string = "The request body value is not base64" //异常： 请求体body编码不为base64
	RESPONSE_MESSAGE_SUCCESS   string = "Successful"
	ERROR_VALUES_NO_REQ_BODY   string = "This request body got an error."
	ERROR_VALUES_DECODE_FAILED string = "Decoding failed"
)