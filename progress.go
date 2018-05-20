package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Progress struct {
	ID    string
	Class string
	Yield func()
}

func (r *Progress) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *Progress) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="progress`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type ProgressBar struct {
	ID       string
	Class    string
	Striped  bool
	Animated bool
	Value    float64
	Yield    func()
}

func (r *ProgressBar) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	fmt.Fprintf(w, ` style="width:%f%%"`, r.Value)
	fmt.Fprintf(w, ` aria-valuenow="%f%%" `, r.Value)
	io.WriteString(w, ` aria-valuemin="0"`)
	io.WriteString(w, ` aria-valuemax="100"`)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ProgressBar) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="progress-bar`)
	if r.Striped {
		io.WriteString(w, ` progress-bar-striped`)
	}
	if r.Animated {
		io.WriteString(w, ` progress-bar-animated`)
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
