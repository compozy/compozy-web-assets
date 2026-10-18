// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ca60c57f66f72eab1bc0eba86c652e5139f6464d777ffaf52d4aea415ada981e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "727b9785d25c01a9283e86dba47c714fb2dc0a32"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
