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
	"html/template"
	"net/http"
)

const program = "Sylvan"
const author = "David R. Larsen"
const version = "0.0.1"

var update = "2021-09-01"
var units = "centimeters"

const srvaddress = "localhost"
const srvport string = "7854"

// HTML to define the style and header settings
const head string = `
<html>
<head>
<title> Sylvan Forest Growth Model </title>
<meta charset=utf-8 />
`

// HTML to define the style and header settings
const w1style string = `
<style>
html, body, h1, div, select {
  }
  body {
    background: #ffffff;
    color: #000000;
    font-family: Helvetica, Geneva, Arial, sans-serif;
    padding: 20px;
  }
  @-viewport {
    width: 640px;
  }
  h1 {
    font-size: 28px;
    margin-bottom: 20px;
  }
  select, input, button {
		  display:block;
		  border-radius:8px;
		  -moz-border-radius: 8;
		  -webkit-border-radius: 8px;
		  border: 5px solid Darkgreen;
		  height: 35px;
		  font-size: 20px;
   }
	table {
	  empty-cells: show
   }
    .hdr {
			display:block;
			background: #eeeeee;
			border-radius:8px;
			-moz-border-radius: 8;
			-webkit-border-radius: 8px;
		   border: 2px solid black;
			height: 35px;
			font-size: 20px;
	 }
	 .nic1 {
		   font-size: 20px;
    }
	 .mac1 {
	   	font-size: 20px;
	 }
	 .btn {
         display:block;
			align:center;
			valign:bottom;
         border-radius:8px;
         -moz-border-radius: 8;
         -webkit-border-radius: 8px;
         border: 5px solid Darkgreen;
         height: 35px;
			width: 200px;
         font-size: 20p
    }
	 .intp {
         display:block;
			align:center;
			valign:bottom;
         border-radius:6px;
         -moz-border-radius: 6;
         -webkit-border-radius: 6px;
         border: 3px solid Darkgreen;
         height: 35px;
			width: 70;
         font-size: 16
    }
	 .fileintp {
         display:block;
			align:center;
			valign:bottom;
         border-radius:6px;
         -moz-border-radius: 6;
         -webkit-border-radius: 6px;
         border: 3px solid Darkgreen;
         height: 35px;
			width: 600;
         font-size: 16
    }
}
</style>
`

// HTML to define the style and header settings
const body string = `
</head>
<body>
`

const banner string = `
<div id="header" align="left" style="background-color:Darkgreen;border:2px solid; border-radius:8px; ">
<b align="left" width="90%" style="margin-left:45px; color:white; font-size:38px ">Sylvan</b>
<p style="margin-left:45px; color:white; font-size:12 ">By David R. Larsen - Version {{.Version}}, Protocol {{.Protocol}} - Last Updated {{.Update}} -  <a style="color:white; font-size:12px" href="http://oakmissouri.org">oakmissouri.org</a> </p>
</div>
`

const intro string = `
<h2>Overview</h2>
<p> The Sylvan Forest Growth Model models spatial realationship of trees based on spacing and asymetry.  
`

// structure definition for the Html struct
type Html struct {
	Version  string
	Protocol string
	Update   string
	Address  string
	Port     string
}

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

	var H Html
	H.Version = version
	H.Update = update
	H.Address = srvaddress
	H.Port = srvport

	t, _ := template.New("head").Parse(head)
	t.Execute(w, "head")

	t1, _ := template.New("style").Parse(w1style)
	t1.Execute(w, "style")

	t2, _ := template.New("body").Parse(body)
	t2.Execute(w, "body")

	t3, _ := template.New("webbanner").Parse(banner)
	t3.Execute(w, H)

	t4, _ := template.New("intro").Parse(intro)
	t4.Execute(w, H)

	fmt.Fprintf(w, "Serving webpage at localhost:7854/")

	fmt.Fprintf(w, "</body>\n")
	fmt.Fprintf(w, "</html>\n")

}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	fmt.Printf("%s by %s, Version %s, Units %s, Last Updated %s\n", program, author, version, units, update)
	fmt.Printf("Serving webpage at localhost:7854/")

	http.HandleFunc("/", welcome)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":7854", nil)

}
