// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "962ffd70be467acf714312088059f8a53c3ad256dfe6dc7263ad251a21bfa703"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0f326683c9df954971ce8e043ed08e9f57482c1b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
