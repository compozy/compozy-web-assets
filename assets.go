// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b9520189ff411614b4b51ba84eb2be81f751717ee8d67eb9d045f78ac7c755fc"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a00a9df50ed056e41224cc24bc4299ca7e78de41"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
