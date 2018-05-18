package bootstrap

import (
	"context"
	"html"
	"io"
)

type Button struct {
	ID       string
	Class    string
	NodeName string
	Name     string
	Type     string
	Style    string
	Href     string
	Form     string
	Value    string
	Target   string
	Size     string // "sm","lg"
	Outline  bool
	Block    bool
	Active   bool
	Disabled bool
	Dropdown bool
	Yield    func()
}

func (c *Button) Render(ctx context.Context, w io.Writer) {
	nodeName := c.NodeName
	if c.Href != "" {
		nodeName = "a"
	} else if nodeName == "" {
		nodeName = "button"
	}

	io.WriteString(w, `<`)
	io.WriteString(w, nodeName)
	if nodeName == "a" {
		io.WriteString(w, ` role="button"`)
	} else {
		c.renderType(ctx, w)
	}
	writeAttr(w, "id", c.ID)
	writeAttr(w, "name", c.Name)
	c.renderClass(ctx, w)
	c.renderHref(ctx, w)
	c.renderForm(ctx, w)
	c.renderTarget(ctx, w)
	c.renderValue(ctx, w)
	c.renderDisabled(ctx, w)

	if c.Dropdown {
		io.WriteString(w, ` data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"`)
	}

	// Only allow yield for a/button.
	if nodeName == "input" {
		io.WriteString(w, `/>`)
	} else {
		io.WriteString(w, `>`)

		if c.Yield != nil {
			c.Yield()
		}

		io.WriteString(w, `</`)
		io.WriteString(w, nodeName)
		io.WriteString(w, `>`)
	}
}

func (c *Button) renderType(ctx context.Context, w io.Writer) {
	switch c.Type {
	case "submit", "reset", "button":
		io.WriteString(w, ` type="`)
		io.WriteString(w, c.Type)
		io.WriteString(w, `"`)
	case "":
		io.WriteString(w, ` type="button"`)
	default:
		Logger.Printf("bootstrap.Button: Invalid type: %q", c.Type)
	}
}

func (c *Button) renderClass(ctx context.Context, w io.Writer) {
	// Write style class.
	io.WriteString(w, ` class="btn`)
	switch c.Style {
	case "primary", "secondary", "success", "danger", "warning", "info", "light", "dark", "link":
		if c.Outline {
			io.WriteString(w, ` btn-outline-`)
		} else {
			io.WriteString(w, ` btn-`)
		}
		io.WriteString(w, c.Style)
	case "":
	default:
		Logger.Printf("bootstrap.Button: Invalid style: %q", c.Style)
	}

	// Set button size.
	switch c.Size {
	case "sm", "lg":
		io.WriteString(w, ` btn-`)
		io.WriteString(w, c.Size)
	case "":
	default:
		Logger.Printf("bootstrap.Button: Invalid size: %q", c.Size)
	}

	// Enable block-level button.
	if c.Block {
		io.WriteString(w, ` btn-block`)
	}

	// Enable active look.
	if c.Active {
		io.WriteString(w, ` active`)
	}

	// Enable dropdown.
	if c.Dropdown {
		io.WriteString(w, ` dropdown-toggle`)
	}

	// Write user-defined class.
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

func (c *Button) renderHref(ctx context.Context, w io.Writer) {
	if c.Href != "" {
		io.WriteString(w, ` href="`)
		io.WriteString(w, html.EscapeString(c.Href))
		io.WriteString(w, `"`)
	}
}

func (c *Button) renderForm(ctx context.Context, w io.Writer) {
	if c.Form != "" {
		io.WriteString(w, ` form="`)
		io.WriteString(w, html.EscapeString(c.Form))
		io.WriteString(w, `"`)
	}
}

func (c *Button) renderTarget(ctx context.Context, w io.Writer) {
	if c.Target != "" {
		io.WriteString(w, ` target="`)
		io.WriteString(w, html.EscapeString(c.Target))
		io.WriteString(w, `"`)
	}
}

func (c *Button) renderValue(ctx context.Context, w io.Writer) {
	if c.Value != "" {
		io.WriteString(w, ` value="`)
		io.WriteString(w, html.EscapeString(c.Value))
		io.WriteString(w, `"`)
	}
}

func (c *Button) renderDisabled(ctx context.Context, w io.Writer) {
	if c.Disabled {
		io.WriteString(w, ` disabled`)
	}
}

type ButtonGroup struct {
	ID       string
	Class    string
	Size     string
	Vertical bool
	Yield    func()
}

func (c *ButtonGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="group"`)
	writeAttr(w, "id", c.ID)
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
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type ButtonToolbar struct {
	ID       string
	Class    string
	Size     string
	Vertical bool
	Yield    func()
}

func (c *ButtonToolbar) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="toolbar"`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *ButtonToolbar) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="btn-toolbar`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CloseButton struct {
	ID      string
	Class   string
	Dismiss string
}

func (c *CloseButton) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<button type="button"`)
	writeAttr(w, "id", c.ID)
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
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
