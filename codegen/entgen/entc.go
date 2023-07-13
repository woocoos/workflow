//go:build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/woocoos/entco/genx"
	"log"
	"os"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("codegen/gqlgen/gqlgen.yaml"),
		entgql.WithSchemaPath("api/graphql/ent.graphql"),
		entgql.WithSchemaHook(genx.ChangeRelayNodeType()),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	os.MkdirAll("./api/graphql", os.ModePerm)
	opts := []entc.Option{
		entc.Extensions(ex),
		genx.GlobalID(),
		genx.SimplePagination(),
	}
	err = entc.Generate("./codegen/entgen/schema", &gen.Config{
		Package: "github.com/woocoos/workflow/ent",
		Features: []gen.Feature{gen.FeatureVersionedMigration, gen.FeatureUpsert,
			gen.FeatureIntercept, gen.FeatureSchemaConfig},
		Target: "./ent",
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
