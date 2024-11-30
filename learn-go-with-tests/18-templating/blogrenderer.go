package blogrenderer

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type PostRenderer struct {
	templ          *template.Template
	markdownParser *parser.Parser
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, fmt.Errorf("ParseFS error: %s", err.Error())
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, markdownParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", newPostViewModel(p, r)); err != nil {
		return fmt.Errorf("ExecuteTemplate error: %s", err.Error())
	}

	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	err := r.templ.ExecuteTemplate(w, "index.gohtml", posts)
	if err != nil {
		return fmt.Errorf("ExecuteTemplate error: %s", err.Error())
	}

	return nil
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

// we'll trust the security of template.Template:
//
//nolint:gosec
func newPostViewModel(p Post, r *PostRenderer) postViewModel {
	htmlBody := template.HTML(markdown.ToHTML([]byte(p.Body), r.markdownParser, nil))
	viewModel := postViewModel{Post: p, HTMLBody: htmlBody}

	return viewModel
}
