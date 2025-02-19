package model

type Permission struct {
	ID int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	V1 string `json:"v1" gorm:"column:v1"`
	V2 string `json:"v2" gorm:"column:v2"`
}

func (Permission) TableName() string {
	return "permission"
}

type PermissionRole struct {
	ID  int64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	PID int64 `json:"pid" gorm:"column:pid"`
	RID int64 `json:"rid" gorm:"column:rid"`
}

func (PermissionRole) TableName() string {
	return "permission_role"
}
