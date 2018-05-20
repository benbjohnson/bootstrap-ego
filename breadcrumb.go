package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Breadcrumb struct {
	ID    string
	Class string
	Yield func()
}

func (r *Breadcrumb) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<ol`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</ol>`)
}

func (r *Breadcrumb) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="breadcrumb`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type BreadcrumbItem struct {
	ID     string
	Class  string
	Active bool
	Yield  func()
}

func (r *BreadcrumbItem) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<li`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</li>`)
}

func (r *BreadcrumbItem) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="breadcrumb-item`)
	if r.Active {
		io.WriteString(w, ` active`)
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
