// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c947fd1f1fa95eef7be7862fd573b76eafdf91f0150b11af1eccfc5b6b6f2fd0"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3d56a7e865fcc245fee7dd5afef91958bd87a444"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
