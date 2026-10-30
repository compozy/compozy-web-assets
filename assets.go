// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "26e96f7af973c1422e8d42c7c346bbe3cd0c5897e9054ef240111f8f1f3ae42b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ab2a55309b6e48adb68cb99f8a6123d8b062e27b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
