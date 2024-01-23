package templates

import (
	"embed"
	"io"
	"text/template"

	"github.com/ShawnROGrady/goenum/model"
)

var (
	//go:embed template/*
	templateContents embed.FS

	enumNameTemplate = template.Must(
		template.New("enumNameTemplate").
			Funcs(template.FuncMap{
				"unexported":    unexported,
				"paddedGoNames": paddedGoNames,
			}).
			ParseFS(templateContents,
				"template/enum_name.tmpl",
			),
	)

	enumValueTemplate = template.Must(
		template.New("enumValueTemplate").
			Funcs(template.FuncMap{
				"unexported":    unexported,
				"paddedStrings": paddedStrings,
			}).
			ParseFS(templateContents,
				"template/enum_value.tmpl",
			),
	)

	stringMethodTemplate = template.Must(
		template.New("stringMethodTemplate").
			Funcs(template.FuncMap{
				"unexported": unexported,
			}).
			ParseFS(templateContents,
				"template/string_method.tmpl",
			),
	)

	invalidNameErrorTemplate = template.Must(
		template.New("invalidNameErrorTemplate").
			ParseFS(templateContents,
				"template/invalid_name_error.tmpl",
			),
	)

	fromStringTemplate = template.Must(
		template.New("fromStringTemplate").
			Funcs(template.FuncMap{
				"unexported": unexported,
			}).
			ParseFS(templateContents,
				"template/from_string.tmpl",
			),
	)

	implTextMarshalerTemplate = template.Must(
		template.New("implTextMarshalerTemplate").
			ParseFS(templateContents,
				"template/impl_text_marshaler.tmpl",
			),
	)

	implTextUnmarshalerTemplate = template.Must(
		template.New("implTextUnmarshalerTemplate").
			Funcs(template.FuncMap{
				"unexported": unexported,
			}).
			ParseFS(templateContents,
				"template/impl_text_unmarshaler.tmpl",
			),
	)

	implSqlScannerTemplate = template.Must(
		template.New("implSqlScannerTemplate").
			ParseFS(templateContents,
				"template/impl_sql_scanner.tmpl",
			),
	)

	implSqlDriverValuerTemplate = template.Must(
		template.New("implSqlDriverValuerTemplate").
			ParseFS(templateContents,
				"template/impl_sql_driver_valuer.tmpl",
			),
	)

	headerTemplate = template.Must(
		template.New("headerTemplate").
			Funcs(template.FuncMap{
				"imports": imports,
			}).
			ParseFS(templateContents,
				"template/header.tmpl",
			),
	)

	goStringMethodTemplate = template.Must(
		template.New("goStringMethodTemplate").
			ParseFS(templateContents,
				"template/go_string.tmpl",
			),
	)
)

func WriteEnumNames(w io.Writer, enumSpec *model.EnumSpec) error {
	return enumNameTemplate.ExecuteTemplate(w, "enum_name", enumSpec)
}

func WriteEnumValues(w io.Writer, enumSpec *model.EnumSpec) error {
	return enumValueTemplate.ExecuteTemplate(w, "enum_value", enumSpec)
}

func WriteStringMethod(w io.Writer, enumSpec *model.EnumSpec) error {
	return stringMethodTemplate.ExecuteTemplate(w, "string_method", enumSpec)
}

func WriteInvalidNameError(w io.Writer, enumSpec *model.EnumSpec) error {
	return invalidNameErrorTemplate.ExecuteTemplate(w, "invalid_name_error", enumSpec)
}

func WriteFromString(w io.Writer, params *FromStringParams) error {
	return fromStringTemplate.ExecuteTemplate(w, "from_string", params)
}

func WriteImplTextMarshaler(w io.Writer, enumSpec *model.EnumSpec) error {
	return implTextMarshalerTemplate.ExecuteTemplate(w, "impl_text_marshaler", enumSpec)
}

func WriteImplTextUnmarshaler(w io.Writer, enumSpec *model.EnumSpec) error {
	return implTextUnmarshalerTemplate.ExecuteTemplate(w, "impl_text_unmarshaler", enumSpec)
}

func WriteImplSqlScanner(w io.Writer, enumSpec *model.EnumSpec) error {
	return implSqlScannerTemplate.ExecuteTemplate(w, "impl_sql_scanner", enumSpec)
}

func WriteImplSqlDriverValuer(w io.Writer, enumSpec *model.EnumSpec) error {
	return implSqlDriverValuerTemplate.ExecuteTemplate(w, "impl_sql_driver_valuer", enumSpec)
}

func WriteHeader(w io.Writer, enumSpec *model.EnumSpec) error {
	return headerTemplate.ExecuteTemplate(w, "header", enumSpec)
}

func WriteGoStringMethod(w io.Writer, enumSpec *model.EnumSpec) error {
	return goStringMethodTemplate.ExecuteTemplate(w, "go_string", enumSpec)
}
