// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0b55efd80772014885525a9d679c86c2733d16c18869cdbf827d11f40dce7a3d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b153f4a736e73e6f8c6f68a3db0088d3bdf0a703"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
