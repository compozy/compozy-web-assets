// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3dce0ec22f87b012810994c78652a8f33afc8b75e2b46733a9c86328b2d4cd02"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6358a84084adcf16b663203a0fb1027bec8050d5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
