package timetrack

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type Task struct {
	Task  string
	Spent time.Duration
}

func (tracker *Tracker) tasksFile(flags int) (*os.File, error) {
	now := time.Now()
	year := fmt.Sprintf("%d", now.Year())
	baseDir := tracker.pathTo(year)

	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		return nil, err
	}

	month := fmt.Sprintf("%d", now.Month())
	location := path.Join(baseDir, month)
	return os.OpenFile(location, flags, 0644)
}

func (tracker *Tracker) Track(bucket, task string, duration time.Duration) (*Task, error) {
	err := tracker.AddBucket(bucket)
	if err != nil {
		return nil, err
	}

	tasks, err := tracker.ListTasks()
	if err != nil {
		return nil, err
	}

	var match *Task
	for _, t := range tasks[bucket] {
		if !strings.EqualFold(t.Task, task) {
			continue
		}

		match = t
		break
	}

	if match != nil {
		match.Spent += duration
	} else {
		match = &Task{
			Task:  task,
			Spent: duration,
		}

		list, ok := tasks[bucket]
		if !ok {
			list = []*Task{match}
		} else {
			list = append(list, match)
		}

		tasks[bucket] = list
	}

	file, err := tracker.tasksFile(os.O_CREATE | os.O_RDWR)
	if err != nil {
		return nil, err
	}

	defer mustClose(file)

	if err := file.Truncate(0); err != nil {
		return nil, err
	}

	if err := toml.NewEncoder(file).Encode(tasks); err != nil {
		return nil, err
	}

	return match, nil
}

func (tracker *Tracker) ListTasks() (map[string][]*Task, error) {

	file, err := tracker.tasksFile(os.O_CREATE | os.O_RDONLY)
	if err != nil {
		return nil, err
	}

	defer mustClose(file)

	tasks := make(map[string][]*Task)
	if _, err := toml.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
