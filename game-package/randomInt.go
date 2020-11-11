package game_package

import (
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

func randomGen(min, max int) int {
	once.Do(func() {
		rand.Seed(time.Now().Unix())
	})
	return rand.Intn(max-min) + min
}
