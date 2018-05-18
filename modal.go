package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Modal struct {
	ID    string
	Class string
	Yield func()
}

func (r *Modal) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, ` tabindex="-1"`)
	io.WriteString(w, ` role="dialog"`)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *Modal) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="modal`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type ModalDialog struct {
	ID       string
	Class    string
	Size     string
	Centered bool
	Yield    func()
}

func (r *ModalDialog) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, ` role="document"`)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ModalDialog) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="modal-dialog`)

	if r.Centered {
		io.WriteString(w, ` modal-dialog-centered`)
	}

	switch r.Size {
	case "sm", "lg":
		fmt.Fprintf(w, ` modal-%s`, r.Size)
	case "":
	default:
		Logger.Printf("bootstrap.ModalDialog: Invalid size: %q", r.Size)
	}

	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type ModalContent struct {
	ID    string
	Class string
	Yield func()
}

func (r *ModalContent) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ModalContent) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="modal-content`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type ModalHeader struct {
	ID    string
	Class string
	Yield func()
}

func (r *ModalHeader) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ModalHeader) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="modal-header`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type ModalBody struct {
	ID    string
	Class string
	Yield func()
}

func (r *ModalBody) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ModalBody) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="modal-body`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type ModalFooter struct {
	ID    string
	Class string
	Yield func()
}

func (r *ModalFooter) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ModalFooter) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="modal-footer`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
