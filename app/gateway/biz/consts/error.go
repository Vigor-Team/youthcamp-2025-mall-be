package consts

var (
	ErrBanUser            int32 = 10011
	ErrBindPermissionRole int32 = 10012
	ErrBindRoleUser       int32 = 10013
	ErrCreatePermission   int32 = 10014
	ErrCreateRole         int32 = 10015
	ErrGetPermission      int32 = 10016
	ErrGetRole            int32 = 10017
)

var (
	ErrJWTNotFound       int32 = 10021
	ErrJWTInvalid        int32 = 10022
	ErrJWTCreate         int32 = 10023
	ErrJWTRefreshExpired int32 = 10024
)
