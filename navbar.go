package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Navbar struct {
	ID     string
	Class  string
	Style  string
	Expand string
	Yield  func()
}

func (r *Navbar) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *Navbar) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="navbar`)

	switch r.Style {
	case "light", "dark":
		fmt.Fprintf(w, ` navbar-%s`, r.Style)
	case "":
	default:
		Logger.Printf("bootstrap.Navbar: Invalid style: %q", r.Style)
	}

	switch r.Expand {
	case "sm", "md", "lg", "xl":
		fmt.Fprintf(w, ` navbar-expand-%s`, r.Expand)
	case "":
	default:
		Logger.Printf("bootstrap.Navbar: Invalid expand: %q", r.Expand)
	}

	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type NavbarBrand struct {
	ID       string
	NodeName string
	Class    string
	Href     string
	Target   string
	Yield    func()
}

func (r *NavbarBrand) Render(ctx context.Context, w io.Writer) {
	nodeName := r.NodeName
	if nodeName == "" {
		nodeName = "a"
	}
	fmt.Fprintf(w, `<%s`, nodeName)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	writeAttr(w, "href", r.Href)
	writeAttr(w, "target", r.Target)
	fmt.Fprintf(w, `</%s>`, nodeName)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</a>`)
}

func (r *NavbarBrand) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="navbar-brand`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type NavbarToggler struct {
	ID     string
	Class  string
	Target string
}

func (r *NavbarToggler) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<button`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	writeAttr(w, "data-target", r.Target)
	io.WriteString(w, ` type="button"`)
	io.WriteString(w, ` data-toggle="collapse"`)
	io.WriteString(w, ` aria-expanded="false"`)
	io.WriteString(w, `>`)
	io.WriteString(w, `<span class="navbar-toggler-icon"></span>`)
	io.WriteString(w, `</button>`)
}

func (r *NavbarToggler) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="navbar-toggler`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type NavbarCollapse struct {
	ID    string
	Class string
	Yield func()
}

func (r *NavbarCollapse) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *NavbarCollapse) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="collapse navbar-collapse`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type NavbarNav struct {
	ID       string
	NodeName string
	Class    string
	Yield    func()
}

func (r *NavbarNav) Render(ctx context.Context, w io.Writer) {
	nodeName := r.NodeName
	if nodeName == "" {
		nodeName = "ul"
	}
	fmt.Fprintf(w, `<%s`, nodeName)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	fmt.Fprintf(w, `</%s>`, nodeName)
}

func (r *NavbarNav) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="navbar-nav`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type NavbarText struct {
	ID    string
	Class string
	Yield func()
}

func (r *NavbarText) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<span`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</span>`)
}

func (r *NavbarText) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="navbar-text`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
