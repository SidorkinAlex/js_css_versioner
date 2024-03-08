package version_replace

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/SidorkinAlex/js_css_versioner/internal/models"
)

// VersionReplace
type VersionReplace struct {
	rgExp *regexp.Regexp
}

// New constructs VersionReplace
func New(exts ...models.VersionedExtension) (*VersionReplace, error) {
	if len(exts) == 0 {
		return nil, fmt.Errorf("exts is required argument")
	}

	extsStr := make([]string, len(exts))
	for i := range exts {
		if exts[i] == "" {
			return nil, fmt.Errorf("versioned extension cannot be empty string")
		}
		extsStr[i] = string(exts[i])
	}

	rgExp, errReg := regexp.Compile(fmt.Sprintf("\\.(%s)([v0-9\\?\"|\\'=]+)", strings.Join(extsStr, "|")))
	if errReg != nil {
		return nil, fmt.Errorf("failed to compile regexp from exts: %w", errReg)
	}

	return &VersionReplace{
		rgExp: rgExp,
	}, nil
}

// Execute reads content of source replace version for all versioned extensions and writes content to destination
func (v *VersionReplace) Execute(source io.Reader, destination io.Writer, version int64) error {
	srcContent, errRead := io.ReadAll(source)
	if errRead != nil {
		return fmt.Errorf("VersionReplaceService.Replace failed source read: %w", errRead)
	}

	dstContent := v.rgExp.ReplaceAllFunc(srcContent, func(b []byte) []byte {
		closingQuote := b[len(b)-1]
		extensionGroup := v.rgExp.ReplaceAllString(string(b), `$1`)
		return []byte(fmt.Sprintf(".%s?v=%d%c", extensionGroup, version, closingQuote))
	})

	nWrite, errWrite := destination.Write(dstContent)
	if errWrite != nil {
		return fmt.Errorf("VersionReplaceService.Replace failed destination write: %w", errRead)
	}
	log.Printf("successfully write %d bytes to destination\n", nWrite)

	return nil
}
