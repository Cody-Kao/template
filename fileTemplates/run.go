package fileTemplates

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type Creator struct {
	folderPath string
	moduleName string
	packages   []string
	frameWork  FrameWork
}

func CreatorInit() *Creator {
	c := &Creator{}
	return c
}

func getUserDeskTop() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	// 創造使用者桌面路徑
	desktopPath := filepath.Join(currentUser.HomeDir, "Desktop")

	// 檢查該路徑是否存在
	_, err = os.Stat(desktopPath)
	if os.IsNotExist(err) {
		return "", err
	}

	return desktopPath, nil
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
	addr, err := getUserDeskTop()
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
		Message: "This is question1 Select the programming language:",
		Options: append(Options, "End"),
	}
	survey.AskOne(prompt, &framework)
	if framework == "End" {
		return fmt.Errorf("terminated by user")
	}

	c.frameWork = FrameWorkMap[framework]
	fmt.Printf("You Choose language %s\n", FrameWorkMap[framework])
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

func (c *Creator) CreateFiles(cmd *cobra.Command) error {
	addr, content := mainContent(c.moduleName)
	destinationPath := filepath.Join(c.folderPath, addr)
	err := createFoldersAndFiles(destinationPath, content)
	if err != nil {
		return err
	}

	err = c.frameWork.Create(c.moduleName, c.folderPath)
	if err != nil {
		return err
	}

	addr, content = templateContent()
	destinationPath = filepath.Join(c.folderPath, addr)
	err = createFoldersAndFiles(destinationPath, content)
	if err != nil {
		return err
	}

	return nil
}

func createFoldersAndFiles(destinationPath string, content string) error {
	parentDir := filepath.Dir(destinationPath) // filepath.Dir會把path的最後一項去掉，只保留前段，等於是創造parent directory的意思
	err := os.Mkdir(parentDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	fmt.Printf("parent path %s is successfully created\n", parentDir)
	file, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}
	file.Close()

	fmt.Printf("file %s is successfully created\n", destinationPath)
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
