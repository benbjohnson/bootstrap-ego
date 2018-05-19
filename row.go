package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Row struct {
	ID    string
	Class string
	Yield func()
}

func (r *Row) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *Row) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="row`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
