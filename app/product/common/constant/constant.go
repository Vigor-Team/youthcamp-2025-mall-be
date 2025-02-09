package constant

type ProductStatus = uint32

const (
	ProductStatusOnline  ProductStatus = 0
	ProductStatusOffline ProductStatus = 1
	ProductStatusDelete  ProductStatus = 2
)

var ProductStatusDescMap = map[ProductStatus]string{
	ProductStatusOnline:  "上架",
	ProductStatusOffline: "下架",
	ProductStatusDelete:  "删除",
}

type StateOperationType = uint32

const (
	StateOperationTypeAdd     StateOperationType = 1
	StateOperationTypeSave    StateOperationType = 2
	StateOperationTypeDel     StateOperationType = 3
	StateOperationTypeOffline StateOperationType = 4
	StateOperationTypeOnline  StateOperationType = 5
)
