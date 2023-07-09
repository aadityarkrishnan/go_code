// Go program to illustrate the
// use of function
package main
import "fmt"

// area() is used to find the
// area of the rectangle
// area() function two parameters,
// i.e, length and width
func area(length, width int)int{
	
	Ar := length* width
	return Ar
}

// function which swap values
func swap(a, b int)int{
 
    var o int
    o= a
    a=b
    b=o
    
   return o
}

func swap_by_ref(a, b *int)int{
    var o int
    o = *a
    *a = *b
    *b = o
     
   return o
}


// Main function
func main() {

// Display the area of the rectangle
// with method calling
fmt.Printf("Area of rectangle is : %d", area(12, 10))
var p int = 10
var q int = 20
 fmt.Printf("p = %d and q = %d", p, q)
 
// call by values
swap(p, q)

  fmt.Printf("\np = %d and q = %d",p, q)
  swap_by_ref(&p, &q)
  fmt.Printf("\np = %d and q = %d",p, q)
}
