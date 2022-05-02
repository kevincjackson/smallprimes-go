package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kevincjackson/smallprimes-go/pkg/primedata"
)

const host = "localhost:3000"

func main() {
	fmt.Printf("API server running on %s\n", host)

	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/is", isHandler)
	http.HandleFunc("/api/upto", uptoHandler)
	http.HandleFunc("/api/between", betweenHandler)

	http.ListenAndServe(host, nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func isHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	x := r.URL.Query().Get("x")
	if x == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Expecting x in query. Got empty x.")
	} else {
		xint, err := strconv.Atoi(x)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Expecting integer for x, got non-integer.")
		} else {
			res := primedata.Is(xint)
			json, _ := json.Marshal(res)
			w.Write(json)
		}
	}
}

func uptoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	x := r.URL.Query().Get("x")
	if x == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Expecting x in query. Got empty x.")
	} else {
		xint, err := strconv.Atoi(x)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Expecting integer x, got non-integer.")
		} else {
			res := primedata.Upto(xint)
			json, _ := json.Marshal(res)
			w.Write(json)
		}
	}
}

func betweenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	x, y := r.URL.Query().Get("x"), r.URL.Query().Get("y")
	if x == "" || y == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Expecting x and y in query.")
	} else {
		xint, xerr := strconv.Atoi(x)
		yint, yerr := strconv.Atoi(y)
		if xerr != nil || yerr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Expecting integers for x and y.")
		} else {
			res := primedata.Between(xint, yint)
			json, _ := json.Marshal(res)
			w.Write(json)
		}
	}
}
