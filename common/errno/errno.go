package errno

var (
	SuccessCode         int32 = 0
	ErrInternal         int32 = 10000
	ErrMysql            int32 = 10001
	ErrRedis            int32 = 10002
	ErrMongo            int32 = 10003
	ErrHTTPRequestParam int32 = 10004
	ErrGRPCRequestParam int32 = 10005
)
