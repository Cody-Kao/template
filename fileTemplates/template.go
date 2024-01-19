package fileTemplates

type Template interface {
	TemplateContent() ([]string, []string)
}

type NetHttpTemplate struct{}

func (n *NetHttpTemplate) TemplateContent() ([]string, []string) {
	return []string{"templates/index.html"}, []string{`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
</head>
<body>
	<h1>Hello World</h1>
</body>
</html>
	`}
}
