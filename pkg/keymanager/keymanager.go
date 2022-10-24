package keymanager

import (
	"crypto/rand"
	"fmt"
	"os"
	"sync"

	"github.com/gotti/meshover/spec"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// Manager manages api keys
type Manager interface {
	GetKeysDigest() *spec.AgentKeys
	AddKey(key *spec.AgentKey) error
	GenerateKey() (*spec.AgentKey, error)
	RemoveKey(keyID int64) error
	IsValid(key string) error
}

// APIKeys is a token struct
type APIKeys struct {
	//path is a file path to save token
	path string
	mtx  sync.Mutex
	keys *spec.AgentKeys
	//MD is a metadata for authorization to gRPC server
	MD metadata.MD
}

// NewAPIKeys loads old settings if exists.
func NewAPIKeys(path string) (*APIKeys, error) {
	ret := &APIKeys{path: path}
	if err := ret.loadKeys(); err != nil {
		return nil, fmt.Errorf("failed to load key, err=%w", err)
	}
	return ret, nil
}

// GetKeysDigest returns digest of keys
func (c *APIKeys) GetKeysDigest() (ret *spec.AgentKeys) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	ret.Keys = make([]*spec.AgentKey, 0)
	for _, k := range c.keys.Keys {
		ret.Keys = append(ret.Keys, &spec.AgentKey{Id: k.GetId(), Key: k.GetKey()[:5]})
	}
	return ret
}

// AddKey adds key
func (c *APIKeys) AddKey(key *spec.AgentKey) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, k := range c.keys.Keys {
		if k.GetId() == key.GetId() {
			return fmt.Errorf("already exists such id, id=%d", key.GetId())
		}
	}
	c.keys.Keys = append(c.keys.Keys, key)
	return c.savekeys()
}

// GenerateKey generate new key
func (c *APIKeys) GenerateKey() (*spec.AgentKey, error) {
	var maxid int64 = 0
	for _, k := range c.keys.Keys {
		if k.GetId() > maxid {
			maxid = k.GetId()
		}
	}
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("random number generation failed, err=%w", err)
	}
	k := &spec.AgentKey{Id: maxid + 1, Key: fmt.Sprintf("%x", buf)}
	c.keys.Keys = append(c.keys.Keys, k)
	c.savekeys()
	return k, nil
}

// RemoveKey removes specified key
func (c *APIKeys) RemoveKey(keyID int64) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for i, k := range c.keys.Keys {
		if k.GetId() == keyID {
			c.keys.Keys = append(c.keys.Keys[:i], c.keys.Keys[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no such id, id=%d", keyID)
}

// IsValid returns nil if provided key is valid
func (c *APIKeys) IsValid(key string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, k := range c.keys.Keys {
		if k.GetKey() == key {
			return nil
		}
	}
	return fmt.Errorf("not found, unauthorized")
}

// savekeys saves keys
func (c *APIKeys) savekeys() error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	buf, err := proto.Marshal(c.keys)
	if err != nil {
		return fmt.Errorf("failed to marshal for saving token, err=%w", err)
	}
	if err := os.WriteFile(c.path, buf, 0600); err != nil {
		return fmt.Errorf("failed to write token, err=%w", err)
	}
	return nil
}

// loadKeys loads keys
func (c *APIKeys) loadKeys() error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if _, err := os.Stat(c.path); os.IsNotExist(err) {
		f, err := os.OpenFile(c.path, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return fmt.Errorf("failed to create key file")
		}
		defer f.Close()
	}
	f, err := os.ReadFile(c.path)
	if err != nil {
		return fmt.Errorf("failed to load key, err=%w", err)
	}
	buf := new(spec.AgentKeys)
	if err := proto.Unmarshal(f, buf); err != nil {
		return fmt.Errorf("failed to unmarshal for loading keys, err=%w", err)
	}
	c.keys = buf
	return nil
}
