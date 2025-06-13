package xtui

import (
	c "github.com/rafa-mori/xtui/components"
	t "github.com/rafa-mori/xtui/types"
)

type Config struct{ t.Config }
type FormFields = t.FormFields
type FormField = t.FormInputObject[any]
type InputField = *t.InputObject[any]

func ShowForm(form Config) (map[string]string, error) {
	return c.ShowForm(form.Config)
}

func NewConfig(title string, fields FormFields) Config {
	return Config{Config: t.Config{Title: title, Fields: fields}}
}
func NewInputField(placeholder string, typ string, value string, required bool, minValue int, maxValue int, err string, validation func(string) error) FormField {
	return &t.InputObject[any]{
		Val: value,
	}
}

func NewFormFields(title string, fields []FormField) FormFields {
	ffs := make([]FormField, len(fields))
	for i, f := range fields {
		ffs[i] = f
	}
	return FormFields{
		Title:  title,
		Fields: ffs,
	}
}
func NewFormModel(config t.Config) (map[string]string, error) { return c.ShowForm(config) }
