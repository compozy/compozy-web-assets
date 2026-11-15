// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "72ae402887ce74d87d60ab426554278365e9a6da478591ab0dc519fb9b64db6e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b494e65f74e45277f7f0ae583cda2a487d7812e0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
