package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/shogo82148/base45"
	"github.com/shogo82148/ridgenative"
)

//go:embed show-link.html
var showLinkHTML string

var showLinkTemplate = template.Must(template.New("show-link").Parse(showLinkHTML))

//go:embed playground.html playground.js url-compressor.js huffman.js
var playgroundFiles embed.FS

var playgroundHandler = http.FileServer(http.FS(playgroundFiles))

func main() {
	http.HandleFunc("/", serveRoot)
	ridgenative.ListenAndServe(":8080", nil)
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/playground.html", "/playground.js", "/url-compressor.js", "/huffman.js":
		serveIndex(w, r)
		return
	}

	url, err := decode(strings.TrimPrefix(r.URL.Path, "/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = showLinkTemplate.Execute(w, struct{ Link string }{Link: "https://" + string(url)})
	if err != nil {
		panic(err)
	}
}

// serveIndex only serves the explicitly embedded playground assets.
func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path == "/" {
		r = r.Clone(r.Context())
		r.URL.Path = "/playground.html"
	}
	playgroundHandler.ServeHTTP(w, r)
}

var escaper = strings.NewReplacer(
	" ", ".A",
	// "$", ".B",
	"%", ".C",
	"*", ".D",
	"+", ".E",
	// "-", ".F",
	".", ".G",
	"/", ".H",
	// ":", ".I",
)

func encode(data []byte) string {
	buf := NewBuffer([]byte{})

	for _, c := range data {
		e := encodeTable[c]
		buf.WriteBitsLSB(e.code, e.length)
	}

	// Encode the terminator.
	for buf.Len()%8 != 0 {
		buf.WriteBit(1)
	}

	return escaper.Replace("0" + base45.EncodeToString(buf.Bytes()))
}

var decoder = strings.NewReplacer(
	".A", " ",
	".B", "$",
	".C", "%",
	".D", "*",
	".E", "+",
	".F", "-",
	".G", ".",
	".H", "/",
	".I", ":",
)

func decode(data string) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data")
	}
	version := data[0]
	if version == '0' {
		return decode0(data[1:])
	}
	return nil, fmt.Errorf("unknown version: %c", version)
}

func decode0(data string) ([]byte, error) {
	tmp, err := base45.DecodeString(decoder.Replace(data))
	if err != nil {
		return nil, err
	}
	buf := NewBuffer(tmp)

	ret := []byte{}
	n := decodeTree
	for {
		bit, err := buf.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		n = n.children[bit]
		if n.children[0] == nil && n.children[1] == nil {
			ret = append(ret, byte(n.value))
			n = decodeTree
		}
	}
	return ret, nil
}
