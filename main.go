package main
import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"path"
)

type Profile struct {
	Name string
	Age int
	Hobies []string
}

type xmlProfile struct {
  Name    string
  Hobbies []string `xml:"Hobbies>Hobby"`
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Profile{Name: "Join", Age: 26, Hobies: []string{"Reading Book", "research"}}
	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(js)
}

func handlerServeFile(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("", "index.html")
	http.ServeFile(w, r, fp)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func xmlHandler(w http.ResponseWriter, r *http.Request) {
	profile := xmlProfile{"Alex", []string{"snowboarding", "programming"}}

  	x, err := xml.MarshalIndent(profile, "", "  ")
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
  	}

  	w.Header().Set("Content-Type", "application/xml")
  	w.Write(x)
}
func main() {
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/file", handlerServeFile)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/xml", xmlHandler)
	http.ListenAndServe(":3002", nil)
}