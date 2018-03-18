package bootstrap

import (
	"context"
	"io"
)

type ButtonGroup struct {
	ID       string
	Class    string
	Size     string
	Vertical bool
	Yield    func()
}

func (c *ButtonGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="group"`)
	appendAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *ButtonGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="btn-group`)

	// Write size class.
	switch c.Size {
	case "sm", "lg":
		io.WriteString(w, ` btn-group-`)
		io.WriteString(w, c.Size)
	case "":
	default:
		Logger.Printf("bootstrap.ButtonGroup: Invalid size: %q", c.Size)
	}

	// Enable vertical group.
	if c.Vertical {
		io.WriteString(w, ` btn-group-vertical`)
	}

	// Write user-defined class.
	appendHTML(w, c.Class)

	io.WriteString(w, `"`)
}
