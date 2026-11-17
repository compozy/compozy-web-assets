// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f49b5121f3a55258dd14271f14c356f02e68c702749a05ff509172599ac1a6dd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7dc53a2092d148661e2ea156e0d77d86c34dd03d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
