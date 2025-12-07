package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	sum := 0
	lowest := 7
	for range 4 {
		roll := rand.Intn(6) + 1
		if roll < lowest {
			lowest = roll
		}
		sum += roll
	}
	sum -= lowest
	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	cons := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: cons,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(cons),
	}
}
