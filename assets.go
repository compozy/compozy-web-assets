// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9c90b97e73ed6c708e1ab2b1223baeb470d130963ce1c9221f82084a423f9414"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "d581ee7d4c77b099e96c53ef42f742e8f0e101d2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
