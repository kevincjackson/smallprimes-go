package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/kevincjackson/smallprimes-go/pkg/primedata"
)

const host = "localhost:3001"

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	fmt.Printf("Web server runnung on %s\n", host)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/is", isHandler)

	http.ListenAndServe(host, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server OK")
}

func isHandler(w http.ResponseWriter, r *http.Request) {
	var res XData
	var err error
	xstr := r.URL.Query().Get("x")
	if xstr == "" {
		res = XData{"", "", ""}
	} else {
		xint, err := strconv.Atoi(xstr)
		if err != nil {
			res = XData{xstr, "", "Ooops. Input isn't a number."}
		} else {
			if xint > primedata.MaxInt {
				res = XData{xstr, "", "Ooops. Input is too large. "}
			} else {
				xres := primedata.Is(xint)
				res = XData{xstr, strconv.FormatBool(xres), ""}
			}
		}
	}
	err = t.ExecuteTemplate(w, "is.gohtml", res)
	if err != nil {
		log.Fatal(err)
	}
}

type XData struct {
	X      string
	Result string
	Error  string
}

type XyData struct {
	X      string
	Y      string
	Result string
	Error  string
}
