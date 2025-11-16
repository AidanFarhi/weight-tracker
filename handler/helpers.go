package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderPage(w http.ResponseWriter, page string, data any) error {
	t, err := template.ParseFiles(
		"./web/templates/base.html",
		fmt.Sprintf("./web/templates/pages/%s.html", page),
	)
	if err != nil {
		return err
	}
	err = t.ExecuteTemplate(w, page, data)
	if err != nil {
		return err
	}
	return nil
}
