package fileTemplates

type Static interface {
	StaticContent() ([]string, []string)
}

type DefaultStatic struct{}

func (d *DefaultStatic) StaticContent() ([]string, []string) {
	return []string{"static"}, []string{""}
}
