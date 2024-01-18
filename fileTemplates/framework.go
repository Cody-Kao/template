package fileTemplates

import (
	"fmt"
	"path/filepath"
)

var Options []string = []string{"net/http", "chi", "gin", "echo", "fiber"}

var FrameWorkMap map[string]FrameWork = map[string]FrameWork{"net/http": NetHttpInit()}

type FrameWork interface {
	Create(string, string) error
}

type NetHttp struct {
	Handler Handler
	Server  Server
}

func NetHttpInit() *NetHttp {
	return &NetHttp{Handler: &NetHttpHandler{}, Server: &NetHttpServer{}}
}

func (n *NetHttp) Create(moduleName string, folderPath string) error {
	fmt.Println("NetHttp Creates!")
	addr, content := n.Handler.HandlerContent()
	destinationPath := filepath.Join(folderPath, addr)
	err := createFoldersAndFiles(destinationPath, content)
	if err != nil {
		return err
	}

	addr, content = n.Server.serverContent(moduleName)
	destinationPath = filepath.Join(folderPath, addr)
	err = createFoldersAndFiles(destinationPath, content)
	if err != nil {
		return err
	}

	return nil
}
