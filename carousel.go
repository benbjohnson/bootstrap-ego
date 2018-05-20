package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Carousel struct {
	ID         string
	Class      string
	Indicators func()
	Inner      func()
	Controls   func()
}

func (r *Carousel) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, ` data-ride="carousel"`)
	io.WriteString(w, `>`)

	if r.Indicators != nil {
		r.Indicators()
	}

	if r.Inner != nil {
		io.WriteString(w, `<div class="carousel-inner">`)
		r.Inner()
		io.WriteString(w, `</div>`)
	}

	if r.Controls != nil {
		r.Controls()
	}

	io.WriteString(w, `</div>`)
}

func (r *Carousel) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel slide`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CarouselItem struct {
	ID     string
	Class  string
	Active bool
	Yield  func()
}

func (r *CarouselItem) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<div`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</div>`)
}

func (r *CarouselItem) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-item`)
	if r.Active {
		io.WriteString(w, ` active`)
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CarouselIndicators struct {
	ID    string
	Class string
	Yield func()
}

func (r *CarouselIndicators) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<ol`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</ol>`)
}

func (r *CarouselIndicators) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-indicators`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

type CarouselIndicator struct {
	ID      string
	Class   string
	Target  string
	SlideTo int
	Active  bool
}

func (r *CarouselIndicator) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<li`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	r.renderDataTarget(ctx, w)
	r.renderDataSlideTo(ctx, w)
	io.WriteString(w, `>`)
	io.WriteString(w, `</li>`)
}

func (r *CarouselIndicator) renderClass(ctx context.Context, w io.Writer) {
	if !r.Active && r.Class == "" {
		return
	}
	io.WriteString(w, ` class="`)
	if r.Active {
		io.WriteString(w, `active`)
	}
	if r.Active && r.Class != "" {
		io.WriteString(w, ` `)
	}
	if r.Class != "" {
		io.WriteString(w, html.EscapeString(r.Class))
	}
	io.WriteString(w, `"`)
}

func (r *CarouselIndicator) renderDataTarget(ctx context.Context, w io.Writer) {
	if r.Target == "" {
		return
	}
	fmt.Fprintf(w, ` data-target="#%s"`, html.EscapeString(r.Target))
}

func (r *CarouselIndicator) renderDataSlideTo(ctx context.Context, w io.Writer) {
	fmt.Fprintf(w, ` data-slide-to="%d"`, r.SlideTo)
}

type CarouselControlPrev struct {
	ID     string
	Class  string
	Target string
}

func (r *CarouselControlPrev) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<a`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	r.renderHref(ctx, w)
	io.WriteString(w, ` role="button" data-slide="prev"`)
	io.WriteString(w, `>`)
	io.WriteString(w, `<span class="carousel-control-prev-icon" aria-hidden="true"></span>`)
	io.WriteString(w, `<span class="sr-only">Previous</span>`)
	io.WriteString(w, `</a>`)
}

func (r *CarouselControlPrev) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-control-prev`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

func (r *CarouselControlPrev) renderHref(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` href="#`)
	if s := r.Target; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type CarouselControlNext struct {
	ID     string
	Class  string
	Target string
}

func (r *CarouselControlNext) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<a`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	r.renderHref(ctx, w)
	io.WriteString(w, ` role="button" data-slide="next"`)
	io.WriteString(w, `>`)
	io.WriteString(w, `<span class="carousel-control-next-icon" aria-hidden="true"></span>`)
	io.WriteString(w, `<span class="sr-only">Nextious</span>`)
	io.WriteString(w, `</a>`)
}

func (r *CarouselControlNext) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="carousel-control-next`)
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

func (r *CarouselControlNext) renderHref(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` href="#`)
	if s := r.Target; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
