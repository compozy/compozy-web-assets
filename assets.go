// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dcb11ce25c2b8786689c7668c4c32c1effbc9ffc0569eacfc0acc451884634b8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e226c9ccf5dc9e27106cf6c70f786a8c818b24b8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
