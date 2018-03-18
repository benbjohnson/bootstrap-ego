package bootstrap

import (
	"context"
	"html"
	"io"
)

type CloseButton struct {
	Class   string
	Dismiss string
}

func (c *CloseButton) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<button type="button"`)
	c.renderClass(ctx, w)

	// Render 'dismiss' data.
	io.WriteString(w, ` data-dismiss="`)
	if c.Dismiss != "" {
		io.WriteString(w, html.EscapeString(c.Dismiss))
	} else {
		io.WriteString(w, "alert")
	}
	io.WriteString(w, `"`)

	// Write accessible label.
	io.WriteString(w, ` aria-label="Close"`)

	// Close tag.
	io.WriteString(w, `><span aria-hidden="true">&times;</span></button>`)
}

func (c *CloseButton) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="close`)
	if c.Class != "" {
		io.WriteString(w, ` `)
		io.WriteString(w, c.Class)
	}
	io.WriteString(w, `"`)
}
