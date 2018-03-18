package bootstrap

import (
	"context"
	"io"
)

type Card struct {
	Class  string
	Header func()
	Body   func()
	Yield  func()
	Footer func()
}

func (c *Card) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
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
	if c.Class != "" {
		io.WriteString(w, ` `)
		io.WriteString(w, c.Class)
	}
	io.WriteString(w, `"`)
}
