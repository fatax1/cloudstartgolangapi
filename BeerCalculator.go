package main

import "errors"

func CanBuyBeer(location string, age int, promille float32) (bool, error) {
	if location != "K" && location != "S" {
		return false, errors.New("Invalid location")
	}
	if age < 1 || age > 150 {
		return false, errors.New("Invalid age")
	}
	// if promille < 0 || promille > 5 {
	// 	return false, errors.New("Invalid promille")
	// }
	// if location == "K" && age >= 18 &&  promille > 1.0  {
	// 	return true, nil
	// }
	// if location == "S" && age >= 20 &&  promille > 1.0  {
	// 	return true, nil
	// }
	// return false, nil

	if promille > 1.0 {
		return false, nil
	}
	if location == "K" && age >= 18 {
		return true, nil
	}
	if location == "S" && age >= 20 {
		return true, nil
	}
	return false, nil
}
