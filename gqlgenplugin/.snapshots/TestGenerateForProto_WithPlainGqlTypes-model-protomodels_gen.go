// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	todo_pb "apis/go/todo"
	"fmt"
	"io"
	"strconv"
)

type Task struct {
	ID uint64

	Title string

	Status *TaskStatus

	AssigneeIds []uint64

	AuthorID uint64
}

func TaskListFromRepeatedProto(in []*todo_pb.Task) []*Task {
	out := make([]*Task, len(in))
	for i, m := range in {
		out[i] = TaskFromProto(m)
	}
	return out
}

func TaskFromProto(in *todo_pb.Task) *Task {
	return &Task{

		ID: in.Id,

		Title: in.Title,

		Status: TaskStatusFromProto(in.Status),

		AssigneeIds: in.AssigneeIds,

		AuthorID: in.AuthorId,
	}
}

func TaskListToRepeatedProto(in []*Task) []*todo_pb.Task {
	out := make([]*todo_pb.Task, len(in))
	for i, m := range in {
		out[i] = TaskToProto(m)
	}
	return out
}

func TaskToProto(in *Task) *todo_pb.Task {
	return &todo_pb.Task{

		Id: in.ID,

		Title: in.Title,

		Status: TaskStatusToProto(in.Status),

		AssigneeIds: in.AssigneeIds,

		AuthorId: in.AuthorID,
	}
}

type TaskInput struct {
	ID uint64

	Title string

	Status *TaskStatus

	AssigneeIds []uint64

	AuthorID uint64
}

func TaskInputListFromRepeatedProto(in []*todo_pb.Task) []*TaskInput {
	out := make([]*TaskInput, len(in))
	for i, m := range in {
		out[i] = TaskInputFromProto(m)
	}
	return out
}

func TaskInputFromProto(in *todo_pb.Task) *TaskInput {
	return &TaskInput{

		ID: in.Id,

		Title: in.Title,

		Status: TaskStatusFromProto(in.Status),

		AssigneeIds: in.AssigneeIds,

		AuthorID: in.AuthorId,
	}
}

func TaskInputListToRepeatedProto(in []*TaskInput) []*todo_pb.Task {
	out := make([]*todo_pb.Task, len(in))
	for i, m := range in {
		out[i] = TaskInputToProto(m)
	}
	return out
}

func TaskInputToProto(in *TaskInput) *todo_pb.Task {
	return &todo_pb.Task{

		Id: in.ID,

		Title: in.Title,

		Status: TaskStatusToProto(in.Status),

		AssigneeIds: in.AssigneeIds,

		AuthorId: in.AuthorID,
	}
}

type CreateTaskPayload_Proto struct {
	Task *todo_pb.Task
}

func CreateTaskPayloadListFromRepeatedProto(in []*CreateTaskPayload_Proto) []*CreateTaskPayload {
	out := make([]*CreateTaskPayload, len(in))
	for i, m := range in {
		out[i] = CreateTaskPayloadFromProto(m)
	}
	return out
}

func CreateTaskPayloadFromProto(in *CreateTaskPayload_Proto) *CreateTaskPayload {
	return &CreateTaskPayload{

		Task: TaskFromProto(in.Task),
	}
}

func CreateTaskPayloadListToRepeatedProto(in []*CreateTaskPayload) []*CreateTaskPayload_Proto {
	out := make([]*CreateTaskPayload_Proto, len(in))
	for i, m := range in {
		out[i] = CreateTaskPayloadToProto(m)
	}
	return out
}

func CreateTaskPayloadToProto(in *CreateTaskPayload) *CreateTaskPayload_Proto {
	return &CreateTaskPayload_Proto{

		Task: TaskToProto(in.Task),
	}
}

type TaskStatus struct {
	Proto todo_pb.Task_Status
}

func TaskStatusListFromRepeatedProto(in []todo_pb.Task_Status) []*TaskStatus {
	out := make([]*TaskStatus, len(in))
	for i, m := range in {
		out[i] = TaskStatusFromProto(m)
	}
	return out
}

func TaskStatusFromProto(in todo_pb.Task_Status) *TaskStatus {
	return &TaskStatus{Proto: in}
}

func TaskStatusListToRepeatedProto(in []*TaskStatus) []todo_pb.Task_Status {
	out := make([]todo_pb.Task_Status, len(in))
	for i, m := range in {
		out[i] = TaskStatusToProto(m)
	}
	return out
}

func TaskStatusToProto(in *TaskStatus) todo_pb.Task_Status {
	return in.Proto
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.Proto.String()))
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	e.Proto = todo_pb.Task_Status(todo_pb.Task_Status_value[str])
	return nil
}

