// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4d8307ca58c2b3b5e6b2cb2d00f753342212748a461ba42ccc9b4559bcc1c826"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4d32815602d5caca92c61e42ba77b0f77e97de3d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
