// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c833f5f1888c200ce6686b037161921387aa38d1fd49229563d81f61ea1841cb"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "351f35351fc07abf4b523e3574d5d204109e37a1"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
