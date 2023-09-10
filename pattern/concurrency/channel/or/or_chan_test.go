package or_chan_test

import (
	"testing"
	"time"

	or_chan "github.com/onee-only/go-lab/pattern/concurrency/channel/or"
)

func TestOr(t *testing.T) {
	start := time.Now()

	<-or_chan.Or(
		time.After(time.Second),
		time.After(time.Second*2),
		time.After(time.Minute),
	)

	t.Logf("finished after: %v", time.Since(start))
}
