package bootstrap

import (
	"context"
	"io"
)

type Alert struct {
	Class       string
	Style       string
	Dismissable bool
	Yield       func()
}

func (c *Alert) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="alert"`)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *Alert) renderClass(ctx context.Context, w io.Writer) {
	// Write style class.
	io.WriteString(w, ` class="alert`)
	switch c.Style {
	case "primary", "secondary", "success", "danger", "warning", "info", "light", "dark":
		io.WriteString(w, ` alert-`)
		io.WriteString(w, c.Style)
	case "":
		Logger.Printf("bootstrap.Alert: Style required")
	default:
		Logger.Printf("bootstrap.Alert: Invalid style: %q", c.Style)
	}

	// Write dismissable class.
	if c.Dismissable {
		io.WriteString(w, ` alert-dismissable`)
	}

	// Write user-defined class.
	if c.Class != "" {
		io.WriteString(w, ` `)
		io.WriteString(w, c.Class)
	}

	io.WriteString(w, `"`)
}
