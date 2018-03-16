package bootstrap

import (
	"log"
	"os"
)

// Logger is the default logger for reporting render errors.
var Logger = log.New(os.Stderr, "", 0)
