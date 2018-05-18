package bootstrap

import (
	"context"
	"html"
	"io"
)

type FormGroup struct {
	ID    string
	Class string
	Yield func()
}

func (c *FormGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *FormGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-group`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type FormRow struct {
	ID    string
	Class string
	Yield func()
}

func (c *FormRow) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *FormRow) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-row`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type FormText struct {
	ID    string
	Class string
	Yield func()
}

func (c *FormText) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *FormText) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-text`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type FormCheck struct {
	ID     string
	Class  string
	Inline bool
	Yield  func()
}

func (c *FormCheck) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *FormCheck) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-check`)
	if c.Inline {
		io.WriteString(w, ` form-check-inline`)
	}
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

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
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Prepend != nil {
		io.WriteString(w, `<div class="input-group-prepend">`)
		io.WriteString(w, `<span class="input-group-text">`)
		c.Prepend()
		io.WriteString(w, `</span>`)
		io.WriteString(w, `</div>`)
	}

	if c.Yield != nil {
		c.Yield()
	}

	if c.Append != nil {
		io.WriteString(w, `<div class="input-group-append">`)
		io.WriteString(w, `<span class="input-group-text">`)
		c.Append()
		io.WriteString(w, `</span>`)
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

	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type InputGroupText struct {
	ID    string
	Class string
	Yield func()
}

func (c *InputGroupText) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}
	io.WriteString(w, `</div>`)
}

func (c *InputGroupText) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="input-group-text`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type InvalidFeedback struct {
	ID    string
	Class string
	Yield func()
}

func (c *InvalidFeedback) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *InvalidFeedback) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="invalid-feedback`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type ValidFeedback struct {
	ID    string
	Class string
	Yield func()
}

func (c *ValidFeedback) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *ValidFeedback) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="valid-feedback`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
