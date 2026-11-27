// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0ea6104e3bf56d11168360587f7d2e45f9c4068ac88c65977beee5d29939c84d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "27b8a232a04eac2c3da117438c7558202fa9b667"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
