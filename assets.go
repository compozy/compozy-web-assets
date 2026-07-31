// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "75b3cf99f875225794755d309906e5fc96705f0f76fe3471d23b7328e81140f0"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "599b32ddc9ef92cfe1bcc2f939403a8a8911018b"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
