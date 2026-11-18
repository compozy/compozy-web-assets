// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "857c8e5428f782916f401f03d9f8ed2c7abc86a72599c0402a58ecca32f34d81"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e649e96abb6d0357438168450ed4bec35ac63230"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
