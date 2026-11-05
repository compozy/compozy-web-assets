// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "06c1dc09320a97ab5fdcc06b9d50f462e57ff03962e4eab6bfad77c90cb05aa6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6f7d828f14a1af5f86b1e2689f354327c4c21cfc"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
