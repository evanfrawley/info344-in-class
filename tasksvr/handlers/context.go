package handlers

import "github.com/evanfrawley/info344-in-class/tasksvr/models/tasks"

//Context holds context values
//used by multiple handler functions.
type Context struct {
	//TODO: add a field that will hold
	//a tasks.Store implementation.
	//Our handlers will use this to
	//insert, update, and get tasks
	tasksStore tasks.Store
}

func NewHandlerContext(tasksStore tasks.Store) *Context {
    return &Context {
        tasksStore: tasksStore,
    }
}