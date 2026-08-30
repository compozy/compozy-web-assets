// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3c929e9c505913977fb5ac13ed36ab311e1887197b4228215c6c2e3b4f0758ad"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "98e4aa7c81f0726916c01c4a08d153d6b924be69"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
