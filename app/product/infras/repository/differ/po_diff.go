package differ

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
	"github.com/r3labs/diff/v3"
)

type productPODiffer struct{}

var ProductPODiffer *productPODiffer

func (differ *productPODiffer) GetChangedMap(origin, target *po.Product) map[string]interface{} {
	d, _ := diff.NewDiffer(diff.TagName("json"))
	changedMap := make(map[string]interface{})
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
