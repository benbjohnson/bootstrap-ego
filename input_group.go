package bootstrap

import (
	"context"
	"io"
)

type InputGroup struct {
	ID      string
	Class   string
	Size    string
	Prepend func()
	Append  func()
	Yield   func()
}

func (c *InputGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Prepend != nil {
		io.WriteString(w, `<div class="input-group-prepend">`)
		c.Prepend()
		io.WriteString(w, `</div>`)
	}

	if c.Yield != nil {
		c.Yield()
	}

	if c.Append != nil {
		io.WriteString(w, `<div class="input-group-append">`)
		c.Append()
		io.WriteString(w, `</div>`)
	}

	io.WriteString(w, `</div>`)
}

func (c *InputGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="input-group`)

	switch c.Size {
	case "sm", "lg":
		io.WriteString(w, ` input-group-`)
		io.WriteString(w, c.Size)
	case "":
	default:
		Logger.Printf("bootstrap.InputGroup: Invalid size: %q", c.Size)
	}

	appendHTML(w, c.Class)
	io.WriteString(w, `"`)
}
