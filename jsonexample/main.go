package main

import (
    "bytes"
    "fmt"
    "github.com/tdewolff/minify"
)

func HtmlMinify(html string) string {
    m := minify.NewMinifierDefault()
    b := &bytes.Buffer{}
    if err := m.HTML(b, bytes.NewBufferString(html)); err != nil {
        panic(err)
    }
    return b.String()
}

func main() {

    htmlExample := `<li>
                        <a>Hello</a>
                    </li>`
    minifiedHtml := HtmlMinify(htmlExample)
    fmt.Println(minifiedHtml) //  <li><a>Hello</a>
}