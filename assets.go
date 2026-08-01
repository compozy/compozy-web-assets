// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "426d1bbc4004b17bc063d5e48caea6a025d707a74f693974032b90af20af7686"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "226137ec68cedf19ad60ef29cd1f7230fbd3e576"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
