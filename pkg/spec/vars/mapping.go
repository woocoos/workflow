package vars

// Mapping holds the values of inputs or outputs, value must be serializable.
type Mapping map[string]any

func (Mapping) Equals(a, b interface{}) bool {
	return a == b
}
