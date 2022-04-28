package main

import (
	"html/template"
	_ "net/http"
	"os"
)

func main() {

	/*
		type Person struct{ Name string }
		p1 := Person{"Mr. Jones"}

		t, err := template.New("test").Parse("Congratulation {{.Name}}. You won a $1,000,000!")
		if err != nil {
			panic(err)
		}
	*/

	t, _ := template.ParseFiles("hello.gohtml")
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// fPrintf
			fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		})

		http.ListenAndServe(":8080", nil)
	*/
}
