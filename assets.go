// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6e002ffa754533830f83dd75105ab4135542b37ba04518008a56e969211e9692"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "aba30d136e827ef8f51870f28e473ee83fc3cbbc"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
