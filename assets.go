// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "af9c3f5b415d8332d7d18c5255521cb57f3821202a57f1193cc11a788574cccf"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0a5754ad2603074342b1cfaa1c958ffecd8ea8da"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
