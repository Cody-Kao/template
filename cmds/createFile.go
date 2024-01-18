/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmds

import (
	"fmt"
	"os"

	f "github.com/Cody-Kao/template/fileTemplates"
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

	// greetCmd.Flags().StringP("modName", "m", "default", "type a module name for mod init")
}

func generateTemplate(cmd *cobra.Command, args []string) {
	// initialize the Creator
	c := f.CreatorInit()

	// 創造project name跟 name of mod file
	err := c.SetUpProjectNameAndModuleName()
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// select framework
	err = c.SelectFrameWork()
	if err != nil {
		fmt.Printf("Error selecting framework: %v\n", err)
		os.Exit(1)
	}

	// run go mod init
	err = c.RunInitModCommand()
	if err != nil {
		fmt.Printf("Error running 'go mod init': %v\n", err)
		os.Exit(1)
	}

	// create files
	err = c.Create(cmd)
	if err != nil {
		fmt.Printf("Error creating template files: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("templates generated successfully!")

	// 記得註冊會使用到的所有外部package
	c.AddPackage("github.com/gorilla/mux")

	// 建立完檔案並且註冊所有外部package之後就能執行go get [package name]
	err = c.GetPackages()

	if err != nil {
		fmt.Printf("Error getting packages: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("getting packages successfully!")

	// 完成!
	fmt.Println("Done!")

	err = c.OpenProject()
	if err != nil {
		fmt.Printf("Error opening project: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Now, You Can Build Your Project!")
}
