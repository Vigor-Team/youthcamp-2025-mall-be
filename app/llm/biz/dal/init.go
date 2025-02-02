package dal

import "github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal/mongo"

func Init() {
	//redis.Init()
	//mysql.Init()
	mongo.Init()
}
