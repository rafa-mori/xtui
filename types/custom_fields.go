package types

// New interfaces and structs for customization options, validation, and layout

type CustomizableField interface {
	FormInputObject[any]
	Label() string
	DefaultValue() string
	Group() string
}
type CustomField struct {
	Lbl  string
	DVal string
	Grp  string
}

func (f CustomField) Label() string        { return f.Lbl }
func (f CustomField) DefaultValue() string { return f.DVal }
func (f CustomField) Group() string        { return f.Grp }
