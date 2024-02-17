package templating

import (
	"bytes"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title: "hello world",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)

	// t.Run("it converts a single post into HTML", func (t *testing.T){
	// 	buf := bytes.Buffer{}
	// 	err := Render(&buf, aPost)

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	got := buf.String()
	// 	want := "<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n\t<title>My amazing blog!</title>\r\n\t<meta charset=\"UTF-8\"/>\r\n\t<meta name=\"description\" content=\"Wow, like and subscribe, it really helps the channel guys\" lang=\"en\"/>\r\n</head>\r\n<body>\r\n<nav role=\"navigation\">\r\n\t<div>\r\n\t\t<h1>Budding Gopher's blog</h1>\r\n\t\t<ul>\r\n\t\t\t<li><a href=\"/\">home</a></li>\r\n\t\t\t<li><a href=\"about\">about</a></li>\r\n\t\t\t<li><a href=\"archive\">archive</a></li>\r\n\t\t</ul>\r\n\t</div>\r\n</nav>\r\n<main>\r\n\r\n\r\n<h1>hello world</h1>\r\n\r\n<p>This is a description</p>\r\n\r\nTags: <ul><li>go</li><li>tdd</li></ul>\r\n\r\n</main>\r\n<footer>\r\n\t<ul>\r\n\t\t<li> <a href=\"https://twitter.com/quii\">Twitter</a></li>\r\n\t\t<li> <a href=\"https://github.com/quii\">Github</a></li>\r\n\t</ul>\r\n</footer>\r\n</body>\r\n</html>\r\n"
	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// })

	postRenderer, err := NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func (t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		// install approval for the comparison
	})

	t.Run("it renders an index of posts", func (t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "hello-world"}, {Title: "hello-gophers"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/post/hello-world">hello-world</a></li><li><a href="/post/hello-gophers">hello-gophers</a></li></ol>`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = Post{
			Title: "hello world",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)

	postRenderer, err := NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
