package membersClub

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
	//tmpl := template.Must(template.ParseFiles("forms.html"))
	tmpl, err := template.New("forms.html").Funcs(funcMap).ParseFiles("forms.html")
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
