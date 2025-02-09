package entity

import "github.com/jinzhu/copier"

type ProductEntity struct {
	ID          uint32
	Name        string
	Description string
	Picture     string
	SpuName     string
	SpuPrice    float32
	Price       float32
	Stock       uint32
	Status      uint32
	Categories  []*CategoryEntity
}

func (entity *ProductEntity) Clone() (*ProductEntity, error) {
	ret := &ProductEntity{}
	err := copier.Copy(ret, entity)
	return ret, err
}

type ProductES struct {
	ID            uint32
	Name          string
	Description   string
	Picture       string
	SpuName       string
	SpuPrice      float32
	Price         float32
	Stock         uint32
	Status        uint32
	CategoryNames []string
}
