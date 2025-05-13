package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_When_Krogen_And_19_And_Not_Drunk_I_Should_Be_Allowed(t *testing.T) {
	// Arrange
	location := "K"
	age := 19
	promille := 0.0

	//Act
	//canBuy, err := CanBuyBeer(location, age, float32(promille))
	canBuy, err := CanBuyBeer(location, age, float32(promille))

	//Assert
	assert.True(t, canBuy)
	assert.Nil(t, err)
	// if canBuy == false || err != nil {
	//t.Fatal("Not allowed but should be")
	// }
}

func Test_When_Krogen_And_17_years_old_is_not_drunk_Then_I_should_not_get_beer(t *testing.T) {
	// Arrange
	location := "K"
	age := 17
	promille := 0.0

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))

	//Assert
	assert.False(t, canBuy)
	assert.Nil(t, err)
	// if canBuy == false || err != nil {
	// 	t.Fatal("Not allowed but should be")
	// }
}

func Test_When_Systemenet_And_20_years_old_is_not_drunk_Then_I_should_get_beer(t *testing.T) {
	// Arrange
	location := "S"
	age := 20
	promille := 0.0

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))

	//Assert
	assert.True(t, canBuy)
	assert.Nil(t, err)
	// if canBuy == false || err != nil {
	// 	t.Fatal("Not allowed but should be")
	// }

}

// PRIO:
//	- först "affärskritisk" kod
// 	- sånt som "går" att testa utan alltför mycket strul
//  - buggar - så dom inte kommer tillbaka
//	- ny kod

// alla varianter
// extra vid  gränsvärden för att fånga >=

// SKriva ett antal automatiska tester
// Köra dessa lokalt
// Ändra i kod och se vad bra med tester
//>>>>> Deploy to cloud (AWS CodeBuild)
// Stefan visar "jag kan inte skriva egna tester i denna kurs" - t.Fatal

//Test_When_Krogen_And_19_And_Not_Drunk_I_Should_Be_Allowed failed
//Test_XXX failed
