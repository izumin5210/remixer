package gqls

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/izumin5210/remixer/cmd/protoc-gen-graphql/protoutil"
)

var (
	_ Type      = (*InputObjectType)(nil)
	_ Definable = (*InputObjectType)(nil)
)

func NewInputObjectType(objType *ObjectType) *InputObjectType {
	return &InputObjectType{base: objType}
}

type InputObjectType struct {
	base *ObjectType
}

func (t *InputObjectType) Name() string {
	n := t.base.Name()
	n = strings.TrimSuffix(n, "Request")
	return n + "Input"
}

func (t *InputObjectType) IsNullable() bool                         { return t.base.IsNullable() }
func (t *InputObjectType) IsList() bool                             { return t.base.IsList() }
func (t *InputObjectType) TypeAST() *ast.Type                       { return ast.NonNullNamedType(t.Name(), nil) }
func (t *InputObjectType) ProtoDescriptor() protoreflect.Descriptor { return t.base.Proto }

func (t *InputObjectType) DefinitionAST() (*ast.Definition, error) {
	def := &ast.Definition{
		Kind:       ast.InputObject,
		Name:       string(t.Name()),
		Directives: definitionDelectivesAST(t.base.Proto),
	}

	err := protoutil.RangeFields(t.base.Proto, func(fd protoreflect.FieldDescriptor) error {
		ft, err := TypeFromProtoField(fd)
		if err != nil {
			return err
		}
		origType := ft
		for {
			if mt, ok := origType.(ModifiedType); ok {
				origType = mt.Original()
			} else {
				break
			}
		}
		if ot, ok := origType.(*ObjectType); ok {
			origType = NewInputObjectType(ot)
		}
		if ft.IsNullable() {
			origType = NullableType(origType)
		}
		if ft.IsList() {
			origType = ListType(origType)
		}
		ft = origType
		def.Fields = append(def.Fields, &ast.FieldDefinition{
			Name: strcase.ToLowerCamel(string(fd.Name())),
			Type: ft.TypeAST(),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return def, nil
}