package differ

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
	"github.com/r3labs/diff/v3"
)

type productPODiffer struct{}

var ProductPODiffer *productPODiffer

func (differ *productPODiffer) GetChangedMap(origin, target *po.Product) map[string]interface{} {
	changedMap := make(map[string]interface{})

	if !compareCategories(origin.Categories, target.Categories) {
		changedMap["Categories"] = target.Categories
		changedMap["ID"] = target.ID
	}

	d, _ := diff.NewDiffer(diff.TagName("json"))
	changeLog, _ := d.Diff(origin, target)

	for _, change := range changeLog {
		if depth := len(change.Path); depth != 1 {
			continue
		}
		if change.Type == diff.UPDATE {
			changedMap[change.Path[0]] = change.To
		}
	}
	return changedMap
}

func compareCategories(origin, target []po.Category) bool {
	if len(origin) != len(target) {
		return false
	}

	for i := range origin {
		if origin[i].ID != 0 && target[i].ID != 0 {
			if origin[i].ID != target[i].ID {
				return false
			}
		} else {
			if origin[i].Name != target[i].Name {
				return false
			}
		}
	}
	return true
}
