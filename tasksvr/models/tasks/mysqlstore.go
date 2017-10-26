package tasks

import (
	"database/sql"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//various SQL statements we will need to execute
//TODO: write me some SQL! Refer to sql/schema.sql
//for the table/column names

//SQL to insert a new task row
//use `?` for column values that we will get at runtime
const sqlInsertTask = `insert into tasks(id,title,completed,createdAt) values (?,?,?,?)`

//SQL to insert a tag for a task
const sqlInsertTag = `insert into tags(taskID,tag) values (?,?)`

//SQL to select all tasks/tags with a particular task.completed value
//join tasks to tags so we get everything with only one query
const sqlSelectAll = ``

//SQL to select a particular task/tags by id
//join tasks to tags so we get everything with only one query
const sqlSelectID = ``

//SQL to update task.completed and task.modifiedAt where id = ?
const sqlUpdate = ``

//taskRow represents the data we will get for each row
//from the sqlSelectAll and selSelectID join queries
type taskRow struct {
	id         string
	title      string
	completed  bool
	createdAt  time.Time
	modifiedAt *time.Time
	tag        string
}

//MySQLStore implements Store for a MySQL database
type MySQLStore struct {
	db *sql.DB
}

//NewMySQLStore constructs a MySQLStore
func NewMySQLStore(db *sql.DB) *MySQLStore {
	if db == nil {
		panic("nil pointer passed to NewMySQLStore")
	}

	return &MySQLStore{
		db: db,
	}
}

//Insert inserts a new task into the database
func (s *MySQLStore) Insert(nt *NewTask) (*Task, error) {
	task, err := nt.ToTask()
	if err != nil {
		return nil, err
	}
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(sqlInsertTask, task.ID, task.Title, task.Completed, task.CreatedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, tag := range task.Tags {
		_, err = tx.Exec(sqlInsertTag, task.ID, tag)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}


	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return task, nil
}

//GetAll gets all tasks with a particular completed state
func (s *MySQLStore) GetAll(completed bool) ([]*Task, error) {
	panic("TODO:")
}

//Update applies the values in `tu` to the task with `id`
func (s *MySQLStore) Update(id bson.ObjectId, tu *TaskUpdates) (*Task, error) {
	panic("TODO:")
}

//scanTasks scans query result rows into a []*Task.
//Since the sqlSelectAll and sqlSelectID queries do a
//join between tasks and tags, we will get multiple rows
//for a single task if it has multiple tags. This function
//iterates over the result rows, constructing a *Task
//as it finds new task IDs, else it appends the next
//tag to the last-created Task.
func scanTasks(rows *sql.Rows) ([]*Task, error) {
	panic("TODO:")
}
