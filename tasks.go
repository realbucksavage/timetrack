package timetrack

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type TaskFile struct {
	Tasks []*Task
}

type Task struct {
	Task   string
	Bucket string
	Spent  time.Duration
}

func (t *Task) String() string {
	return fmt.Sprintf("%s: %s (%v)", t.Bucket, t.Task, t.Spent)
}

func (tracker *Tracker) taskBaseDir() string {
	now := time.Now()
	base := path.Join(fmt.Sprintf("%d", now.Year()), fmt.Sprintf("%d", now.Month()))

	return tracker.pathTo(base)
}

func (tracker *Tracker) tasksFile(flags int) (*os.File, error) {
	baseDir := tracker.taskBaseDir()
	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		return nil, err
	}

	location := path.Join(baseDir, "tasks")
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
	for _, t := range tasks {
		if !strings.EqualFold(t.Task, task) && !strings.EqualFold(t.Bucket, bucket) {
			continue
		}

		match = t
		break
	}

	if match != nil {
		match.Spent += duration
	} else {
		match = &Task{
			Task:   task,
			Bucket: bucket,
			Spent:  duration,
		}

		tasks = append(tasks, match)
	}

	file, err := tracker.tasksFile(os.O_CREATE | os.O_RDWR)
	if err != nil {
		return nil, err
	}

	defer mustClose(file)

	if err := toml.NewEncoder(file).Encode(&TaskFile{tasks}); err != nil {
		return nil, err
	}

	return match, nil
}

func (tracker *Tracker) ListTasks() ([]*Task, error) {

	file, err := tracker.tasksFile(os.O_CREATE | os.O_RDONLY)
	if err != nil {
		return nil, err
	}

	defer mustClose(file)

	tasks := new(TaskFile)
	if _, err := toml.NewDecoder(file).Decode(tasks); err != nil {
		return nil, err
	}

	return tasks.Tasks, nil
}
