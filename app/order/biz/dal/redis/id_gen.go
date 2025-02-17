package redis

import (
	"context"
	"github.com/bwmarrin/snowflake"
)

var (
	OrderNode    *snowflake.Node
	PreOrderNode *snowflake.Node
)

func InitSnowflake() error {
	if OrderNode == nil {
		node, err := snowflake.NewNode(111)
		if err != nil {
			return err
		}
		OrderNode = node
	}

	if PreOrderNode == nil {
		node, err := snowflake.NewNode(222)
		if err != nil {
			return err
		}
		PreOrderNode = node
	}

	return nil
}

func NextId(_ context.Context, node *snowflake.Node) (uint32, error) {
	id := node.Generate()
	return uint32(id.Int64() & 0xFFFFFFFF), nil
}
