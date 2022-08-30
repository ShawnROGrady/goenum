package templates

import "github.com/ShawnROGrady/goenum/model"

type FromStringParams struct {
	// The enum spec.
	Enum *model.EnumSpec

	// If true, don't prefix the method with the enum's name. By default
	// for an enum type named "Foo" the generated function will be named
	// "FooFromString", but if there is only one main enum type in the
	// package it may be preferable to just have this method be named
	// "FromString". This is primarily meant for cases where the package
	// name matches the enum type name, since if the "Foo" enum type is
	// defined in a "foo" package it makes more sense to have
	// "foo.FromString" instead of "foo.FooFromString".
	ExcludePrefix bool
}
