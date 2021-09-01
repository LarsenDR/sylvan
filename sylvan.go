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

import "fmt"

var Program string
var Author string
var Version string
var Units string

type Section struct {
	Azi   float64
	GR    float64
	GHt   float64
	BHRad float64
	BHHt  float64
	Tht   float64
}

type Tree struct {
	Id   int64
	X    float64
	Y    float64
	Elev float64
	Spp  string
	//Section []slice
}

func main() {
	Program = "Sylvan"
	Author = "David R. Larsen"
	Version = "0.0.1"
	Units = "centimeters"

	fmt.Printf("%s by %s\n", Program, Author)
	fmt.Printf("Version %s, Units %s\n", Version, Units)

}
