package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Badge struct {
	ID    string
	Class string
	Style string
	Pill  bool
	Yield func()
}

func (r *Badge) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<span role="badge"`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</span>`)
}

func (r *Badge) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="badge`)

	// Write style class.
	switch r.Style {
	case "primary", "secondary", "success", "danger", "warning", "info", "light", "dark":
		fmt.Fprintf(w, ` badge-%s`, r.Style)
	case "":
		Logger.Printf("bootstrap.Badge: Style required")
	default:
		Logger.Printf("bootstrap.Badge: Invalid style: %q", r.Style)
	}

	// Pill style
	if r.Pill {
		io.WriteString(w, ` badge-pill`)
	}

	// Write user-defined class.
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
