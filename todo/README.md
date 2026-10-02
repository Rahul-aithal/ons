# Todo CLI

`todo` is a small Go command-line application for storing tasks in a CSV file. It can add tasks, list the stored tasks, change task status, and delete tasks.

## Requirements

- Go 1.27.0 or later
- Write access to the data file location

## Build

```sh
go build -o todo .
```

The repository also contains a previously built `todo` executable.

## Data storage

The application reads and writes one CSV file:

- If `DB_FILE` is set, that path is used.
- Otherwise, the default path is `$HOME/data.csv`.

The file has no header row. Each record uses this field order:

| Field | Example | Notes |
| --- | --- | --- |
| ID | `96c00a4f-6157-4e2d-8af3-921745f5caa4` | Generated UUID |
| Name | `Write documentation` | Supplied by the user |
| Status | `Pending`, `Late`, or `Completed` | Set by the application |
| Due | `31-12-2026` | Required `DD-MM-YYYY` date |
| Description | | Always written as empty by `add` |
| Priority | `low` | Always written as `low` by `add` |

## Commands

### Add a todo

```sh
./todo add "Write documentation" 31-12-2026
```

The command creates a UUID, validates the due date as `DD-MM-YYYY`, and stores the task. A future due date is stored as `Pending`; a due date that is not in the future is stored as `Late`.

### List todos

```sh
./todo list
```

`get` is an alias for `list`. The command prints an ID, name, status, and due-date table to standard output. If the data file does not exist, `list` creates it.

### Change a todo status

```sh
./todo completed ID
```

Current implementation detail: this command rewrites rows whose ID differs from the supplied ID as `Completed`; it does not rewrite the row whose ID matches. This repository currently documents that behavior without changing it.

### Delete a todo

```sh
./todo delete ID
```

The command removes the row whose ID matches the supplied ID.

## Logging and output

- `list` prints the task table to standard output.
- Successful `add`, `completed`, and `delete` operations print timestamped log messages to standard error.
- Invalid input and file failures print timestamped fatal log messages to standard error.

Example:

```text
2026/10/02 17:21:42 Added new todo Write documentation
```
