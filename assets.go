// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "078ed4bb2336ffcdd3d8c701a0c50f5c5d6781d6e1b8bbda44aa3ed32f49a85d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "16cf61ea3bd59911ba4bd0647ed479b4611f7510"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
