package fileTemplates

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type Creator struct {
	folderPath string
	moduleName string
	packages   []string
	frameWork  *FrameWork
}

func CreatorInit() *Creator {
	c := &Creator{}
	return c
}

func (c *Creator) SetUpProjectNameAndModuleName() error {
	// 創造使用者想要的專案名稱
	var templateName string
	fmt.Print("What is the name of project: ")
	_, err := fmt.Scan(&templateName)
	if err != nil {
		return err
	}

	// 把使用者的桌面路徑找出，這樣才能把檔案建立在桌面
	addr, err := GetUserDeskTop()
	if err != nil {
		return err
	}
	templateDir := filepath.Join(addr, templateName) // 組合成完整路徑
	err = os.Mkdir(templateDir, os.ModePerm)         // 創造空資料夾
	if err != nil {
		return err
	}
	c.folderPath = templateDir

	// 設定name for the mod init
	var modName string
	fmt.Print("What is the name for the mod init: ")
	_, err = fmt.Scan(&modName)
	if err != nil {
		return err
	}
	c.moduleName = modName

	return nil
}

func (c *Creator) SelectFrameWork() error {
	var framework string
	prompt := &survey.Select{
		Message: "Please choose a framework:",
		Options: append(Options, "End"),
	}
	survey.AskOne(prompt, &framework)
	if framework == "End" {
		return fmt.Errorf("terminated by user")
	}

	c.frameWork = FrameWorkMap[framework]
	fmt.Printf("You Choose Framework %s\n", framework)
	return nil
}

func (c *Creator) RunGitInit() error {
	// Run 'git init' command
	cmdGitInit := exec.Command("git", "init")
	// Set the working directory for the command
	cmdGitInit.Dir = c.folderPath
	cmdGitInit.Stdout = os.Stdout
	cmdGitInit.Stderr = os.Stderr

	err := cmdGitInit.Run()
	if err != nil {
		return err
	}

	fmt.Println("git init execute successfully")
	return nil
}

func (c *Creator) RunInitModCommand() error {
	fmt.Println("project path: ", c.folderPath)

	// Run 'go mod init' command
	cmdGoModInit := exec.Command("go", "mod", "init", c.moduleName)
	// Set the working directory for the command
	cmdGoModInit.Dir = c.folderPath
	cmdGoModInit.Stdout = os.Stdout
	cmdGoModInit.Stderr = os.Stderr

	err := cmdGoModInit.Run()
	if err != nil {
		return err
	}

	fmt.Printf("Go module initialized with name '%s'\n", c.moduleName)
	return nil
}

func (c *Creator) AddPackage(packageName string) {
	c.packages = append(c.packages, packageName)
}

func (c *Creator) GetPackages() error {
	if len(c.packages) != 0 {
		for _, p := range c.packages {
			// Run 'go mod init' command
			cmdGoModInit := exec.Command("go", "get", p)
			// Set the working directory for the command
			cmdGoModInit.Dir = c.folderPath
			cmdGoModInit.Stdout = os.Stdout
			cmdGoModInit.Stderr = os.Stderr

			err := cmdGoModInit.Run()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Creator) Create(cmd *cobra.Command) error {
	err := c.frameWork.StartCreate(c.moduleName, c.folderPath)
	if err != nil {
		return err
	}

	return nil
}

func (c *Creator) OpenProject() error {
	// Run 'code .' command
	cmdGoModInit := exec.Command("code", ".")
	// Set the working directory for the command
	cmdGoModInit.Dir = c.folderPath
	cmdGoModInit.Stdout = os.Stdout
	cmdGoModInit.Stderr = os.Stderr

	err := cmdGoModInit.Run()
	if err != nil {
		return err
	}

	return nil
}
