# Task Tracker

Task Tracker is a simple command-line application for managing tasks. It allows you to add, update, list, and delete tasks, with support for tracking the progress status of each task.

## Features

- Add new tasks with descriptions.
- Update existing tasks.
- List all tasks or filter tasks by their status (Not Started, In Progress, Done).
- Delete tasks.

## Installation

To install the Task Tracker, clone the repository and build the application:

```sh
git clone https://github.com/BDrimus/task-tracker.git
cd task-tracker
go build
```

## Usage

### Add a Task

To add a new task, use the `add` command followed by the task description:

```sh
./task-tracker add "This is a new task"
```

### Update a Task

To update an existing task, use the `update` command followed by the task ID and the new description:

```sh
./task-tracker update [task-id] "This is an updated task"
```

### List Tasks

To list all tasks, use the `list` command:

```sh
./task-tracker list
```

You can also filter tasks by their status:

```sh
./task-tracker list todo
./task-tracker list inProgress
./task-tracker list done
```

### Delete a Task

To delete a task, use the `delete` command followed by the task ID:

```sh
./task-tracker delete [task-id]
```
