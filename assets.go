// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b4d4be90518e13f75541249107c712dfb1602d7913f15e2530e66f92cd27c393"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "30da9b257bcb7e7dba23b8616dc7e9767fc7f08a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
