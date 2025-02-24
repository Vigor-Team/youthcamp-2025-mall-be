package redis

func GetBlacklistUserIDKey(userID int32) string {
	return "blacklist:userid:" + string(userID)
}
