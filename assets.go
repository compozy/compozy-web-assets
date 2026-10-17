// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ffda90838ee09bc54b073e76195440766f86a02a4628dd60c29fc105a9877718"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3b83024757c38bdca11ef94555fba9f6d2c6ab80"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
