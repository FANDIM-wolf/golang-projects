package main 

import (
	"fmt"
)

func main(){
	//variables
	var type_of_drink string 

	fmt.Println("Cafe Versetti\n")
	fmt.Scanf("%s\n",&type_of_drink)
	
	//condition 
	if type_of_drink == "coffe"{
		fmt.Printf("Thanks for %s",type_of_drink)
	}
	
}