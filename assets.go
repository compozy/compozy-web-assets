// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5ee8511b3ad4076db5e8d61394d955561bf2a84ee4986499b19569b29918a5c6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ecdca847b3df9325b45751e8cc9a6182f8e0aea6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
