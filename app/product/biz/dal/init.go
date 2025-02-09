package dal

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
