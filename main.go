package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/shogo82148/base45"
	"github.com/shogo82148/ridgenative"
)

//go:embed show-link.html
var showLinkHTML string

var showLinkTemplate = template.Must(template.New("show-link").Parse(showLinkHTML))

func main() {
	http.HandleFunc("/", serveRoot)
	ridgenative.ListenAndServe(":8080", nil)
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
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

//go:embed show-preview.html
var showPreviewHTML string

var showPreviewTemplate = template.Must(template.New("show-preview").Parse(showPreviewHTML))

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Query().Get("url")

	// remove the scheme
	u := url
	u = strings.TrimPrefix(u, "http://")
	u = strings.TrimPrefix(u, "https://")

	host := os.Getenv("COMPRESSOR_HOSTNAME")
	if host == "" {
		host = r.Host
	}

	encoded := "HTTPS://" + strings.ToUpper(host) + "/" + encode([]byte(u))
	err := showPreviewTemplate.Execute(w, struct{ Encoded, Link string }{Encoded: encoded, Link: url})
	if err != nil {
		panic(err)
	}
}

var escaper = strings.NewReplacer(
	" ", "%20",
	"%", "%25",
	"+", "%2B",
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
	tmp, err := base45.DecodeString(data)
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
