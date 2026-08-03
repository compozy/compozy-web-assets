// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8f3678454b0f421f79a61aac9de09fa6289d2d80097af3ea280e19f0335fd268"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "530a78dff0cd719eb4be660c4484fa45a4158721"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
