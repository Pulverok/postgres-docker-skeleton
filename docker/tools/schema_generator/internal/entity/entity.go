package entity

type MigrateConfig struct {
	Migrations []map[string][]string `yaml:"migrations"`
}

type DBItem struct {
	Type string
	Path string
	Data []byte
}

type Entity struct {
	Name  string
	Files []DBItem
}

// Entities is a list of entities.
type Entities []Entity
