package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Jumbotron struct {
	ID    string
	Class string
	Yield func()
}

func (r *Jumbotron) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}
	io.WriteString(w, `</div>`)
}

func (r *Jumbotron) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="jumbotron`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
