# Simple Task Manager

This simple task manager allows users to create, read, update, and delete tasks. The application is built in Golang and uses MySQL for data storage.

## System Requirements:

- Golang installed
- MySQL installed

## Initial Setup:

1. Open `sql_files/schema.sql` and run the commands in the MySQL CLI to create the required database and tables.
2. Update the username and password details of MySQL in `internal\db\database.go` at line:17.

## Running the Application:

1. Navigate to the `cmd` directory by running the following command: `cd cmd`.
2. Run the command: `go run main.go` to start the application.

## Example usage:

### Create Task:

```
curl --location 'http://localhost:8080/tasks/create' --header 'Content-Type: application/json' --data '{"Title": "test title", "Description": "test description"}'
```

### Update Task:

```
curl --location 'http://localhost:8080/tasks/update' --header 'Content-Type: application/json' --data '{"ID": 0, "Title": "test title", "Description": "updated description"}'
```
### Read Specific Task:

```
curl --location 'http://localhost:8080/tasks/read?id=0'
```

### Delete Specific Task:

```
curl --location --request DELETE 'http://localhost:8080/tasks/delete?id=0' --header 'Content-Type: application/json'
```


Feel free to modify the data in the cURL commands according to your requirements.


### Testing:

To run the tests, run the following command:
```
go test ./tests
```
