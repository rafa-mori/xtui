package types

type Config struct {
	Title  string
	Fields FormFields
}

func (c Config) GetTitle() string      { return c.Title }
func (c Config) GetFields() FormFields { return c.Fields }

type FormConfig struct {
	Title string
	*FormFields
}

func (f FormConfig) GetTitle() string                  { return f.Title }
func (f FormConfig) GetFields() []FormInputObject[any] { return f.Fields }
