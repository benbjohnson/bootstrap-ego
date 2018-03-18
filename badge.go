package bootstrap

import (
	"context"
	"io"
)

type Badge struct {
	Class string
	Style string
	Pill  bool
	Yield func()
}

func (c *Badge) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<span role="badge"`)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</span>`)
}

func (c *Badge) renderClass(ctx context.Context, w io.Writer) {
	// Write style class.
	io.WriteString(w, ` class="badge`)
	switch c.Style {
	case "primary", "secondary", "success", "danger", "warning", "info", "light", "dark":
		io.WriteString(w, ` badge-`)
		io.WriteString(w, c.Style)
	case "":
		Logger.Printf("bootstrap.Badge: Style required")
	default:
		Logger.Printf("bootstrap.Badge: Invalid style: %q", c.Style)
	}

	// Pill style
	if c.Pill {
		io.WriteString(w, ` badge-pill`)
	}

	// Write user-defined class.
	if c.Class != "" {
		io.WriteString(w, ` `)
		io.WriteString(w, c.Class)
	}

	io.WriteString(w, `"`)
}
