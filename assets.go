// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e97434f5cf03d4c955d44d8bb411565f5fc66b0dc016b8e1df0445c4f3e9a6d0"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b616c61c6fa373ba55143bae242c2d54092b966b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
