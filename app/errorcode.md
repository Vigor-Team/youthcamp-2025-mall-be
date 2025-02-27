Common Error Codes:
    ErrInternal: 10000
    ErrMysql: 10001
    ErrRedis: 10002
    ErrMongo:  10003
    ErrParam: 10004

    
Auth:
    ErrAuth: 20000

User:
var (
ErrUserExist    int32 = 20001
ErrUserNotFound int32 = 20002
ErrPassword     int32 = 20003
)

product:
var (
ErrDecrProductStock int32 = 30001
ErrIncrProductStock int32 = 30002
ErrGetProduct       int32 = 30004
ErrStateOperation   int32 = 30005
ErrUpdateProduct    int32 = 30006
)

var (
ErrGetCategory int32 = 30011
)

payment:
var (
ErrCardValidate     int32 = 40001
ErrCreatePaymentLog int32 = 40002
)

order:
var (
ErrGetOrder         int32 = 50001
ErrUpdateOrder      int32 = 50002
ErrCreateOrder      int32 = 50003
ErrPreOrderValidate int32 = 50004
)

var (
ErrPubMessage int32 = 50021
)

var (
ErrOutStock int32 = 50031
ErrDuplicateUser int32 = 50032
)

llm:
var (
ErrDelConversation int32 = 60001
ErrGetHistory      int32 = 60002
ErrBuildAgent	  int32 = 60003
ErrInvokeAgent	  int32 = 60004
)

checkout:
var (
ErrRPCGetCart       int32 = 70001
ErrRPCGetProduct    int32 = 70002
ErrRPCPlaceOrder    int32 = 70003
ErrRPCEmptyCart     int32 = 70004
ErrRPCCharge        int32 = 70005
ErrRPCMarkOrderPaid int32 = 70006
)