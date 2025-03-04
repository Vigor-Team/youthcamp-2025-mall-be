package model

import (
	"context"
	"gorm.io/gorm"
)

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

func CreatePermission(db *gorm.DB, _ context.Context, permission *Permission) (*Permission, error) {
	err := db.Create(permission).Error
	return permission, err
}

func BindPermissionRole(db *gorm.DB, _ context.Context, permissionRole *PermissionRole) error {
	return db.Create(permissionRole).Error
}

func ListPermissions(db *gorm.DB, _ context.Context) ([]*Permission, error) {
	var permissions []*Permission
	err := db.Find(&permissions).Error
	return permissions, err
}
