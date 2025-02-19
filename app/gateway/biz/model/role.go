package model

import (
	"context"
	"gorm.io/gorm"
)

type Role struct {
	ID   int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}

func (Role) TableName() string {
	return "role"
}

type UserRole struct {
	ID  int64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UID int64 `json:"uid" gorm:"column:uid"`
	RID int64 `json:"rid" gorm:"column:rid"`
}

func (UserRole) TableName() string {
	return "user_role"
}

func GetRolesByUid(db *gorm.DB, _ context.Context, uid int64) ([]Role, error) {
	var roles []Role
	err := db.Model(&UserRole{}).
		Joins("LEFT JOIN role on role.id=user_role.rid").
		Select("role.id,role.name").
		Where("user_role.uid=?", uid).
		Scan(&roles).Error
	return roles, err
}

func GetRoleByName(db *gorm.DB, _ context.Context, name string) (*Role, error) {
	var role Role
	err := db.Where("name= ?", name).First(&role).Error
	return &role, err
}

func CreateUserRole(db *gorm.DB, _ context.Context, uid, rid int64) error {
	return db.Create(&UserRole{UID: uid, RID: rid}).Error
}
