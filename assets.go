// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dd139525e6dd9108486375b3cf777490e835abd16101ba000130a8dcc55548c4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a35eda6d3a2ec47995c19a14a5a01d4f9452cf1c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
