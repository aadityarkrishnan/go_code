package main 

import "fmt"

func main(){
	for i :=0; i<4; i++ {
		fmt.Printf("Aaditya R Krishnan\n")
	}
	// for {
	// 	fmt.Printf("Infinte Loop\n")  
	// }
	z:= 0
    for z < 3 {
       z += 2
    }
	rvariable:= []string{"GFG", "Geeks", "GeeksforGeeks"} 
      
    // i and j stores the value of rvariable
    // i store index number of individual string and
    // j store individual string of the given array
    for i, j:= range rvariable {
       fmt.Println(i, j) 
    }

	for i, j:= range "XabCd" {
		fmt.Printf("The index number of %U is %d\n", j, i) 
	 }
	
	mmap := map[int]string{
        22:"Geeks",
        33:"GFG",
        44:"GeeksforGeeks",
    }
    for key, value:= range mmap {
       fmt.Println(key, value) 
    }

	chnl := make(chan int)
    go func(){
        chnl <- 100
        chnl <- 1000
       chnl <- 10000
       chnl <- 100000
       close(chnl)
    }()
    for i:= range chnl {
       fmt.Println(i) 
    }
}