package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"testdata/task/graph/generated"
	"testdata/task/graph/model"
)

func (r *taskResolver) ID(ctx context.Context, obj *model.Task) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *taskResolver) AssigneeIds(ctx context.Context, obj *model.Task) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *taskResolver) Id(ctx context.Context, obj *model.Task) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
