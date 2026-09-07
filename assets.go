// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1065cbaa46202debfcf47aae8b1cf84d2cb852323424a17f03861e87a20f6bb3"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "52d2c4a63f8dbaa429ff61f31671f52bee8a518b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
