package bootstrap

import (
	"context"
	"html"
	"io"
)

type Card struct {
	ID     string
	Class  string
	Header func()
	Body   func()
	Yield  func()
	Footer func()
}

func (c *Card) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Header != nil {
		io.WriteString(w, `<div class="card-header">`)
		c.Header()
		io.WriteString(w, `</div>`)
	}

	if c.Body != nil {
		io.WriteString(w, `<div class="card-body">`)
		c.Body()
		io.WriteString(w, `</div>`)
	}

	if c.Yield != nil {
		c.Yield()
	}

	if c.Footer != nil {
		io.WriteString(w, `<div class="card-footer">`)
		c.Footer()
		io.WriteString(w, `</div>`)
	}

	io.WriteString(w, `</div>`)
}

func (c *Card) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CardGroup struct {
	ID    string
	Class string
	Yield func()
}

func (c *CardGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *CardGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card-group`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CardDeck struct {
	ID    string
	Class string
	Yield func()
}

func (c *CardDeck) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *CardDeck) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card-deck`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CardColumns struct {
	ID    string
	Class string
	Yield func()
}

func (c *CardColumns) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", c.ID)
	c.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if c.Yield != nil {
		c.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (c *CardColumns) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card-columns`)
	if s := c.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
