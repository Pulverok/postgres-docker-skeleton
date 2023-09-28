package processor

import (
	"fmt"
	"os"

	"github.com/Pulverok/postgres-docker-skeleton/schema_generator/internal/config"
	"github.com/Pulverok/postgres-docker-skeleton/schema_generator/internal/entity"
	"github.com/Pulverok/postgres-docker-skeleton/schema_generator/internal/enum"

	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

// Processor handles the processing of entities and generates migration SQL files.
type Processor struct {
	cfg *config.AppConfig
}

// New creates a new Processor instance.
func New(cfg *config.AppConfig) *Processor {
	return &Processor{cfg: cfg}
}

// Process generates migration SQL files for the configured entities.
func (p *Processor) Process() error {
	mEntities, err := p.getEntities()
	if err != nil {
		return err
	}

	if len(mEntities) == 0 {
		return fmt.Errorf("no entities to import")
	}

	res, err := os.Create(fmt.Sprintf("%s/%s", p.cfg.DataPath, "migrate.sql"))
	if err != nil {
		return err
	}
	defer res.Close()

	for _, e := range mEntities {
		for _, f := range e.Files {
			_, err := res.Write(f.Data)
			if err != nil {
				return err
			}
			_, err = res.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// getMigrateConfig retrieves the migration configuration from the config file.
func (p *Processor) getMigrateConfig() (*entity.MigrateConfig, error) {
	cfg := entity.MigrateConfig{}
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", p.cfg.DataPath, "config.yml"))
	if err != nil {
		return nil, fmt.Errorf("can't read config file: %w", err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// getEntities retrieves entities from the MigrateConfig and initializes their data.
func (p *Processor) getEntities() (entity.Entities, error) {
	var entities entity.Entities
	migrateEntities, err := p.getMigrateConfig()
	if err != nil {
		return nil, err
	}

	schemas, err := getSchemaFiles(*migrateEntities, p.cfg.DataPath)
	if err != nil {
		return nil, err
	}
	entities = append(entities, entity.Entity{
		Name:  "schemas",
		Files: schemas,
	})

	for _, v := range migrateEntities.Migrations {
		for k, l := range v {
			item := getEntity(l, k, p.cfg.DataPath)
			entities = append(entities, *item)
		}
	}

	return entities, nil
}

// getEntity retrieves an entity's data including its schema and SQL files.
func getEntity(e []string, schema, dataPath string) *entity.Entity {
	var item entity.Entity
	item.Name = schema

	for _, el := range e {
		for _, it := range enum.AllTypes {
			path := fmt.Sprintf("%s/%s/%s/%s.sql", dataPath, schema, it, el)
			content, err := os.ReadFile(path)
			if err == nil {
				item.Files = append(item.Files, entity.DBItem{
					Type: string(it),
					Path: path,
					Data: content,
				})
			}
		}
	}

	return &item
}

func getSchemaFiles(cfg entity.MigrateConfig, dataPath string) ([]entity.DBItem, error) {
	var schemasPath []string
	var schemas []entity.DBItem
	for _, e := range cfg.Migrations {
		for k := range e {
			if k != "public" {
				schemaPath := fmt.Sprintf("%s/%s/initialize-schema.sql", dataPath, k)
				content, err := os.ReadFile(schemaPath)
				if err != nil {
					return nil, fmt.Errorf("schema file %s not found", k)
				}

				if !slices.Contains(schemasPath, schemaPath) {
					schemas = append(schemas, entity.DBItem{
						Type: string(enum.Schema),
						Path: schemaPath,
						Data: content,
					})

					schemasPath = append(schemasPath, schemaPath)
				}
			}
		}
	}

	return schemas, nil
}
