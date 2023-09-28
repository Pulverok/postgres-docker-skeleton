package enum

import "github.com/Pulverok/postgres-docker-skeleton/schema_generator/internal/entity"

const (
	Schema            entity.Type = "schema"
	Types             entity.Type = "types"
	Enums             entity.Type = "enums"
	Tables            entity.Type = "tables"
	Sequences         entity.Type = "sequences"
	Seeds             entity.Type = "seeds"
	Fixtures          entity.Type = "fixtures"
	Functions         entity.Type = "functions"
	MaterializedViews entity.Type = "materialized_views"
	Views             entity.Type = "views"
)

// AllTypes is a list of all entities.
var AllTypes = []entity.Type{
	Types,
	Enums,
	Tables,
	Sequences,
	Seeds,
	Fixtures,
	Functions,
	MaterializedViews,
	Views,
}
