package cardenal

// Variable for available cardinal points
var Cardenal = []string{
	"N",
	"E",
	"S",
	"W",
}

// Gets the index of the received cardinal point
// Return an int from index
func FindIndex(cardenalToFind string) int {
	for i := 0; i < len(Cardenal); i++ {
		if Cardenal[i] == cardenalToFind {
			return i
		}
	}
	return -1
}
