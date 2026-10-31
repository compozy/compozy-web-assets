// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e248947cf62f59669e39077fd80b5462d8824f594ec7c35e8df7ffdf60f379ce"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "be0d735b4113ef3ed91f6e2068c756535a4eabab"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
