// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e1ff907890f6497ee7e1dd160acab161331603acdcfe3f89e5485d09deb57d36"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "33a4fab4b2db4b2387d19d7418e640b55d4b1ddf"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
