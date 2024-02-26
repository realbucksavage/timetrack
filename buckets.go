package timetrack

import (
	"bufio"
	"os"
)

func (tracker *Tracker) bucketFile() (*os.File, error) {
	location := tracker.pathTo("buckets")
	file, err := os.Open(location)
	if os.IsNotExist(err) {
		file, err = os.OpenFile(location, os.O_CREATE|os.O_RDWR, 644)
	}

	return file, nil
}

func (tracker *Tracker) ListBuckets() ([]string, error) {
	file, err := tracker.bucketFile()
	if err != nil {
		return nil, err
	}

	defer must(file.Close())

	buckets := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		buckets = append(buckets, scanner.Text())
	}

	return buckets, nil
}
