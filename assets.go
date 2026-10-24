// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e4996eef4cd621c529b22aa36a2f5f486e615b7d436980a6c736843638b6be84"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8bee74634f134de8215cbe2990d0c06f9f579535"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
