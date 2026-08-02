// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "85f38aba3a431c29dc1f7626ae570ef9cfde885cb6620a0fda55f8ca9edd134d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4f4992466423c5562bd30e96d27846684fb377d6"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
