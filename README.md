`timetrack` is supposed to be a simple time tracking tool based on plaintext. This tool can be used to log time spent
on various "buckets" of tasks in the current month. The goal for this tool is to be an easy-to-use and extensible CLI.

This tool is supposed to provide these things OOTB:

- [ ] Management of "buckets" of tasks
    - [x] Listing and adding these buckets
    - [ ] Totalling time spent on a bucket
- [ ] Tracking time spent during a month
    - [x] Adding/Updating time spent on a task
    - [ ] Listing tasks
        - [x] of the current month
        - [ ] of the current year
        - [ ] everything
    - [ ] Output formats
        - [x] a sane default terminal output
        - [ ] CSV

Installation:

```shell
go install github.com/realbucksavage/timetrack/cmd/timetrack@latest
```

Usage:

- `timetrack status`: check for errors and such.
- `timetrack b[ucket][s]`: bucket management
  - `timetrack b[ucket][s] list|ls`: show all buckets
- `timetrack t[ask][s]`: task management
  - `timetrack t[ask][s] list|ls`: list tasks of the current month
  - `timetrack t[ask][s] a[dd] <bucket> <task> <duration>`: log some time on a task

Examples:

```shell
# log some time on a task
timetrack task add 'go-dev' 'build the timetrack utility' 30m

# show tracked time of the current month
timetrack t ls

# show the help message
timetrack h
```
