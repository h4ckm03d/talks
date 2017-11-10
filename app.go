package main

import (
	"mime"
	"os"

	"golang.org/x/tools/present"
)

func main() {

	basePath := os.Getenv("TPL_DIR")
	if basePath == "" {
		basePath = "."
	}

	initTemplates(basePath)

	present.PlayEnabled = true
	initPlayground(basePath, nil)

	// App Engine has no /etc/mime.types
	mime.AddExtensionType(".svg", "image/svg+xml")
}
