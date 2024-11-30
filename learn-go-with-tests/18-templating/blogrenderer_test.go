package blogrenderer_test

import (
	"bytes"
	"go_academy/learn-go-with-tests/18-templating"
	"testing"

	"github.com/approvals/go-approval-tests" // added via `$ go get github.com/approvals/go-approval-tests`
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, newRendererErr := blogrenderer.NewPostRenderer()

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		if newRendererErr != nil {
			t.Fatal(newRendererErr)
		}

		buf := bytes.Buffer{}
		if renderErr := postRenderer.Render(&buf, aPost); renderErr != nil {
			t.Fatal(renderErr)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{
			Title:       "Hello World",
			Description: "",
			Body:        "",
			Tags:        []string{},
		}, {
			Title:       "Hello World 2",
			Description: "",
			Body:        "",
			Tags:        []string{},
		}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

// issue intro'd with markdown.ToHTML: https://github.com/quii/learn-go-with-tests/issues/638
// func BenchmarkRender(b *testing.B) {
// 	var (
// 		aPost = blogrenderer.Post{
// 			Title:       "hello world",
// 			Body:        "This is a post",
// 			Description: "This is a description",
// 			Tags:        []string{"go", "tdd"},
// 		}
// 	)
// 	postRenderer, newRendererErr := blogrenderer.NewPostRenderer()
// 	if newRendererErr != nil {
// 		b.Fatal(newRendererErr)
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		postRenderer.Render(io.Discard, aPost)
// 	}
// }
