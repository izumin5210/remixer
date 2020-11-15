package gqls

import (
	"github.com/iancoleman/strcase"
	"github.com/izumin5210/remixer/codegen/protoutil"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	_ Type      = (*ObjectType)(nil)
	_ Definable = (*ObjectType)(nil)
	_ ProtoType = (*ObjectType)(nil)
)

type ObjectType struct {
	Proto *protogen.Message
}

func newObjectType(proto *protogen.Message) *ObjectType        { return &ObjectType{Proto: proto} }
func (t *ObjectType) Name() string                             { return nameWithParent(t.Proto.Desc) }
func (t *ObjectType) IsNullable() bool                         { return false }
func (t *ObjectType) IsList() bool                             { return false }
func (t *ObjectType) TypeAST() *ast.Type                       { return ast.NonNullNamedType(t.Name(), nil) }
func (t *ObjectType) ProtoDescriptor() protoreflect.Descriptor { return t.Proto.Desc }
func (t *ObjectType) GoIdent() protogen.GoIdent                { return t.Proto.GoIdent }

func (t *ObjectType) DefinitionAST() (*ast.Definition, error) {
	def := &ast.Definition{
		Kind:        ast.Object,
		Name:        string(t.Name()),
		Directives:  messageDirectivesAST(t.Proto),
		Description: protoutil.FormatComments(t.Proto.Comments),
	}

	for _, f := range t.Proto.Fields {
		ft, err := TypeFromProtoField(f)
		if err != nil {
			return nil, err
		}
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name:        strcase.ToLowerCamel(string(f.Desc.Name())),
			Type:        ft.TypeAST(),
			Directives:  fieldDirectivesAST(f, ft),
			Description: protoutil.FormatComments(f.Comments),
		})
	}

	return def, nil
}
