package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Container struct {
	ID    string
	Class string
	Fluid bool
	Yield func()
}

func (r *Container) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *Container) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="`)
	if r.Fluid {
		io.WriteString(w, "container-fluid")
	} else {
		io.WriteString(w, "container")
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
