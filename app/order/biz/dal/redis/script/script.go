package script

import (
	"os"
	"path/filepath"
	"sync"
)

var (
	scriptOnce       sync.Once
	preSeckillScript string
)

func GetPreSeckillScript() (string, error) {
	var err error
	scriptOnce.Do(func() {
		prefix := "redis/script"
		scriptFileRelPath := filepath.Join(prefix, "pre_seckill.lua")
		content, err := os.ReadFile(scriptFileRelPath)
		if err != nil {
			return
		}
		preSeckillScript = string(content)
	})
	return preSeckillScript, err
}
