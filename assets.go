// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "748de31e6ab104c1d4c3a5e2060d4ae2e9e45a86082fd1de54868a0798e817dc"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8fff21d200831b4c10ea1ce08e1bc47f5f41689d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
