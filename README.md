# Worklog

Worklog is a lightweight CLI tool for tracking what you work on during the week. It lets you quickly add, edit, remove, and report tasks, helping you stay organized and reflect on your weekly progress.

## Features

- Add new tasks to your weekly log
- Edit or remove tasks from history
- Generate weekly reports to review your progress
- Search for tasks in your history

## Installation

Clone the repository and make sure the executable is available in your PATH.

```bash
git clone <repository-url>
cd worklog
```

## Usage

```bash
worklog [command]
```

### Available Commands

- **completion**: Generate the autocompletion script for the specified shell
- **drop**: Delete all tasks from history
<!-- - **find**: Search for tasks in your history -->
- **help**: Help about any command
- **new**: Add a new task to your weekly log
- **remove**: Delete a task from your history
- **report**: Show a summary of tasks for the current week

### Flags

- `-p, --db-path string` : Location of the database file (default `.config/worklog`)
- `-h, --help` : Help for worklog

Use `worklog [command] --help` for more information about a command.

## Example

```bash
worklog new "Finish writing documentation"
worklog report
```

## License

This project is licensed under the MIT License.
