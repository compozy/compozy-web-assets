// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c4bfc3f70c35cd6140460a2d01749713b8df9993f02dc7b3ad4570a4342d6af4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4d92cd254b9f5d4ae211158edd140be68d2ab2b3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
