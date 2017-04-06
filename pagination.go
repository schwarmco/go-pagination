package pagination

import "math"

type Values struct {
	NumPages           int
	HasPrev, HasNext   bool
	PrevPage, NextPage int
	ItemsPerPage       int
	CurrentPage        int
	NumItems           int
}

// pages start at 1 - not 0
func Calculate(currentPage, itemsPerPage, numItems int) (v Values) {

	v = Values{}

	v.CurrentPage = currentPage
	v.ItemsPerPage = itemsPerPage
	v.NumItems = numItems

	// calculate number of pages
	d := float64(v.NumItems) / float64(v.ItemsPerPage)
	v.NumPages = int(math.Ceil(d))

	// HasPrev, HasNext?
	v.HasPrev = v.CurrentPage > 1
	v.HasNext = v.CurrentPage < v.NumPages

	// calculate them
	if v.HasPrev {
		v.PrevPage = v.CurrentPage - 1
	}
	if v.HasNext {
		v.NextPage = v.CurrentPage + 1
	}

	return
}
