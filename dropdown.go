package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Dropdown struct {
	ID    string
	Class string
	Yield func()
}

func (r *Dropdown) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *Dropdown) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="dropdown`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type DropdownMenu struct {
	ID    string
	Class string
	Yield func()
}

func (r *DropdownMenu) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *DropdownMenu) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="dropdown-menu`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type DropdownDivider struct {
	ID    string
	Class string
}

func (r *DropdownDivider) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)
	io.WriteString(w, `</div>`)
}

func (r *DropdownDivider) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="dropdown-divider`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
