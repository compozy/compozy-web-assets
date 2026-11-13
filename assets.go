// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1f878a294f59e2f35233a8d53d23616a5a52bc2b63cc81d407cfda9bc5ca13a5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f6b76ef54f5de2b38884cd6a510c488af2e4505d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
