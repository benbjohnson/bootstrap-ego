package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Alert struct {
	ID          string
	Class       string
	Style       string
	Dismissable bool
	Yield       func()
}

func (r *Alert) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="alert"`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *Alert) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="alert`)

	// Write style class.
	switch r.Style {
	case "primary", "secondary", "success", "danger", "warning", "info", "light", "dark":
		fmt.Fprintf(w, ` alert-%s`, r.Style)
	case "":
		Logger.Printf("bootstrap.Alert: Style required")
	default:
		Logger.Printf("bootstrap.Alert: Invalid style: %q", r.Style)
	}

	// Write dismissable class.
	if r.Dismissable {
		io.WriteString(w, ` alert-dismissable`)
	}

	// Write user-defined class.
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
