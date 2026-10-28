// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f6081e40e3291b3b906cdf6883fbcf9d3570d55c8d717726cc2059ab95f8d23b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "670ab2aa5f545cf9a81c751f532825747358b41c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
