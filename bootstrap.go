package bootstrap

import (
	"html"
	"io"
	"log"
	"os"
)

// Logger is the default logger for reporting render errors.
var Logger = log.New(os.Stderr, "", 0)

// writeAttr appends a space and key/value pair to w if value is non-blank.
func writeAttr(w io.Writer, key, value string) {
	if value == "" {
		return
	}
	io.WriteString(w, ` `)
	io.WriteString(w, key)
	io.WriteString(w, `="`)
	io.WriteString(w, html.EscapeString(value))
	io.WriteString(w, `"`)
}
