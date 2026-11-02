// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "421e723b6adc8e419121fe9dc77aa4ab0899dd1faa03eb00ac5228db5c758477"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "543f879012e9b658668603f5977aff4f35c250d1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
