// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5ba414effd1ceafb6dbf8f72b1dea4aaf70fa8c044d5f8715adad4b0d92e4aea"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "23b1626ae659f337c3d48a05501b26128d309893"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
