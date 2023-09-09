package goroutine

import (
	"math/rand"
	"time"
)

func ReallyLongTask() {
	<-time.After(time.Millisecond * time.Duration(rand.Int63n(500)))
}
