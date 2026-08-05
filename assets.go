// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3d581902ac048b6fcb25f95be0756d7086cdb3acd3d4f94ff50a9281de15734e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3a95b5bb972cf35a37cd2765fa23a12f0f88d1b5"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
