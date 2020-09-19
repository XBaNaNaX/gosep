package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

type rrWrit struct {
	io.Writer
	err error
}

func (w *rrWrit) Write(p []byte) (n int, rr error) {
	if w.err != nil {
		return 0, w.err
	}

	n, w.err = w.Writer.Write(p)

	return n, nil
}

func New(w io.Writer) *rrWrit {
	return &rrWrit{Writer: w}
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ww := New(w)
	fmt.Fprintf(ww, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ww, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprint(ww, "\r\n")

	io.Copy(w, body)
	return ww.err
}

func main() {
	var buf bytes.Buffer
	st := Status{Code: 555, Reason: "account not found."}
	headers := []Header{
		{"Content-Type", "application/json"},
	}
	body := strings.NewReader("this is a body")

	err := WriteResponse(&buf, st, headers, body)

	fmt.Printf("buf: \n%s\n", buf.String())
	fmt.Println("error:", err)
}
