// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e248947cf62f59669e39077fd80b5462d8824f594ec7c35e8df7ffdf60f379ce"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4dbd3ad5a0c10cba1b39639d23d44c98890ec5f0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
