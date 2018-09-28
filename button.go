package bootstrap

import (
	"context"
	"fmt"
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
	Title    string
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

	fmt.Fprintf(w, `<%s`, nodeName)
	if nodeName == "a" {
		io.WriteString(w, ` role="button"`)
	} else {
		r.renderType(ctx, w)
	}
	writeAttr(w, "id", r.ID)
	writeAttr(w, "name", r.Name)
	r.renderClass(ctx, w)
	writeAttr(w, "href", r.Href)
	writeAttr(w, "form", r.Form)
	writeAttr(w, "target", r.Target)
	writeAttr(w, "value", r.Value)
	writeAttr(w, "title", r.Title)
	if r.Disabled {
		io.WriteString(w, ` disabled`)
	}
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

		fmt.Fprintf(w, `</%s>`, nodeName)
	}
}

func (r *Button) renderType(ctx context.Context, w io.Writer) {
	switch r.Type {
	case "submit", "reset", "button":
		fmt.Fprintf(w, ` type="%s"`, r.Type)
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
			fmt.Fprintf(w, ` btn-outline-%s`, r.Style)
		} else {
			fmt.Fprintf(w, ` btn-%s`, r.Style)
		}
	case "":
	default:
		Logger.Printf("bootstrap.Button: Invalid style: %q", r.Style)
	}

	// Set button size.
	switch r.Size {
	case "sm", "lg":
		fmt.Fprintf(w, ` btn-%s`, r.Size)
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
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
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
		fmt.Fprintf(w, ` btn-group-%s`, r.Size)
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
		fmt.Fprintf(w, " %s", html.EscapeString(s))
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
		fmt.Fprintf(w, " %s", html.EscapeString(s))
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
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}
