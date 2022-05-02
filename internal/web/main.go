package main

import (
	"fmt"
	_ "html/template"
	"net/http"
	_ "os"
	"strconv"
	"text/template"

	"github.com/kevincjackson/smallprimes-go/pkg/primedata"
)

const host = "localhost:3001"

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
	t, _ := template.ParseFiles("is.gohtml")
	xstr := r.URL.Query().Get("x")
	if xstr == "" {
		res = XData{"2", "true", ""}
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
	err = t.Execute(w, res)
	if err != nil {
		panic(err)
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
