// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "426d1bbc4004b17bc063d5e48caea6a025d707a74f693974032b90af20af7686"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6065799625ae03fb3c052721e414bccb72b0d692"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
