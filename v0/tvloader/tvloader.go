// Package tvloader is used to load and parse SPDX tag-value documents
// into tools-golang data structures.
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package tvloader

import (
	"io"

	"github.com/spdx/tools-golang/v0/spdx"
	"github.com/spdx/tools-golang/v0/tvloader/parser2v1"
	"github.com/spdx/tools-golang/v0/tvloader/reader"
)

// Load2_1 takes an io.Reader and returns a fully-parsed SPDX Document
// (version 2.1) if parseable, or error if any error is encountered.
func Load2_1(content io.Reader) (*spdx.Document2_1, error) {
	tvPairs, err := reader.ReadTagValues(content)
	if err != nil {
		return nil, err
	}

	doc, err := parser2v1.ParseTagValues(tvPairs)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
