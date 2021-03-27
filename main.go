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

	//dictionary  to keep data 
	Dictionary :=map[string]string{
		"login": "",
		"password":"",
	}

	//print title
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
		Dictionary["login"]="none"
		fmt.Println("Drink:",coffe.name)
		fmt.Println("Price: ",coffe.price)
		fmt.Printf("Thanks for %s",type_of_drink)
	}
	
}