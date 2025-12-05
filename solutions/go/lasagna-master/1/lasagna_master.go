package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer <= 0 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "sauce" {
			sauce += 0.2
		} else if layer == "noodles" {
			noodles += 50
		}
	}
	return
}

func AddSecretIngredient(friendList, myList []string) {
	myList = append(myList[:len(myList)-1], friendList[len(friendList)-1])
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	newValues := make([]float64, len(amounts))
	for ind, value := range amounts {
		newValues[ind] = (value / 2.0) * float64(portions)
	}
	return newValues
}
