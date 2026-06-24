package riot

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
)

// Decodes a response body as JSON. If the "Content-Encoding" header indicates
// that it's compressed in a format we recognize, then the body will be
// decompressed accordingly. At the time of writing, the recognized encodings
// are: gzip, x-gzip.
//
// Note: The response body will not be closed by this function, even though it
// will be fully read.
func decodeBody(dst any, resp *http.Response) error {

	var body io.Reader

	switch resp.Header.Get("Content-Encoding") {
	case "gzip", "x-gzip":
		r, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		defer r.Close()
		body = r
	default:
		body = resp.Body
	}

	decoder := json.NewDecoder(body)
	return decoder.Decode(dst)
}
