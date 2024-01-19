package fileTemplates

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func GetUserDeskTop() (string, error) {
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

func CreateParentFolders(destinationPath string) error {
	exist, err := PathExists(destinationPath)
	if err != nil {
		return err
	}
	if !exist {
		err = CreateParentFolders(filepath.Dir(destinationPath)) // create parent folders recursively
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

func createFolder(destinationPath string) error {
	err := os.Mkdir(destinationPath, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("folder path %s is successfully created\n", destinationPath)

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

	fmt.Printf("file path %s is successfully created\n", destinationPath)
	return nil
}

func CreateFoldersAndFiles(moduleName string, folderPath string, addrSlice []string, contentSlice []string) error {
	addrSliceLen, contentSliceLen := len(addrSlice), len(contentSlice)
	if addrSliceLen != contentSliceLen {
		return fmt.Errorf("lengthes of two slices are not equal")
	}
	for i := 0; i < addrSliceLen; i++ {
		destinationPath := filepath.Join(folderPath, addrSlice[i])
		// filepath.Dir會把path的最後一項去掉，只保留前段，等於是創造parent directory的意思
		parentPath := filepath.Dir(destinationPath)
		err := CreateParentFolders(parentPath)
		if err != nil {
			return err
		}

		// 把所有parent path創造完了之後，最後判斷該路徑是一個folder或file，如果是前者則content內容應該為空字串，反之
		// 一般的方法沒辦法判斷一個non-existing path是否為folder或file，因為都有可能，所以這邊使用這種自定義的方法
		emptyString := ""
		if contentSlice[i] == emptyString {
			err = createFolder(destinationPath)
			if err != nil {
				return err
			}
		} else {
			err = CreateFiles(destinationPath, contentSlice[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
