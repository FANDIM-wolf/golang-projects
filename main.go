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

//dictionary  to keep data 
Dictionary :=map[string]string{
	"login": "",
	"password":"",

}
//dictionary of basket 
Dictionary_of_basket:=map[string]int{
	"price":0,
}
func remove_all_from_basket(){
	Dictionary_of_basket["price"]=0
	fmt.Printf("Basket is clear")
}

func main(){
	//variables
	var type_of_drink string 

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
		Dictionary_of_basket["price"] += 45 
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
		Dictionary_of_basket["price"] += 55 
		Dictionary["login"]="none"
		fmt.Println("Drink:",coconut_coctail.name)
		fmt.Println("Price: ",coconut_coctail.price)
		fmt.Printf("Thanks for %s",type_of_drink)
	}
	if type_of_drink == "clear" || type_of_drink == "none" {
		remove_all_from_basket()
	}
	
	fmt.Printf("bill for service of resturant: %d\n",Dictionary_of_basket["price"])
}