// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "018bda38be2ce1cbd7a606cc1504662f552b93747d92c8055a868f5ad846d473"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "928848ba47f8ce3f8f37ccc95a971c2c50545ac7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
