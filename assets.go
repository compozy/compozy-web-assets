// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b684df12a45c127f76b50b2792acb829293298e3fc94a9c22a0fe03dc4e1ede3"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "082b33a92dc82ecfb2f6c0c7ecd3459cf3b1e9fa"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
