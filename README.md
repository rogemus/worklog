# Worklog

`Worklog` is a lightweight command-line tool for tracking what you're working on. It records short task entries in a local text file so you can quickly add, edit, remove, and review work without any cloud or subscription.

Key points:
- Local-first: data is stored on your machine (no cloud).
- Simple format: human-editable text file.
- Fast workflow: add tasks from the terminal and generate weekly reports.

## Data format

Entries are single-line records starting with a timestamp, a title, and optional metadata. Example entries:
```
[2006-01-02 15:04:00] Finish documentation @tags:docs,writing @issue_id:XXX-123 @repo:super-repo @branch:master
[2006-01-02 15:10:00] Quick note about a task
[2006-01-02 15:12:00] Follow up @issue_id:XXX-123
```

By default the database is a text file under your config directory (default: `.config/worklog`) and can be edited directly with your editor.

## Features

- Create new task entries
- Edit or remove past entries
- Generate weekly summary reports
- Search and filter historical tasks

## Installation

Install via Go:

```bash
go install github.com/rogemus/worklog@latest
```
Ensure the worklog binary is in your PATH.


## Usage

```bash
worklog [command]
```

### Common commands:

- **completion** — generate shell autocompletion
- **drop** — delete all tasks
- **find** — search task history
- **new** — add a new task
- **remove** — delete a specific task
- **report** — show the current week summary
- **version** — show app version
- **help** — command help

### Flags

- `-p, --db-path string` : Location of the database file (default `.config/worklog`)
- `-h, --help` : Help for worklog

Use `worklog [command] --help` for details on a specific command.

## Example

```bash
worklog new "Finish writing documentation"
worklog report
worklog find "documentation"
```
