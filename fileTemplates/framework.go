package fileTemplates

var Options []string = []string{"net/http", "chi", "gin", "echo", "fiber"}

var FrameWorkMap map[string]*FrameWork = map[string]*FrameWork{"net/http": NetHttpInit()}

// create a base class, and then instantiate differently to behaive differently
type FrameWork struct {
	Main     Main
	Handler  Handler
	Server   Server
	Template Template
	Static   Static
}

func (f *FrameWork) StartCreate(moduleName string, folderPath string) error {
	addrSlice, contentSlice := f.Main.MainContent(moduleName)
	err := CreateFoldersAndFiles(moduleName, folderPath, addrSlice, contentSlice)
	if err != nil {
		return err
	}

	addrSlice, contentSlice = f.Handler.HandlerContent()
	err = CreateFoldersAndFiles(moduleName, folderPath, addrSlice, contentSlice)
	if err != nil {
		return err
	}

	addrSlice, contentSlice = f.Server.ServerContent(moduleName)
	err = CreateFoldersAndFiles(moduleName, folderPath, addrSlice, contentSlice)
	if err != nil {
		return err
	}

	addrSlice, contentSlice = f.Template.TemplateContent()
	err = CreateFoldersAndFiles(moduleName, folderPath, addrSlice, contentSlice)
	if err != nil {
		return err
	}

	addrSlice, contentSlice = f.Static.StaticContent()
	err = CreateFoldersAndFiles(moduleName, folderPath, addrSlice, contentSlice)
	if err != nil {
		return err
	}

	return nil
}

// 日後要新增模板就去現有的檔案新增模板樣式，或是自己創新的檔案去新增
// 再從這裡創造init function，並跟framework map連結
func NetHttpInit() *FrameWork {
	return &FrameWork{Main: &NetHttpMain{}, Handler: &NetHttpHandler{}, Server: &NetHttpServer{}, Template: &NetHttpTemplate{}, Static: &DefaultStatic{}}
}
