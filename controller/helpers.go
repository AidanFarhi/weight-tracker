package controller

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"strings"
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

func IsValidEmail(email string) bool {
	s := strings.TrimSpace(email)
	if s == "" {
		return false
	}
	addr, err := mail.ParseAddress(s)
	if err != nil {
		return false
	}
	return strings.EqualFold(addr.Address, s)
}

func IsValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 20 {
		return false
	}
	nonValidChars := regexp.MustCompile(`[\s]`)
	if nonValidChars.MatchString(password) {
		return false
	}
	specialChars := regexp.MustCompile(`[!@#\$%\^&\*]`)
	return specialChars.MatchString(password)
}
