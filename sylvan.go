// A new version of the Sylvan growth model using a 3-D data structure.
//
// By David R. Larsen
// Professor Emeritus
// The School of Natural Resources
// University of Missouri
//
// September 1, 2021
//

package main

import (
	"fmt"
	"net/http"
)

var Program string
var Author string
var Version string
var Units string

type Section struct {
	Azi float64
	GR  int64
	BHR int64
	BHt int64
	Tht int64
}

type Tree struct {
	Id    int64
	GX    int64
	GY    int64
	GElev int64
	Spp   string
	//Section []slice
}

func welcome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	Program = "Sylvan"
	Author = "David R. Larsen"
	Version = "0.0.1"
	Units = "centimeters"

	//fmt.Printf("%s by %s, Version %s, Units %s\n", Program, Author, Version, Units)

	http.HandleFunc("/hello", welcome)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)

}
