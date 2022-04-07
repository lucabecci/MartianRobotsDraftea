package cardenal

var Cardenal = []string{
	"N",
	"E",
	"S",
	"W",
}

func FindIndex(cardenalToFind string) int {
	for i := 0; i < len(Cardenal); i++ {
		if Cardenal[i] == cardenalToFind {
			return i
		}
	}
	return -1
}
