package dungeon

import (
	"log"
	"math/rand"
	"time"
)

/*
setSeed défini la seed pour toute les opérations aléatoire pour la génération du donjon à partir du string donné. 
*/
func setSeed(seed string) {
	if seed == "random" || seed == "" {
		alea = rand.New(rand.NewSource(time.Now().UnixNano()))
	} else if seed == "fromCurrentSeed" {
		createNewSeed()
	} else {
		// seed conversion
		seedRune := []rune(seed)
		seedInt := 0
		for i := 0; i < len(seedRune); i++ {
			currentSeedChar := int(seedRune[i])
			seedInt += currentSeedChar
		}
		alea = rand.New(rand.NewSource(int64(seedInt)))
	}
	// log.Println(alea.Float64())
}

/*
Créer une nouvelle seed à partir de la seed actuellement utilisé (set à la seed déterministe)
*/
func createNewSeed() {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := make([]byte, alea.Intn(256))
	for i := range seed {
		seed[i] = charset[alea.Intn(len(charset))]
	}
	setSeed(string(seed))
}

/*
randomIntBetween choisie aléatoirement un entier entre les deux bornes données, incluse.
*/
func randomIntBetween(min, max int) (res int) {
	if min > max {
		log.Panicf("error : min %d > max %d", min, max)
	} else if min == max {
		return min
	}
	res = alea.Intn(max-min+1) + min
	return res
}
