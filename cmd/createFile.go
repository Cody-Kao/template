/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run:   generateTemplate,
}

func init() {
	rootCmd.AddCommand(greetCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetCmd.PersistentFlags().String("foo", "", "A help for foo")

	greetCmd.Flags().StringP("modName", "m", "default", "type a module name for mod init")
}

func runInitModCommand(cmd *cobra.Command) {
	moduleName, _ := cmd.Flags().GetString("modName")
	fmt.Println(moduleName)

	// Run 'go mod init' command
	cmdGoModInit := exec.Command("go", "mod", "init", moduleName)
	cmdGoModInit.Stdout = os.Stdout
	cmdGoModInit.Stderr = os.Stderr

	err := cmdGoModInit.Run()
	if err != nil {
		fmt.Printf("Error running 'go mod init': %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Go module initialized with name '%s'\n", moduleName)
}

func generateTemplate(cmd *cobra.Command, args []string) {
	// Create a directory for the Gin template
	templateDir := "gin-template"
	err := os.Mkdir(templateDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Specify the source directory containing template files
	sourceDir := "fileTemplates" // Replace with the actual path to your source folder

	// Copy all files and subdirectories from the source directory to the template directory
	err = copyFolderContents(sourceDir, templateDir)
	if err != nil {
		fmt.Printf("Error copying template files: %v\n", err)
		os.Exit(1)
	}
	runInitModCommand(cmd)
	fmt.Println("Gin template generated successfully!")
}

func copyFolderContents(src, dest string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectories
			err := os.Mkdir(destPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				os.Exit(1)
			}
			err = copyFolderContents(sourcePath, destPath)
			if err != nil {
				return err
			}
		} else {
			// Copy individual files
			err = copyFile(sourcePath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}
	return nil
}
