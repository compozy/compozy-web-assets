// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "636b8e76ff451b4779203e33d36d44ff6cfa5253d8011738dbff618450451516"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "701e565c534365b9cb0e9f4ed53a18f76d9d4ff4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
