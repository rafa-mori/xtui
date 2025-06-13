[//]: # (![XTui Banner]&#40;./assets/banner.png&#41;)

# XTui

---

**A high-performance, easy-to-use terminal user interface (TUI) library for Go, enabling developers to build interactive and visually appealing terminal applications with minimal effort.**

---

![Go Version](https://img.shields.io/github/go-mod/go-version/faelmori/xtui)
![License](https://img.shields.io/github/license/faelmori/xtui)
![Build Status](https://img.shields.io/github/actions/workflow/status/faelmori/xtui/build.yml)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [CLI Examples](#cli-examples)
- [Module Examples](#module-examples)
- [Hotkeys](#hotkeys)
- [Form Handling](#form-handling)
- [Data Export](#data-export)
- [Contributing](#contributing)
- [License](#license)

## Introduction

**xtui** is a high-performance, easy-to-use terminal user interface (TUI) library for Go. It enables developers to build interactive and visually appealing terminal applications with minimal effort while maintaining flexibility and performance.

## Features

- **Intuitive API** – Simplifies the creation of rich terminal interfaces.
- **Customizable Styles** – Tailor UI components to your specific needs with custom styles and configurations.
- **Interactive Form Handling** – Manage form inputs with validation, password protection, and navigation.
- **Data Filtering, Sorting, and Navigation** – Built-in support for table operations.
- **Keyboard Shortcuts** – Provides an efficient user experience with predefined hotkeys.
- **Paginated Views** – Allows smooth navigation through large datasets.
- **Multi-format Export** – Export data to CSV, YAML, JSON, and XML formats.
- **Error Logging** – Integrated with the **logz** library for error tracking and debugging.

## Installation

To install **xtui**, run the following command:

```sh
go get github.com/rafa-mori/xtui
```

## Usage

Here’s a quick example demonstrating how to use **xtui** for displaying tables:

```go
package main

import (
    "github.com/rafa-mori/xtui"
    "github.com/rafa-mori/xtui/types"
    "github.com/charmbracelet/lipgloss"
)

func main() {
    config := types.FormConfig{
        Fields: []types.Field{
            {Name: "ID", Placeholder: "Unique ID"},
            {Name: "Name", Placeholder: "User Name"},
        },
    }
    
    customStyles := map[string]lipgloss.Color{
        "Info":    lipgloss.Color("#75FBAB"),
        "Warning": lipgloss.Color("#FDFF90"),
    }
    
    if err := xtui.StartTableScreen(config, customStyles); err != nil {
        panic(err)
    }
}
```

For form-based interactions:

```go
package main

import (
    "github.com/rafa-mori/xtui"
    "github.com/rafa-mori/xtui/types"
)

func main() {
    config := types.Config{
        Title: "User Registration",
        Fields: types.FormFields{
            Inputs: []types.InputField{
                {Ph: "Name", Tp: "text", Req: true, Err: "Name is required!"},
                {Ph: "Password", Tp: "password", Req: true, Err: "Password is required!"},
            },
        },
    }
    
    result, err := xtui.ShowForm(config)
    if err != nil {
        panic(err)
    }
    println("Form submitted:", result)
}
```

## CLI Examples

### Install Applications Command

```sh
go run main.go app-install --application app1 --application app2 --path /usr/local/bin --yes --quiet
```

### Table View Command

```sh
go run main.go table-view
```

### Input Form Command

```sh
go run main.go input-form
```

### Loader Form Command

```sh
go run main.go loader-form
```

## Module Examples

### Log Viewer

```go
package main

import (
    "github.com/rafa-mori/xtui/wrappers"
)

func main() {
    if err := wrappers.LogViewer(); err != nil {
        panic(err)
    }
}
```

### Application Manager

```go
package main

import (
    "github.com/rafa-mori/xtui/wrappers"
)

func main() {
    args := []string{"app1", "app2", "/usr/local/bin", "true", "true"}
    if err := wrappers.InstallDependenciesWithUI(args...); err != nil {
        panic(err)
    }
}
```

## Hotkeys

The following keyboard shortcuts are supported out of the box:

- **q, Ctrl+C:** Exit the application.
- **Enter:** Copy selected row or submit form.
- **Ctrl+R:** Change cursor mode.
- **Tab/Shift+Tab, Up/Down Arrows:** Navigate between form fields or table rows.
- **Ctrl+E:** Export data to CSV.
- **Ctrl+Y:** Export data to YAML.
- **Ctrl+J:** Export data to JSON.
- **Ctrl+X:** Export data to XML.

## Form Handling

**xtui** provides an intuitive API for managing forms with validations:
- **Field Validation:** Enforce required fields, minimum/maximum length, and custom validators.
- **Password Input:** Securely handle password fields with hidden characters.
- **Dynamic Properties:** Automatically adapt form inputs based on external configurations.

### Example

```go
field := types.InputField{
    Ph:  "Email",
    Tp:  "text",
    Req: true,
    Err: "Valid email is required!",
    Vld: func(value string) error {
        if !strings.Contains(value, "@") {
            return fmt.Errorf("Invalid email format")
        }
        return nil
    },
}
```

## Data Export

**xtui** supports exporting table data in multiple formats:
- **CSV:** Saves data as a comma-separated values file.
- **YAML:** Outputs data in a structured YAML format.
- **JSON:** Encodes data into a compact JSON format.
- **XML:** Exports data as XML for interoperability.

### Example

To export data to a file, simply use the respective hotkey (e.g., `Ctrl+E` for CSV). Files will be saved with predefined names, such as `exported_data.csv`.

## Contributing

We welcome contributions of all kinds! Whether it’s reporting issues, improving documentation, or submitting new features, your help is appreciated. Please check out our [contributing guidelines](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.