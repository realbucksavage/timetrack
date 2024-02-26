package timetrack

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func (tracker *Tracker) bucketFile(flags int) (*os.File, error) {
	location := tracker.pathTo("buckets")
	return os.OpenFile(location, flags, 0644)
}

func (tracker *Tracker) ListBuckets() ([]string, error) {
	file, err := tracker.bucketFile(os.O_CREATE | os.O_RDONLY)
	if err != nil {
		return nil, err
	}

	defer mustClose(file)

	buckets := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		buckets = append(buckets, scanner.Text())
	}

	return buckets, nil
}

func (tracker *Tracker) AddBucket(bucket string) error {
	buckets, err := tracker.ListBuckets()
	if err != nil {
		return err
	}

	if slices.Contains(buckets, bucket) {
		return nil
	}

	file, err := tracker.bucketFile(os.O_CREATE | os.O_RDWR | os.O_APPEND)
	if err != nil {
		return err
	}

	defer mustClose(file)

	_, err = fmt.Fprintln(file, bucket)
	return err
}
