package templates

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var templ = func() *template.Template {
	t := template.New("")
	err := filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			fmt.Println(path)
			_, err = t.ParseFiles(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}
	return t
}()

type Page struct {
	Title string
}
