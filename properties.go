package events

//Properties event properties
type Properties interface {
	//EventProperties return event properties and any error if raised.
	EventProperties() (map[string]string, error)
}

//Map event map properties struct
type Map map[string]string

//EventProperties return event properties and any error if raised.
func (m Map) EventProperties() (map[string]string, error) {
	return m, nil
}
