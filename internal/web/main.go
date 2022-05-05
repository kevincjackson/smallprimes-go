package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/kevincjackson/smallprimes-go/internal/format"
	"github.com/kevincjackson/smallprimes-go/pkg/primedata"
)

const host = "localhost:3001"

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	log.Printf("Web server runnung on %s\n", host)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/is", isHandler)
	http.HandleFunc("/upto", uptoHandler)
	http.HandleFunc("/between", betweenHandler)

	http.ListenAndServe(host, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server OK")
}

func isHandler(w http.ResponseWriter, r *http.Request) {
	var res XData
	xstr := r.URL.Query().Get("x")
	if xstr == "" {
		res = XData{"", "", ""}
	} else {
		xint, err := strconv.Atoi(xstr)
		if err != nil {
			res = XData{xstr, "", "Ooops. Not a number."}
		} else if xint > primedata.MaxInt {
			res = XData{xstr, "", "Ooops. Number too large."}
		} else {
			xres := primedata.Is(xint)
			res = XData{xstr, strconv.FormatBool(xres), ""}
		}
	}
	err := t.ExecuteTemplate(w, "is.gohtml", res)
	if err != nil {
		log.Fatal(err)
	}
}

func uptoHandler(w http.ResponseWriter, r *http.Request) {
	var res XData
	xstr := r.URL.Query().Get("x")
	if xstr == "" {
		res = XData{"", "", ""}
	} else {
		xint, err := strconv.Atoi(xstr)
		if err != nil {
			res = XData{xstr, "", "Ooops. Not a number."}
		} else if xint > primedata.MaxInt {
			res = XData{xstr, "", "Ooops. Number too large."}
		} else {
			xres := "UPTO RESULT"
			res = XData{xstr, xres, ""}
		}
	}
	err := t.ExecuteTemplate(w, "upto.gohtml", res)
	if err != nil {
		log.Fatal(err)
	}
}

func betweenHandler(w http.ResponseWriter, r *http.Request) {
	var res XyData
	xstr := r.URL.Query().Get("x")
	ystr := r.URL.Query().Get("y")
	if xstr == "" && ystr == "" {
		res = XyData{"", "", "", ""}
	} else {
		xint, xerr := strconv.Atoi(xstr)
		yint, yerr := strconv.Atoi(ystr)
		if xerr != nil || yerr != nil {
			res = XyData{xstr, ystr, "", "Ooops. Not a number."}
		} else if xint > primedata.MaxInt || yint > primedata.MaxInt {
			res = XyData{xstr, ystr, "", "Ooops. Number too large."}
		} else {
			if xint > yint {
				xint, yint = yint, xint
			}
			xyres := format.Between(xint, yint, "", ", ", "")
			res = XyData{xstr, ystr, xyres, ""}
		}
	}
	err := t.ExecuteTemplate(w, "between.gohtml", res)
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
