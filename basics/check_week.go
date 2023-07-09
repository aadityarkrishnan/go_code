package main

import "fmt"

func main(){
	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error reading number:", err)
		return
	}
	fmt.Println("You entered:", number)

	if number == 1 {
		fmt.Print("Sunday")
	}	else if number == 2 {
		fmt.Print("Monday")
	}	else if number == 3 { 
		fmt.Print("Tuesday")
	}  else if number == 4 {
		fmt.Print("Wednesday")
	} else if number == 5 {
		fmt.Print("Thursday")
	} else if number == 6 {
		fmt.Print("Friday")
	} else if number == 7 {
		fmt.Print("Saturday")
	} else {
		fmt.Print("Invalid")
	}
		
	
	
	
	
	
	


}