package handler

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// Parses the "startIndex" and "count" query parameters. This will not return
// an error if the values are absent, instead it will return defaults. An error
// will only be returned if one of the values are present but are not in the
// correct format. The error message is suitable for sending in a plaintext
// error response.
func parseStartCount(r *http.Request) (uint64, uint64, error) {
	var startIndex uint64 = 0
	var count uint64 = 10
	var err error

	if r.URL.Query().Has("startIndex") {
		startIndex, err = strconv.ParseUint(
			r.URL.Query().Get("startIndex"),
			10,
			64,
		)
		if err != nil {
			return 0, 0, errors.New("startIndex must be a positive integer")
		}
	}

	if r.URL.Query().Has("count") {
		count, err = strconv.ParseUint(
			r.URL.Query().Get("count"),
			10,
			64,
		)
		if err != nil {
			return 0, 0, errors.New("count must be a positive integer")
		}
	}

	return startIndex, count, nil
}

// If the request header indicates that it can handle compressed data in one
// of the formats we recognize, then we will send the data as JSON in that 
// compressed format. As a fallback, the uncompressed JSON will be sent.
func sendCompressedJson(w http.ResponseWriter, r *http.Request, data any) {
	w.Header().Add("Content-Type", "application/json")

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Add("Content-Encoding", "gzip")

		compressor := gzip.NewWriter(w)
		defer compressor.Close()

		encoder := json.NewEncoder(compressor)
		encoder.Encode(data)

		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

// Sends uncompressed JSON in the response.
func sendJson(w http.ResponseWriter, r *http.Request, data any) {
	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
