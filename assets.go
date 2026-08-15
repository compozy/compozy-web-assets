// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7f71ded7fbddd85a5ea92e1969d15cbd3463a395d55fcb85959b299dcf453945"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "21b86011ee564da09c00ceb0001b5db3ce70b541"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
