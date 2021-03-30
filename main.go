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


func work(){
	//variables
	var type_of_drink string 
	var continue_to_order string
	var price_in_busket int
	fmt.Scanf("%s\n",&type_of_drink)
	
	//dictionary  to keep data 
	Dictionary:=map[string]string{
		"login": "",
		"password":"",

	}
	

	//condition 
	if type_of_drink == "coffe"{
		coffe := GOOD{
			id :1,
			name : "coffe",
			price: 45,
			additional: []string{"nothing","got"},
		}
		Dictionary["login"]="none"
		price_in_busket =price_in_busket+coffe.price //bill in busket
		fmt.Println("Drink:",coffe.name)
		fmt.Println("Price: ",coffe.price)
		fmt.Printf("Thanks for %s",type_of_drink)

		fmt.Println("Do you wanna order something else:")
		fmt.Scanf("%s\n",&continue_to_order)
		if continue_to_order == "y" || continue_to_order == "yes" {
			work()
		}

	}
	if type_of_drink == "coconut coctail" || type_of_drink == "coconut"{
		coconut_coctail := GOOD{
			id :1,
			name : "coconut coctail",
			price: 55,
			additional: []string{"nothing","got"},
		}
		price_in_busket = price_in_busket+coconut_coctail.price  //bill in busket
		Dictionary["login"]="none"
		fmt.Println("Drink:",coconut_coctail.name)
		fmt.Println("Price: ",coconut_coctail.price)
		fmt.Printf("Thanks for %s",type_of_drink)

		fmt.Println("Do you wanna order something else:")
		fmt.Scanf("%s\n",&continue_to_order)
		if continue_to_order == "y" || continue_to_order == "yes" {
			work()
		}
	}
	if type_of_drink == "clear" || type_of_drink == "none" {
		price_in_busket =0
	}
	
	fmt.Printf("bill for service of resturant: %d\n",price_in_busket)
	}

 
func main(){
	//print title
	fmt.Println("Cafe Versetti\n")

	work()	
	
}