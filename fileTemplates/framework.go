package fileTemplates

var Options []string = []string{"net/http", "chi", "gin", "echo", "fiber"}

var FrameWorkMap map[string]*FrameWork = map[string]*FrameWork{"net/http": NetHttpInit()}

// create a base class, and then using class composition to behaive differently
type FrameWork struct {
	Main     Main
	Handler  Handler
	Server   Server
	Template Template
}

func (f *FrameWork) StartCreate(moduleName string, folderPath string) error {
	addr, content := f.Main.MainContent(moduleName)
	err := CreateFoldersAndFiles(moduleName, folderPath, addr, content)
	if err != nil {
		return err
	}
	addr, content = f.Handler.HandlerContent()
	err = CreateFoldersAndFiles(moduleName, folderPath, addr, content)
	if err != nil {
		return err
	}

	addr, content = f.Server.serverContent(moduleName)
	err = CreateFoldersAndFiles(moduleName, folderPath, addr, content)
	if err != nil {
		return err
	}

	addr, content = f.Template.TemplateContent()
	err = CreateFoldersAndFiles(moduleName, folderPath, addr, content)
	if err != nil {
		return err
	}

	return nil
}

func NetHttpInit() *FrameWork {
	return &FrameWork{Main: &NetHttpMain{}, Handler: &NetHttpHandler{}, Server: &NetHttpServer{}, Template: &NetHttpTemplate{}}
}
