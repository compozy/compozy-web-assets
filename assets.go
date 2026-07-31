// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "70bd06ec4b84b51fec75ce74ef9773792a01e97f2b7328eb6b90788ccaefc5de"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7f3fb64152febb569ba6e0627f1013d687db902f"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
