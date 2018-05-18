package bootstrap

import (
	"context"
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
	// Write style class.
	io.WriteString(w, ` class="alert`)
	switch r.Style {
	case "primary", "secondary", "success", "danger", "warning", "info", "light", "dark":
		io.WriteString(w, ` alert-`)
		io.WriteString(w, r.Style)
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
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
