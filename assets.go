// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "13d7b2dde713bc150c8cafa269cf4559c160a197a56cb67ee139d6a9e4554d8c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2d471bc5cd31a8954e928599a57625767a82095f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
