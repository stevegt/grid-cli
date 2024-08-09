package grid_cli

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Module is an interface for grid-cli modules.
type Module interface {
	Accept(ctx context.Context, parms ...interface{}) (Message, error)
	HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error)
}

// LocalCacheModule implements a module that interacts with a local cache.
type LocalCacheModule struct {
	cacheDir string
}

func NewLocalCacheModule(cacheDir string) *LocalCacheModule {
	return &LocalCacheModule{cacheDir: cacheDir}
}

func constructCacheKey(promiseHash, moduleHash []byte, args ...interface{}) string {
	var keyBuilder strings.Builder

	keyBuilder.WriteString(fmt.Sprintf("%x", promiseHash))
	keyBuilder.WriteString("/")
	keyBuilder.WriteString(fmt.Sprintf("%x", moduleHash))

	for _, arg := range args {
		var encodedArg string
		switch v := arg.(type) {
		case string:
			encodedArg = url.QueryEscape(v)
		case []byte:
			encodedArg = url.QueryEscape(string(v))
		default:
			// Handle unsupported argument types
		}
		keyBuilder.WriteString("/")
		keyBuilder.WriteString(encodedArg)
	}

	return keyBuilder.String()
}
