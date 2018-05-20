package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Card struct {
	ID     string
	Class  string
	Header func()
	Body   func()
	Yield  func()
	Footer func()
}

func (r *Card) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Header != nil {
		io.WriteString(w, `<div class="card-header">`)
		r.Header()
		io.WriteString(w, `</div>`)
	}

	if r.Body != nil {
		io.WriteString(w, `<div class="card-body">`)
		r.Body()
		io.WriteString(w, `</div>`)
	}

	if r.Yield != nil {
		r.Yield()
	}

	if r.Footer != nil {
		io.WriteString(w, `<div class="card-footer">`)
		r.Footer()
		io.WriteString(w, `</div>`)
	}

	io.WriteString(w, `</div>`)
}

func (r *Card) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CardGroup struct {
	ID    string
	Class string
	Yield func()
}

func (r *CardGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *CardGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card-group`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CardDeck struct {
	ID    string
	Class string
	Yield func()
}

func (r *CardDeck) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *CardDeck) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card-deck`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CardColumns struct {
	ID    string
	Class string
	Yield func()
}

func (r *CardColumns) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *CardColumns) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="card-columns`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
