package model

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

type TaskStore struct {
	db *bolt.DB
}

type Task struct {
	ID        int    `json: "id"`
	Task      string `json: "task"`
	Completed bool   `json: "completed"`
}

func StartServer() (*TaskStore, error) {
	db, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Check for bucket
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &TaskStore{db: db}, nil
}

func (ts *TaskStore) Close() {
	ts.db.Close()
}

func (ts *TaskStore) AddItem(task string) error {
	return ts.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))

		if b == nil {
			return fmt.Errorf("the bucket does not exist")
		}

		id, _ := b.NextSequence()

		item := Task{
			ID:        int(id),
			Task:      task,
			Completed: false,
		}

		buf, err := json.Marshal(item)
		if err != nil {
			return err
		}

		key := []byte(strconv.Itoa(item.ID))
		b.Put(key, buf)
		return nil
	})
}

func (ts *TaskStore) UpdateAction(id int) error {

	return ts.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))

		if b == nil {
			return fmt.Errorf("the bucket does not exist")
		}

		key := []byte(strconv.Itoa(id))

		taskJson := b.Get(key)

		if taskJson == nil {
			return fmt.Errorf("the task with the id does not exist")
		}

		var task Task
		if err := json.Unmarshal(taskJson, &task); err != nil {
			return err
		}

		task.Completed = true

		updatedJSON, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(key, updatedJSON)
	})

}

func (ts *TaskStore) ListTasks() ([]Task, error) {
	var task Task
	var tasks []Task
	err := ts.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))

		if b == nil {
			return fmt.Errorf("the bucket does not exist")
		}

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}
			if !task.Completed {
				tasks = append(tasks, task)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("there is some error fetching the list: %s", err)
	}

	return tasks, nil
}

func (ts *TaskStore) GetIncompleteList(index int) error {
	tasks, err := ts.ListTasks()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("index provided is greater than uncompleted list")
	}

	return ts.UpdateAction(tasks[index-1].ID)
}

func (ts *TaskStore) DeleteTask(index int) (*Task, error) {
	var task Task
	err := ts.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("Tasks"))
		if b == nil {
			return fmt.Errorf("the bucket does not exist")
		}

		key := []byte(strconv.Itoa(index))

		taskJson := b.Get(key)

		if err := json.Unmarshal(taskJson, &task); err != nil {
			return err
		}

		return b.Delete(key)
	})

	if err != nil {
		return nil, fmt.Errorf("there is some problem deleting - %s", err)
	}

	return &task, nil
}
