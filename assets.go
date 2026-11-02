// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e1ff907890f6497ee7e1dd160acab161331603acdcfe3f89e5485d09deb57d36"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "49cb8eabbdeca0842d0da2ffa404f5f2d7e311e7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
