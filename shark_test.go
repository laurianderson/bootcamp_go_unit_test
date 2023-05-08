package hunt

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// arrange
	shark := Shark{
		hungry: false,
		tired: true,
		speed: 20,
	}

	prey := Prey{
		name: "fish",
		speed: 10,
	}

	result := fmt.Errorf("cannot hunt, i am really tired")

	// act
	err := shark.Hunt(&prey)

	// assert
	assert.Equal(t, result.Error(), err.Error())

}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	// arrange
	shark := Shark{
		hungry: false,
		tired: false,
		speed: 20,
	}

	prey := Prey{
		name: "",
		speed: 0,
	}

	result := fmt.Errorf("cannot hunt, i am not hungry")

	// act
	err := shark.Hunt(&prey)

	// assert
	assert.Equal(t, result.Error(), err.Error())

}

func TestSharkCannotReachThePrey(t *testing.T) {
	// arrange
	shark := Shark{
		hungry: true,
		tired: false,
		speed: 10,
	}

	prey := Prey{
		name: "fish",
		speed: 20,
	}

	result := fmt.Errorf("could not catch it")

	// act
	err := shark.Hunt(&prey)

	// assert
	assert.Equal(t, result.Error(), err.Error())

}


func TestSharkHuntNilPrey(t *testing.T) {
}
