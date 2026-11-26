// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "909dcd8712dbbd043f0406d7c4e99934a41f3d0265e1fc3195ba4e922d3e2188"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "820fc4d8674a911e9236cb3b363991c0f3df5039"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
