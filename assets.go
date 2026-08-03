// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c596e50caa91fc4379cdaec33d889ad377a06d7b58e8cd56f9427d466a1804d1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "dc65b22c79a9bf82a43b2fd8b8a33146dfd7f87d"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
