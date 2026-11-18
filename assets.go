// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "857c8e5428f782916f401f03d9f8ed2c7abc86a72599c0402a58ecca32f34d81"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fcf81bb4f744e4679c7827806b4df0d3e9420fce"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
