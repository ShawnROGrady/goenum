package model

// EnumSpec describes the specification for an enum.
type EnumSpec struct {
	Package  string
	Name     string
	Variants []Variant
}

// Variant describes an enum variant.
type Variant struct {
	GoName   string
	EnumName string

	// TODO: synonyms?
}

func (v *Variant) String() string {
	if v.EnumName != "" {
		return v.EnumName
	}

	return v.GoName
}
