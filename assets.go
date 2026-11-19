// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "bea30c44f381cbdce1486bc5c7c80dfb4d0bb9d5053aa842737d38cf9ed33b50"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "61a7a8238fc3a30d6e14d97dd48e5f5b4a9cd84b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
