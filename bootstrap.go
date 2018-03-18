package bootstrap

import (
	"html"
	"io"
	"log"
	"os"
)

// Logger is the default logger for reporting render errors.
var Logger = log.New(os.Stderr, "", 0)

// appendAttr appends a space and key/value pair to w if value is non-blank.
func appendAttr(w io.Writer, key, value string) {
	if value == "" {
		return
	}
	io.WriteString(w, ` `)
	io.WriteString(w, key)
	io.WriteString(w, `="`)
	io.WriteString(w, html.EscapeString(value))
	io.WriteString(w, `"`)
}

// appendHTML appends a space and escaped HTML value to w if s is non-blank.
func appendHTML(w io.Writer, s string) {
	if s == "" {
		return
	}
	io.WriteString(w, ` `)
	io.WriteString(w, html.EscapeString(s))
}
