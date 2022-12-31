package restaurantmodel

import (
	"testing"
)

func TestRestaurantCreate_Validate(t *testing.T) {
	dataTest := RestaurantCreate{
		Name: "",
	}

	err := dataTest.Validate()

	if err == ErrNameIsEmpty {
		t.Log("Validate restaurant: passed")
		return
	}

	if err == nil {
		t.Error("Validate restaurant. Input name:", dataTest.Name, ". Expect: ErrNameIsEmpty", "Output: ", err)
		return
	}

	t.Log("Validate restaurant. Input name:", dataTest.Name, ". Expect: ErrNameIsEmpty", "Output: ", err)

}
