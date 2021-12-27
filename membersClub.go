package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"net/mail"
	"regexp"
	"time"
)

type Members struct {
	Name  string
	Email string
	Date  string
}

var static string = "<!DOCTYPE html>\n<html>\n<head>\n    <style type=\"text/css\">\n        table {\n            border-collapse: collapse;\n        }\n        table th,\n        table td {\n            padding: 0 3px;\n        }\n        table.brd th,\n        table.brd td {\n            border: 1px solid #000;\n        }\n    </style>\n</head>\n<body>\n<div\n        style=\"\n      display: flex;\n      justify-content: center;\n      margin: 2rem;\n      font-weight: bold;\n      font-size: 2rem;\n    \"\n>\n    Welcome to the club!\n</div>\n<div style=\"display: flex; flex-direction: column; width: 10.6rem\">\n    <div style=\"font-weight: bold; font-size: 1.3rem; margin-bottom: 5px\">\n        New member\n    </div>\n    <form method=\"POST\">\n        <label>Subject:</label><br />\n        <input type=\"text\" name=\"Name\"><br />\n        <label>Email:</label><br />\n        <input type=\"text\" name=\"Email\"><br />\n        <input type=\"submit\">\n    </form>\n</div>\n<div style=\"display: flex; justify-content: center; font-weight: bold; font-size: 1.3rem; margin-top: 2rem;\">\n    Members\n</div>\n<div style=\"display: flex; justify-content: space-around; margin-top: 15px\">\n    <table class=\"brd\">\n        <tr>\n            <th>#</th>\n            <th>Name</th>\n            <th>Email</th>\n            <th>Registration date</th>\n        </tr>\n        {{ range $i, $e := . }}\n        {{ $i := inc $i }}\n            <tr>\n            <td>{{ $i }}</td>\n            <td>{{ .Name }}</td>\n            <td>{{ .Email }}</td>\n            <td>{{ .Date }}</td>\n            </tr>\n        {{ end }}\n    </table>\n</div>\n</body>\n</html>\n"
var membersList []Members

func validateEmail(Email string) bool {
	_, err := mail.ParseAddress(Email)
	return err == nil
}

func validateName(Name string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z. ]+$", Name); !m {
		return false
	} else {
		return true
	}
}

func addMember(req *http.Request) Members {
	details := Members{
		Name:  req.FormValue("Name"),
		Email: req.FormValue("Email"),
		Date:  time.Now().Format("02.01.2006"),
	}
	return details
}

func logger(req *http.Request) {
	log.Printf(
		"%s\t\t%s\t\t%s\t\t",
		req.Method,
		req.RemoteAddr,
		req.RequestURI,
	)
	response, _ := httputil.DumpRequest(req, true)
	log.Println(fmt.Sprintf("%q", response))
}

var notUniqueFlag bool = false

func main() {
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}
	tmpl, err := template.New("forms.html").Funcs(funcMap).Parse(static)
	if err != nil {
		log.Fatal("Template: ", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch requestType := req.Method; requestType {
		case http.MethodPost:
			if validateEmail(req.FormValue("Email")) && validateName(req.FormValue("Name")) {
				for _, member := range membersList {
					if member.Email == req.FormValue("Email") {
						notUniqueFlag = true
					}
				}
				if notUniqueFlag == false {
					membersList = append(membersList, addMember(req))
				} else {
					log.Println(fmt.Sprintf("%q", "Validation error: Not unique email"))
					notUniqueFlag = false
				}
				if err := tmpl.Execute(w, membersList); err != nil {
					http.Error(w, err.Error(), 500)
				}
			}
			logger(req)
		case http.MethodGet:
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), 500)
			}
			logger(req)
		default:
			log.Println(fmt.Sprintf("%q", "Unsupported request type"))
		}
	})
	fmt.Println("Listening on port: 80")
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
