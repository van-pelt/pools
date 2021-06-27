package template

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Templ = func() *template.Template {
	t := template.New("")
	basePath, err := os.Getwd()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	err = filepath.Walk(basePath+"/web", func(path string, info os.FileInfo, err error) error {
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
	Basepath string
}
