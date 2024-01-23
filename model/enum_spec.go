package model

// EnumSpec describes the specification for an enum.
type EnumSpec struct {
	Package            string
	Name               string
	Variants           []Variant
	UnderlyingTypeName string
	AdditionalImports  []string
}

// Variant describes an enum variant.
type Variant struct {
	GoName   string
	EnumName *string
	Data     any

	// TODO: synonyms?
}

func (v *Variant) String() string {
	if v.EnumName != nil {
		return *v.EnumName
	}

	return v.GoName
}
