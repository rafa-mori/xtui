package cli

import (
	"reflect"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rafa-mori/xtui/components"
	"github.com/rafa-mori/xtui/types"
	"github.com/rafa-mori/xtui/wrappers"
	"github.com/spf13/cobra"
)

func FormsCmdsList() []*cobra.Command {
	inputCmd := InputFormCommand()
	loaderCmd := LoaderFormCommand()

	return []*cobra.Command{
		inputCmd,
		loaderCmd,
	}
}

func InputFormCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "input-form",
		Aliases: []string{"input", "formInput", "inputForm", "formInput", "form-input"},
		Short:   "Form inputs for any command",
		Long:    "Form inputs screen, interactive mode, for any command with flags",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := types.Config{
				Title: "Dynamic Form",
				Fields: types.FormFields{
					Title: "Login",
					Fields: []types.FormInputObject[any]{
						&types.Input[string]{
							Ph:  "Username",
							Tp:  reflect.TypeFor[string](),
							Val: "",
							Req: true,
							Min: 3,
							Max: 20,
							Err: "Username is required and must be between 3 and 20 characters.",
							//							: []types.ValidationRule{
							//,								func(value string) error {
							//									if len(value) < 3 || len(value) > 20 {
							//										return fmt.Errorf("username must be between 3 and 20 characters")
							//									}
							//									return nil
							//								}
							//							},
							ValidationRulesVal: []types.ValidationRule{types.Required},
						},
						&types.Input[string]{
							Ph:  "Password",
							Tp:  reflect.TypeFor[string](),
							Val: "",
							Req: true,
							Min: 6,
							Max: 20,
							Err: "Password is required and must be between 6 and 20 characters.",
							//Vld: func(value string) error {
							//	if len(value) < 6 || len(value) > 20 {
							//		return fmt.Errorf("password must be between 6 and 20 characters")
							//	}
							//	return nil
							//},
							ValidationRulesVal: []types.ValidationRule{types.Required},
						},
					},
				},
			}
			_, err := components.ShowForm(config)
			return err
		},
	}

	return cmd
}

func LoaderFormCommand() *cobra.Command {
	// Configuration file path. This file can be used to load dynamic properties.
	var configFile string
	// Loader settings and properties map. This map can be used to load dynamic properties.
	// Example: {"Loading dynamic properties...": 2, "Dynamic properties loaded successfully.": 1, "Closing loader...": 1}
	var sequenceWithDelay map[string]int
	// Loader icon sequence map. This map can be used to load dynamic properties
	// with sequenceWithDelay to set icons for each message in the loader.
	var sequenceWithIcon map[string]string
	// Loader color sequence map. This map can be used to load dynamic properties with other sequencies
	// to set colors for each message in the loader.
	var sequenceWithColor map[string]string

	cmd := &cobra.Command{
		Use:     "loader-form",
		Aliases: []string{"loader", "formLoader", "loaderForm", "formLoader", "form-loader"},
		Short:   "Form loader for any command",
		Long:    "Form loader screen, interactive mode, for any command with flags",
		RunE: func(cmd *cobra.Command, args []string) error {

			messages := make(chan tea.Msg)
			go func() {
				messages <- wrappers.LoaderMsg{Message: "Loading dynamic properties..."}
				time.Sleep(2 * time.Second)
				messages <- wrappers.LoaderMsg{Message: "Dynamic properties loaded successfully."}
				time.Sleep(1 * time.Second)
				messages <- wrappers.LoaderCloseMsg{}
			}()
			return wrappers.StartLoader(messages)
		},
	}

	cmd.Flags().StringToIntVarP(&sequenceWithDelay, "loader-delay", "l", nil, "Loader messages and delays")
	cmd.Flags().StringToStringVarP(&sequenceWithIcon, "loader-icon", "i", nil, "Loader messages and icons")
	cmd.Flags().StringToStringVarP(&sequenceWithColor, "loader-color", "r", nil, "Loader messages and colors")
	cmd.Flags().StringVarP(&configFile, "loader-config", "L", "", "Loader configuration file for dynamic properties and settings")

	return cmd
}

// Unit tests for InputFormCommand
func TestInputFormCommand(t *testing.T) {
	cmd := InputFormCommand()
	if cmd.Use != "input-form" {
		t.Errorf("expected 'input-form', got '%s'", cmd.Use)
	}
	if cmd.Short != "Form inputs for any command" {
		t.Errorf("expected 'Form inputs for any command', got '%s'", cmd.Short)
	}
	if cmd.Long != "Form inputs screen, interactive mode, for any command with flags" {
		t.Errorf("expected 'Form inputs screen, interactive mode, for any command with flags', got '%s'", cmd.Long)
	}
}

// Unit tests for LoaderFormCommand
func TestLoaderFormCommand(t *testing.T) {
	cmd := LoaderFormCommand()
	if cmd.Use != "loader-form" {
		t.Errorf("expected 'loader-form', got '%s'", cmd.Use)
	}
	if cmd.Short != "Form loader for any command" {
		t.Errorf("expected 'Form loader for any command', got '%s'", cmd.Short)
	}
	if cmd.Long != "Form loader screen, interactive mode, for any command with flags" {
		t.Errorf("expected 'Form loader screen, interactive mode, for any command with flags', got '%s'", cmd.Long)
	}
}
