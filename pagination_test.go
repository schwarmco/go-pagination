package pagination

import (
	"fmt"
)

func ExampleCalculate() {

	currentPage := 2
	itemsPerPage := 10
	numItems := 42

	pager := Calculate(currentPage, itemsPerPage, numItems)

	fmt.Println("NumPages:", pager.NumPages)
	fmt.Println("HasPrev:", pager.HasPrev)
	fmt.Println("PrevPage:", pager.PrevPage)
	fmt.Println("HasNext:", pager.HasNext)
	fmt.Println("NextPage:", pager.NextPage)

	// Output:
	// NumPages: 5
	// HasPrev: true
	// PrevPage: 1
	// HasNext: true
	// NextPage: 3
}
