// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "52fd5849e6963a51a0f6839f6c5fc07c5744f1664582294fb62170f42d8247f4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "66b6a3741aec3456867e11ed08055a4e3ecb97ed"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
