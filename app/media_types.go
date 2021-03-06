// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Algorithms": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/schiob/AAAS/design
// --out=$(GOPATH)/src/github.com/schiob/AAAS
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// A graph result (default view)
//
// Identifier: application/vnd.graphresult+json; view=default
type Graphresult struct {
	// Father of every node
	Fathers []int `form:"fathers" json:"fathers" xml:"fathers"`
	// Resut for every node
	Shortest []int `form:"shortest" json:"shortest" xml:"shortest"`
	// ID of start point
	Start int `form:"start" json:"start" xml:"start"`
}

// Validate validates the Graphresult media type instance.
func (mt *Graphresult) Validate() (err error) {

	if mt.Shortest == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "shortest"))
	}
	if mt.Fathers == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "fathers"))
	}
	return
}
