// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8a1831fd675fc907e8907e8b4c5a4ff70742cfd9acf2ae358150804ebbc3f8fe"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7e3753198b7349727f852ca53a38826c66d8467e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
