package views

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func PopulateTemplates() {
	templatesDir := "././templates/"
	var allFiles []string
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Println(err)
		os.Exit(1) // No point in running app if templates aren't read
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	templates, err := template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	HomeTemplate = templates.Lookup("home.html")
}
