// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "09d223e296f1869f4b2002d4fe5ca43943458883c20919c78931df16199c7a0d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6635200183a928bd7e1d53cae9f542dcf0941e54"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
