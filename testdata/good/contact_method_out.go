// Code generated by goenum. DO NOT EDIT.

package contactmethod

import "fmt"

var preferredContactMethodName = map[PreferredContactMethod]string{
	PreferredContactMethodUnspecified: "PreferredContactMethodUnspecified",
	ContactByEmail:                    "EMAIL",
	ContactByCellphone:                "CELLPHONE",
	ContactByLandLine:                 "LANDLINE",
}

var preferredContactMethodValue = map[string]PreferredContactMethod{
	"PreferredContactMethodUnspecified": PreferredContactMethodUnspecified,
	"EMAIL":                             ContactByEmail,
	"CELLPHONE":                         ContactByCellphone,
	"LANDLINE":                          ContactByLandLine,
}

func (e PreferredContactMethod) String() string {
	return preferredContactMethodName[e]
}

func (e PreferredContactMethod) GoString() string {
	switch e {
	case PreferredContactMethodUnspecified:
		return "contactmethod.PreferredContactMethodUnspecified"
	case ContactByEmail:
		return "contactmethod.ContactByEmail"
	case ContactByCellphone:
		return "contactmethod.ContactByCellphone"
	case ContactByLandLine:
		return "contactmethod.ContactByLandLine"
	default:
		return fmt.Sprintf("contactmethod.PreferredContactMethod(%v)", int(e))
	}
}

type InvalidPreferredContactMethodNameError struct {
	Name string
}

func (e *InvalidPreferredContactMethodNameError) Error() string {
	return fmt.Sprintf("invalid PreferredContactMethod %q", e.Name)
}

func FromString(name string) (PreferredContactMethod, error) {
	val, ok := preferredContactMethodValue[name]
	if !ok {
		return PreferredContactMethod(0), &InvalidPreferredContactMethodNameError{
			Name: name,
		}
	}

	return val, nil
}

func (e PreferredContactMethod) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

func (e *PreferredContactMethod) UnmarshalText(text []byte) error {
	name := string(text)
	val, ok := preferredContactMethodValue[name]
	if !ok {
		return &InvalidPreferredContactMethodNameError{
			Name: name,
		}
	}

	*e = val
	return nil
}
