package main

import (
	"os"
	"time"

	"github.com/ablqk/tiny-go-projects/chapter-03/exercises/7_avoiding_long_messages/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))

	lgr.Infof("A little copying is better than a little dependency.")
	lgr.Errorf("Errors are values. Documentation is for %s.", "users")
	lgr.Debugf("Make the zero (%d) value useful.", 0)

	lgr.Infof("Hallo, %d %v", 2022, time.Now())
}
