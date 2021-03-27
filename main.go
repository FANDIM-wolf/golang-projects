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
	if type_of_drink == "coconut coctail" || type_of_drink == "coconut"{
		coconut_coctail := GOOD{
			id :1,
			name : "coconut coctail",
			price: 55,
			additional: []string{"nothing","got"},
		}
		Dictionary["login"]="none"
		fmt.Println("Drink:",coconut_coctail.name)
		fmt.Println("Price: ",coconut_coctail.price)
		fmt.Printf("Thanks for %s",type_of_drink)
	}
}