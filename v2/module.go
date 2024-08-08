package grid_cli

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// Module is roughly equivalent to an application in a microkernel system.
type Module interface {
	Accept(ctx context.Context, parms ...interface{}) (Message, error)
	HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error)
}

type LocalCacheModule struct {
	cacheDir string
}

func NewLocalCacheModule(cacheDir string) *LocalCacheModule {
	return &LocalCacheModule{cacheDir: cacheDir}
}

func (m *LocalCacheModule) Accept(ctx context.Context, parms ...interface{}) (Message, error) {
	// Logic to check if the module can handle the message, returning a promise of handling ability
	// XXX return a real promise
	return Message{
		Parms: []interface{}{true},
		Payload: map[string]interface{}{
			"message": "Acceptable",
		},
	}, nil
}

func (m *LocalCacheModule) HandleMessage(ctx context.Context, parms ...interface{}) ([]byte, error) {
	if len(parms) < 2 {
		return nil, errors.New("insufficient arguments")
	}

	promiseHash, ok1 := parms[0].([]byte)
	moduleHash, ok2 := parms[1].([]byte)
	if !ok1 || !ok2 {
		return nil, errors.New("invalid argument types")
	}

	cacheKey := constructCacheKey(promiseHash, moduleHash, parms[2:]...)
	return loadFromCache(cacheKey)
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

func loadFromCache(cacheKey string) ([]byte, error) {
	filePath := filepath.Join("cache", cacheKey)
	return afero.ReadFile(fs, filePath)
}

func saveToCache(cacheKey string, data []byte) error {
	cacheDir := "cache"
	if err := fs.MkdirAll(cacheDir, 0755); err != nil {
		return err
	}
	filePath := filepath.Join(cacheDir, cacheKey)
	return afero.WriteFile(fs, filePath, data, 0644)
}
