package cmd

import (
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/kk-no/kul/internal"
	tmp "github.com/kk-no/kul/internal/template"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize command",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var initGoCmd = &cobra.Command{
	Use:   "go",
	Short: "initialize command for the Go",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Println("init go called with module name", name)
		version, err := internal.GetGoModuleVersion()
		if err != nil {
			return err
		}
		fmt.Printf("go version detected: %s\n", version)
		if err := createGoFiles(name, version); err != nil {
			fmt.Printf("go files creation failed: %s\n", err)
			return err
		}
		fmt.Println("create go files succeeded")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initGoCmd)
}

type GoModuleElement struct {
	Name    string
	Version string
}

const (
	GoModuleFileName = "go.mod"
	GoMainFileName   = "main.go"

	GoModuleTemplateName = "go.mod.txt"
	GoMainTemplateName   = "main.go.txt"
)

func createGoFiles(name, version string) error {
	moduleFS := template.Must(template.ParseFS(tmp.Templates, GoModuleTemplateName))
	mainFS := template.Must(template.ParseFS(tmp.Templates, GoMainTemplateName))

	main, err := createFile(GoMainFileName)
	if err != nil {
		return err
	}
	defer main.Close()

	module, err := createFile(GoModuleFileName)
	if err != nil {
		return err
	}
	defer module.Close()

	if err := moduleFS.Execute(module, &GoModuleElement{Name: name, Version: version}); err != nil {
		return err
	}
	if err := mainFS.Execute(main, nil); err != nil {
		return err
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	if _, err := os.Stat(name); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return os.Create(name)
		}
		return nil, fmt.Errorf("error: %w", err)
	}
	return nil, fmt.Errorf("file already exists: %s", name)
}
