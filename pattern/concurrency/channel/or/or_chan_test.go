package or_test

import (
	"testing"
	"time"

	"github.com/onee-only/go-lab/pattern/concurrency/channel/or"
)

func TestOr(t *testing.T) {
	start := time.Now()

	<-or.Or(
		time.After(time.Second),
		time.After(time.Second*2),
		time.After(time.Minute),
	)

	t.Logf("finished after: %v", time.Since(start))
}
