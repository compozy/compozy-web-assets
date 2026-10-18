// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9827372f7afeb6e4d66b86ea6d6659b9dbc90ac83c417df6fc7caafa486d87fe"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "bec3ce3097911bc17c5769d97a6ebf79fa2f3f15"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
