package protomodelgen_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/izumin5210/remixer/gqlgenplugin/protomodelgen"
	"github.com/izumin5210/remixer/gqlgentest"
)

func TestGenerateCode(t *testing.T) {
	rootDir := getModuleRoot()
	testdataDir := filepath.Join(rootDir, "testdata")

	gqlgentest := gqlgentest.New(t)
	gqlgentest.AddGqlGenPlugin(protomodelgen.New())
	gqlgentest.AddGqlSchemaFile(t, filepath.Join(testdataDir, "apis", "graphql", "task", "*.graphqls"))
	gqlgentest.AddGqlSchema("schema.graphqls", `
directive @grpc(service: String!, rpc: String!) on FIELD_DEFINITION
directive @proto(fullName: String!, package: String!, name: String!, goPackage: String!, goName: String!) on OBJECT | INPUT_OBJECT | ENUM
directive @protoField(name: String!, type: String!, goName: String!, goTypeName: String!, goTypePackage: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION

extend type Query {
  tasks: [Task!]!
}`)
	gqlgentest.AddGoModReplace("github.com/izumin5210/remixer", rootDir)
	gqlgentest.AddGoModReplace("apis/go/tasks", filepath.Join(testdataDir, "apis", "go", "task"))

	gqlgentest.Run(t, func(t *testing.T, err error) {
		if err != nil {
			t.Errorf("failed to generate code: %v", err)
		}
		gqlgentest.SnapshotFile(t,
			"model/protomodels_gen.go",
		)
	})
}

func getModuleRoot() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
}