package templates

import (
	"fmt"
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

		if normalized.EnumName == "" {
			normalized.EnumName = rawVariant.GoName
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
		padding := maxLen - len(rawVariant.GoName)

		if normalized.EnumName == "" {
			normalized.EnumName = rawVariant.GoName
		}

		normalized.EnumName = fmt.Sprintf("%q", normalized.EnumName)
		normalized.EnumName += ":"

		if padding != 0 {
			normalized.EnumName += strings.Repeat(" ", padding)
		}

		return normalized
	}

	normalizedVariants := make([]model.Variant, len(rawVariants))
	for i, rawVariant := range rawVariants {
		normalizedVariants[i] = normalizeVariant(rawVariant)
	}

	return normalizedVariants
}
