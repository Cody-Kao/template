package fileTemplates

import (
	"fmt"
	"os"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFolders(destinationPath string) error {
	exist, err := PathExists(destinationPath)
	if err != nil {
		return err
	}
	if !exist {
		err = CreateFolders(filepath.Dir(destinationPath)) // create parent folders recursively
		if err != nil {
			return err
		}
	} else {
		return nil
	}

	err = os.Mkdir(destinationPath, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("parent path %s is successfully created\n", destinationPath)
	return nil
}

func CreateFiles(destinationPath string, content string) error {
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

func CreateFoldersAndFiles(moduleName string, folderPath string, addr string, content string) error {
	destinationPath := filepath.Join(folderPath, addr)
	// filepath.Dir會把path的最後一項去掉，只保留前段，等於是創造parent directory的意思
	parentPath := filepath.Dir(destinationPath)
	err := CreateFolders(parentPath)
	if err != nil {
		return err
	}
	err = CreateFiles(destinationPath, content)
	if err != nil {
		return err
	}

	return nil
}
