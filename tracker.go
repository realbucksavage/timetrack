package timetrack

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/pkg/errors"
)

var (
	homeDir        = os.Getenv("HOME")
	defaultBaseDir = fmt.Sprintf("%s/.local/share/timetrack", homeDir)
)

type Option func(t *Tracker) error

func WithBaseDir(baseDir string) Option {
	return func(t *Tracker) error {
		if baseDir == "" {
			return errors.New("base directory must not be empty")
		}

		t.baseDir = baseDir
		return nil
	}
}

func NewTracker(opts ...Option) (*Tracker, error) {
	t := &Tracker{baseDir: defaultBaseDir}
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return nil, err
		}
	}

	info, err := os.Stat(t.baseDir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(t.baseDir, 0755)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot create base directory %q", t.baseDir)
		}
	}

	if info != nil && !info.IsDir() {
		return nil, errors.Errorf("%q is not a directory", t.baseDir)
	}

	return t, nil
}

type Tracker struct {
	baseDir string
}

func (tracker *Tracker) Status() error {
	buckets, err := tracker.ListBuckets()
	if err != nil {
		return errors.Wrap(err, "reading buckets")
	}

	fmt.Printf("%d buckets", len(buckets))

	return nil
}

func (tracker *Tracker) pathTo(file string) string {
	return path.Join(tracker.baseDir, file)
}

func mustClose(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}
