// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8358df7149ec28402ef3f8565b0f11b6e5dbba5285a6600171341818efd3b861"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6559b7d2f1b7e0e8d44bcbfacc5bb1cb40979432"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
