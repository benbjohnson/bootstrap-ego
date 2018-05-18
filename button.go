package bootstrap

import (
	"context"
	"html"
	"io"
)

type Button struct {
	ID       string
	Class    string
	NodeName string
	Name     string
	Type     string
	Style    string
	Href     string
	Form     string
	Value    string
	Target   string
	Size     string // "sm","lg"
	Outline  bool
	Block    bool
	Active   bool
	Disabled bool
	Dropdown bool
	Yield    func()
}

func (r *Button) Render(ctx context.Context, w io.Writer) {
	nodeName := r.NodeName
	if r.Href != "" {
		nodeName = "a"
	} else if nodeName == "" {
		nodeName = "button"
	}

	io.WriteString(w, `<`)
	io.WriteString(w, nodeName)
	if nodeName == "a" {
		io.WriteString(w, ` role="button"`)
	} else {
		r.renderType(ctx, w)
	}
	writeAttr(w, "id", r.ID)
	writeAttr(w, "name", r.Name)
	r.renderClass(ctx, w)
	r.renderHref(ctx, w)
	r.renderForm(ctx, w)
	r.renderTarget(ctx, w)
	r.renderValue(ctx, w)
	r.renderDisabled(ctx, w)

	if r.Dropdown {
		io.WriteString(w, ` data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"`)
	}

	// Only allow yield for a/button.
	if nodeName == "input" {
		io.WriteString(w, `/>`)
	} else {
		io.WriteString(w, `>`)

		if r.Yield != nil {
			r.Yield()
		}

		io.WriteString(w, `</`)
		io.WriteString(w, nodeName)
		io.WriteString(w, `>`)
	}
}

func (r *Button) renderType(ctx context.Context, w io.Writer) {
	switch r.Type {
	case "submit", "reset", "button":
		io.WriteString(w, ` type="`)
		io.WriteString(w, r.Type)
		io.WriteString(w, `"`)
	case "":
		io.WriteString(w, ` type="button"`)
	default:
		Logger.Printf("bootstrap.Button: Invalid type: %q", r.Type)
	}
}

func (r *Button) renderClass(ctx context.Context, w io.Writer) {
	// Write style class.
	io.WriteString(w, ` class="btn`)
	switch r.Style {
	case "primary", "secondary", "success", "danger", "warning", "info", "light", "dark", "link":
		if r.Outline {
			io.WriteString(w, ` btn-outline-`)
		} else {
			io.WriteString(w, ` btn-`)
		}
		io.WriteString(w, r.Style)
	case "":
	default:
		Logger.Printf("bootstrap.Button: Invalid style: %q", r.Style)
	}

	// Set button size.
	switch r.Size {
	case "sm", "lg":
		io.WriteString(w, ` btn-`)
		io.WriteString(w, r.Size)
	case "":
	default:
		Logger.Printf("bootstrap.Button: Invalid size: %q", r.Size)
	}

	// Enable block-level button.
	if r.Block {
		io.WriteString(w, ` btn-block`)
	}

	// Enable active look.
	if r.Active {
		io.WriteString(w, ` active`)
	}

	// Enable dropdown.
	if r.Dropdown {
		io.WriteString(w, ` dropdown-toggle`)
	}

	// Write user-defined class.
	if s := r.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

func (r *Button) renderHref(ctx context.Context, w io.Writer) {
	if r.Href != "" {
		io.WriteString(w, ` href="`)
		io.WriteString(w, html.EscapeString(r.Href))
		io.WriteString(w, `"`)
	}
}

func (r *Button) renderForm(ctx context.Context, w io.Writer) {
	if r.Form != "" {
		io.WriteString(w, ` form="`)
		io.WriteString(w, html.EscapeString(r.Form))
		io.WriteString(w, `"`)
	}
}

func (r *Button) renderTarget(ctx context.Context, w io.Writer) {
	if r.Target != "" {
		io.WriteString(w, ` target="`)
		io.WriteString(w, html.EscapeString(r.Target))
		io.WriteString(w, `"`)
	}
}

func (r *Button) renderValue(ctx context.Context, w io.Writer) {
	if r.Value != "" {
		io.WriteString(w, ` value="`)
		io.WriteString(w, html.EscapeString(r.Value))
		io.WriteString(w, `"`)
	}
}

func (r *Button) renderDisabled(ctx context.Context, w io.Writer) {
	if r.Disabled {
		io.WriteString(w, ` disabled`)
	}
}

type ButtonGroup struct {
	ID       string
	Class    string
	Size     string
	Vertical bool
	Yield    func()
}

func (r *ButtonGroup) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="group"`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ButtonGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="btn-group`)

	// Write size class.
	switch r.Size {
	case "sm", "lg":
		io.WriteString(w, ` btn-group-`)
		io.WriteString(w, r.Size)
	case "":
	default:
		Logger.Printf("bootstrap.ButtonGroup: Invalid size: %q", r.Size)
	}

	// Enable vertical group.
	if r.Vertical {
		io.WriteString(w, ` btn-group-vertical`)
	}

	// Write user-defined class.
	if s := r.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type ButtonToolbar struct {
	ID       string
	Class    string
	Size     string
	Vertical bool
	Yield    func()
}

func (r *ButtonToolbar) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div role="toolbar"`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *ButtonToolbar) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="btn-toolbar`)
	if s := r.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CloseButton struct {
	ID      string
	Class   string
	Dismiss string
}

func (r *CloseButton) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<button type="button"`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)

	// Render 'dismiss' data.
	io.WriteString(w, ` data-dismiss="`)
	if r.Dismiss != "" {
		io.WriteString(w, html.EscapeString(r.Dismiss))
	} else {
		io.WriteString(w, "alert")
	}
	io.WriteString(w, `"`)

	// Write accessible label.
	io.WriteString(w, ` aria-label="Close"`)

	// Close tag.
	io.WriteString(w, `><span aria-hidden="true">&times;</span></button>`)
}

func (r *CloseButton) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="close`)
	if s := r.Class; s != "" {
		io.WriteString(w, " ")
		io.WriteString(w, html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
