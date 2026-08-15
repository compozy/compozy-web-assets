// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7f71ded7fbddd85a5ea92e1969d15cbd3463a395d55fcb85959b299dcf453945"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1c1564b85cb5203077356bb6eed02ea9daac3160"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
