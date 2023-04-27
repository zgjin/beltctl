package v1

type TypeMeta struct {
	Kind string `json:"kind,omitempty"`

	APIVersion string `json:"apiVersion,omitempty"`
}

type ObjectMeta struct {
	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}
