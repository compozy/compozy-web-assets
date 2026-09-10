// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "130bb405101cc30911bb176e427187ea585d900fedcddb3a341c3bfac85c6e82"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0bbae40b4ebe6b90d76e174f97514d8a29eca68d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
