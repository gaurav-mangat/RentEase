package Tenant

import (
	"RentEase/utils"
	"fmt"
)

func searchProperties() {

	var propertyType int
	var state, city , locality, pincode , choice string

	fmt.Println("Enter the type of property to search(1.Flat , 2. House 3. Commercial)")
	fmt.Scan(&propertyType)

	state=utils.ReadInput("Enter the State :")
	city=utils.ReadInput("Enter the city :")
	locality=utils.ReadInput("Enter the locality :")
	pincode=utils.ReadInput("Enter the pincode :")
	choice=utils.ReadInput("Enter s to search or enter f to apply filters :")

	if choice=="s"{

	}else if choice=="f"{
		ApplyFilters(propertyType int)
	}else{
		fmt.Println("Please enter a valid choice :")
	}

}

func ApplyFilters(propertyType int) {
	var priceRange float64
	var bhk int
	if propertyType==1 {

		fmt.Println("Enter the BHK")
		fmt.Scanf("%d",&bhk)
	}
	fmt.Println("Enter the price range :")
	fmt.Scanf("%f", &priceRange)


}


