// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ca2c688b34c814f9e8d1e6a0726fea8cfa5c82706538c5399e40dbf0d4e30a57"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3f85e8c52716e15616d1847fb2ea3e1aff0d1199"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
