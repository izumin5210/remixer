// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateHelloPayload interface {
	IsCreateHelloPayload()
}

type CreateHelloInput struct {
	Message string `json:"message"`
}

type CreateHelloSuccess struct {
	Hello *Hello `json:"hello"`
}

type HelloMessageInvalid struct {
	Hello   *Hello `json:"hello"`
	Message string `json:"message"`
}

func (HelloMessageInvalid) IsCreateHelloPayload() {}

type UserError struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

func (UserError) IsCreateHelloPayload() {}

