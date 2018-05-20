package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Pagination struct {
	ID    string
	Class string
	Size  string
	Yield func()
}

func (r *Pagination) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<ul`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</ul>`)
}

func (r *Pagination) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="pagination`)

	switch r.Size {
	case "sm", "lg":
		fmt.Fprintf(w, ` pagination-%s`, r.Size)
	case "":
	default:
		Logger.Printf("bootstrap.Pagination: Invalid size: %q", r.Size)
	}

	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type PageItem struct {
	ID       string
	Class    string
	Active   bool
	Disabled bool
	Yield    func()
}

func (r *PageItem) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<li`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</li>`)
}

func (r *PageItem) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="page-item`)
	if r.Active {
		io.WriteString(w, ` active`)
	}
	if r.Disabled {
		io.WriteString(w, ` disabled`)
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type PageLink struct {
	ID       string
	NodeName string
	Class    string
	Href     string
	Target   string
	Disabled bool
	Yield    func()
}

func (r *PageLink) Render(ctx context.Context, w io.Writer) {
	nodeName := r.NodeName
	if nodeName == "" {
		nodeName = "a"
	}
	fmt.Fprintf(w, `<%s`, nodeName)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	writeAttr(w, "href", r.Href)
	writeAttr(w, "target", r.Target)
	if r.Disabled {
		io.WriteString(w, ` tabindex="-1"`)
	}
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	fmt.Fprintf(w, `</%s>`, nodeName)
}

func (r *PageLink) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="page-link`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
