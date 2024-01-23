package templates

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ShawnROGrady/goenum/model"
)

func isAsciiLower(r rune) bool {
	return r >= 97 && r <= 122
}

func isAsciiUpper(r rune) bool {
	return r >= 65 && r <= 90
}

func unexported(name string) string {
	rs := []rune(name)
	if !isAsciiUpper(rs[0]) {
		return name
	}

	rs[0] = rs[0] + 32
	return string(rs)
}

func paddedGoNames(rawVariants []model.Variant) []model.Variant {
	maxLen := 0
	for _, rawVariant := range rawVariants {
		if len(rawVariant.GoName) > maxLen {
			maxLen = len(rawVariant.GoName)
		}
	}

	normalizeVariant := func(rawVariant model.Variant) model.Variant {
		normalized := rawVariant
		normalized.GoName = rawVariant.GoName + ":"
		padding := maxLen - len(rawVariant.GoName)
		if padding != 0 {
			normalized.GoName += strings.Repeat(" ", padding)
		}

		if normalized.EnumName == nil {
			normalized.EnumName = &rawVariant.GoName
		}

		return normalized
	}

	normalizedVariants := make([]model.Variant, len(rawVariants))
	for i, rawVariant := range rawVariants {
		normalizedVariants[i] = normalizeVariant(rawVariant)
	}

	return normalizedVariants
}

func paddedStrings(rawVariants []model.Variant) []model.Variant {
	maxLen := 0
	for _, rawVariant := range rawVariants {
		if len(rawVariant.String()) > maxLen {
			maxLen = len(rawVariant.String())
		}
	}

	normalizeVariant := func(rawVariant model.Variant) model.Variant {
		normalized := rawVariant
		var normalizedEnumName string

		if normalized.EnumName == nil {
			normalizedEnumName = rawVariant.GoName
		} else {
			normalizedEnumName = *normalized.EnumName
		}

		padding := maxLen - len(normalizedEnumName)

		normalizedEnumName = fmt.Sprintf("%q", normalizedEnumName)
		normalizedEnumName += ":"

		if padding != 0 {
			normalizedEnumName += strings.Repeat(" ", padding)
		}

		normalized.EnumName = &normalizedEnumName

		return normalized
	}

	normalizedVariants := make([]model.Variant, len(rawVariants))
	for i, rawVariant := range rawVariants {
		normalizedVariants[i] = normalizeVariant(rawVariant)
	}

	return normalizedVariants
}

func imports(spec model.EnumSpec) []string {
	vs := []string{"fmt"}
	if len(spec.AdditionalImports) != 0 {
		vs = append(vs, spec.AdditionalImports...)
	}
	sort.Strings(vs)
	for i, v := range vs {
		vs[i] = strconv.Quote(v)
	}
	return vs
}
