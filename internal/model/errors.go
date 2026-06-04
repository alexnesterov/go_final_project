package model

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrCreateTask = errors.New("create task")
	ErrUpdateTask = errors.New("update task")
	ErrDoneTask   = errors.New("done task")
)
