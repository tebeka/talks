package main

// meanwhile at stdlib.h:
// typedef struct
//   {
//     int quot;     /* Quotient.  */
//     int rem;      /* Remainder.  */
//   } div_t;

// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	var dt C.div_t  // Look Ma! C types

	dt = C.div(16, 6)
	fmt.Printf("quot: %d, rem: %d\n", dt.quot, dt.rem)

	dt.quot = 17 // Mutate the struct
	fmt.Printf("%v\n", dt) // %v works as well
}

