// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "26c956c0d6c2039697cd59ce64c2d64c636646172f47b082397aa90bad55a013"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "14da8232f7a42098b7ae460f4b4795f655afe66b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
