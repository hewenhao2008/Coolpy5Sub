package WebSite

import (
	"net/http"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("index")
	t, _ = t.ParseFiles("temp/index.html")
	t.Execute(w, nil)
}

type Home struct {
	Header

}

type Header struct {
	Uname string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		s1, _ := template.ParseFiles("temp/home.html", "temp/homeheader.html", "temp/homefooter.html")
		home := &Home{ Header{ Uname:getUserName(r)}}
		s1.ExecuteTemplate(w, "home", home)
		s1.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
