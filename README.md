# Todo List Example

This is simple RESTful API of todo list app built using Nano HTTP Multiplexer & SQLite

## Contents

- [Installation](#installation)
- [API Endpoint](#api-endpoint)
  - [Todo List](#todo-list)
  - [Create Todo](#create-todo)
  - [Todo Detail](#todo-detail)
  - [Toggle Todo Status](#toggle-todo-status)
  - [Remove Todo](#remove-todo)

## Installation

To install this example, you need to install Go, SQLite, and set your Go workspace first.
you need [Go](https://golang.org/) installed (**version 1.11+ is required**), then you can use the below Git command to copy this project into your computer.

```bash
git clone https://github.com/hariadivicky/nano-example
```

Now we assume that you are on `nano-example/` directory.

Then you can create new SQLite database on `database/`

```bash
touch database/todos.db
```

Feel free to choose your database name. You could find this configuration on `main.go` file

```go
models.OpenDatabase("file:./database/todo.db")
```

Now you need to migrate `todos` table schema

```bash
$  sqlite3 database/todo.db
>  .read database/create_todos_table.sql
>  .exit
```

You are ready now, build and run it

```bash
go build && ./todo
```

server runs on `http://localhost:8080`

## API Endpoint

### Todo List

show the collection of todo.

`GET` [http://localhost:8080/todos](http://localhost:8080/todos)

response:

```json
{
    "collection": [
        {
            "id": 1,
            "title": "learn golang",
            "is_done": false
        },
        {
            "id": 2,
            "title": "practice golang",
            "is_done": false
        }
    ]
}
```

### Create Todo

create new todo

`POST` [http://localhost:8080/todos](http://localhost:8080/todos)

Headers:

`Content-Type: application/x-www-form-urlencoded`

`Data: title=use%20nano&is_done=true`

response:

```json
{
    "todo": {
        "id": 3,
        "title": "use nano",
        "is_done": true
    }
}
```

### Todo Detail

display todo detail

`GET` [http://localhost:8080/todos/2](http://localhost:8080/todos/2)

response:

```json
{
    "todo": {
        "id": 2,
        "title": "practice golang",
        "is_done": false
    }
}
```

### Toggle Todo Status

toggle todo done status

`PUT` [http://localhost:8080/todos/2/toggle](http://localhost:8080/todos/2/toggle)

response:

```json
{
    "updated": true,
    "todo": {
        "id": 2,
        "title": "practice golang",
        "is_done": true
    }
}
```

### Remove Todo

delete todo record

`DELETE` [http://localhost:8080/todos/2](http://localhost:8080/todos/2)

response:

```json
{
    "deleted": true
}
```
