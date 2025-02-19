package redis

import "strconv"

func GetBlacklistUserIDKey(userID int64) string {
	return "blacklist:userid:" + strconv.FormatInt(userID, 10)
}
