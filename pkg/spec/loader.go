package spec

// Loader is an interface for bpmn sepc and dmn spec
type Loader interface {
	Load(data []byte) error
	LoadFromFile(path string) error
	GetId() string
	GetName() string
}

func MergeMapping(src, dst map[string]any) {
	for k, v := range src {
		dst[k] = v
	}
}
