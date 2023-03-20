package util

import (
	"github.com/tsingsun/woocoo/pkg/conf"
	"os"
	"path/filepath"
)

func CreateFileName(id, key, name string) string {
	fn := key + ":" + id + filepath.Ext(name)
	return filepath.Join(conf.Abs(conf.Global().String("workflow.deploymentDir")), fn)
}

func SaveResource(id, key, name string, data []byte) error {
	filename := CreateFileName(id, key, name)
	return os.WriteFile(filename, data, 0644)
}

func OpenResource(id, key, name string) ([]byte, error) {
	fn := CreateFileName(id, key, name)
	// read file
	f, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return f, nil
}
