// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f4e2284924cbb81a12c18fb70efa9bf7d95f4246e2ce98aebd67e3b327b0433a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0baeae0cacb7c20777f1cd990d9d61422c062817"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
