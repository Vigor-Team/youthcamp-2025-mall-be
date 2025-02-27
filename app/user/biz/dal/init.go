package dal

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
