package main 

import (
	"fmt"
)

type GOOD struct {
	id int64
	name string 
	price int 
	additional [] string
}

func main(){
	//variables
	var type_of_drink string 

	fmt.Println("Cafe Versetti\n")
	fmt.Scanf("%s\n",&type_of_drink)
	
	//condition 
	if type_of_drink == "coffe"{
		coffe := GOOD{
			id :1,
			name : "coffe",
			price: 45,
			additional: []string{"nothing","got"},
		}
		fmt.Println(coffe.price)
		fmt.Printf("Thanks for %s",type_of_drink)
	}
	
}