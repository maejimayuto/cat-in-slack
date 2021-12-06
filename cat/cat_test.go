package cat

import (
	"fmt"
)

func ExampleToS() {
	w := new(Cat)
	w.Url = "https://cdn2.thecatapi.com/images/UYLI_E-SZ.jpg"
	w.Width = 750
	w.Height = 500

	fmt.Println(w.ToS())
	// Output:
	// url: https://cdn2.thecatapi.com/images/UYLI_E-SZ.jpg
	// width: 750
	// height: 500
}
