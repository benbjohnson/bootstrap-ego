package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type FormGroup struct {
	ID    string
	Class string
	Yield func()
}

func (r *FormGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *FormGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-group`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type FormRow struct {
	ID    string
	Class string
	Yield func()
}

func (r *FormRow) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *FormRow) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-row`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type FormText struct {
	ID    string
	Class string
	Yield func()
}

func (r *FormText) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *FormText) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-text`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type FormCheck struct {
	ID     string
	Class  string
	Inline bool
	Yield  func()
}

func (r *FormCheck) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *FormCheck) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="form-check`)
	if r.Inline {
		io.WriteString(w, ` form-check-inline`)
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type InputGroup struct {
	ID      string
	Class   string
	Size    string
	Prepend func()
	Append  func()
	Yield   func()
}

func (r *InputGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Prepend != nil {
		io.WriteString(w, `<div class="input-group-prepend">`)
		io.WriteString(w, `<span class="input-group-text">`)
		r.Prepend()
		io.WriteString(w, `</span>`)
		io.WriteString(w, `</div>`)
	}

	if r.Yield != nil {
		r.Yield()
	}

	if r.Append != nil {
		io.WriteString(w, `<div class="input-group-append">`)
		io.WriteString(w, `<span class="input-group-text">`)
		r.Append()
		io.WriteString(w, `</span>`)
		io.WriteString(w, `</div>`)
	}

	io.WriteString(w, `</div>`)
}

func (r *InputGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="input-group`)

	switch r.Size {
	case "sm", "lg":
		fmt.Fprintf(w, ` input-group-%s`, r.Size)
	case "":
	default:
		Logger.Printf("bootstrap.InputGroup: Invalid size: %q", r.Size)
	}

	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type InputGroupText struct {
	ID    string
	Class string
	Yield func()
}

func (r *InputGroupText) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}
	io.WriteString(w, `</div>`)
}

func (r *InputGroupText) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="input-group-text`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type InvalidFeedback struct {
	ID    string
	Class string
	Yield func()
}

func (r *InvalidFeedback) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *InvalidFeedback) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="invalid-feedback`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type ValidFeedback struct {
	ID    string
	Class string
	Yield func()
}

func (r *ValidFeedback) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ValidFeedback) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="valid-feedback`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
