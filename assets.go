// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c596e50caa91fc4379cdaec33d889ad377a06d7b58e8cd56f9427d466a1804d1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0385983a40187815751801db67eb99f611327f36"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
