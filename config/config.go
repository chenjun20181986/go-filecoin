package config

import (
	"os"

	"gx/ipfs/QmWHbPAp5UWfwZE3XCgD93xsCYZyk12tAAQVL3QXLKcWaj/toml"
)

// Config is an in memory representation of the filecoin configuration file
type Config struct {
	API       *APIConfig       `toml:"api"`
	Bootstrap *BootstrapConfig `toml:"bootstrap"`
	Datastore *DatastoreConfig `toml:"datastore"`
	Swarm     *SwarmConfig     `toml:"swarm"`
}

// APIConfig holds all configuration options related to the api.
type APIConfig struct {
	Address                  string   `toml:"address"`
	AccessControlAllowOrigin []string `toml:"accessControlAllowOrigin"`
}

func newDefaultAPIConfig() *APIConfig {
	return &APIConfig{
		Address:                  ":3453",
		AccessControlAllowOrigin: []string{"http://localhost:8080"},
	}
}

// DatastoreConfig holds all the configuration options for the datastore.
// TODO: use the advanced datastore configuration from ipfs
type DatastoreConfig struct {
	Type string `toml:"type"`
	Path string `toml:"path"`
}

func newDefaultDatastoreConfig() *DatastoreConfig {
	return &DatastoreConfig{
		Type: "badgerds",
		Path: "badger",
	}
}

// SwarmConfig holds all configuration options related to the swarm.
type SwarmConfig struct {
	Address string `toml:"address"`
}

func newDefaultSwarmConfig() *SwarmConfig {
	return &SwarmConfig{
		Address: "/ip4/127.0.0.1/tcp/6000",
	}
}

// BootstrapConfig holds all configuration options related to bootstrap nodes
type BootstrapConfig struct {
	Addresses []string `toml:"addresses"`
}

// TODO: define default bootstrap node addresses
func newDefaultBootstrapConfig() *BootstrapConfig {
	return &BootstrapConfig{
		Addresses: []string{
			"TODO",
		},
	}
}

// NewDefaultConfig returns a config object with all the fields filled out to
// their default values
func NewDefaultConfig() *Config {
	return &Config{
		API:       newDefaultAPIConfig(),
		Bootstrap: newDefaultBootstrapConfig(),
		Datastore: newDefaultDatastoreConfig(),
		Swarm:     newDefaultSwarmConfig(),
	}
}

// WriteFile writes the config to the given filepath.
func (cfg *Config) WriteFile(file string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	if err := toml.NewEncoder(f).Encode(*cfg); err != nil {
		return err
	}

	return f.Close()
}

// ReadFile reads a config file from disk.
func ReadFile(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	cfg := NewDefaultConfig()
	if _, err := toml.DecodeReader(f, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}