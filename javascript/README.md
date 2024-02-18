# JavaScript

This is a simple API in Node.js using express with an in-memory data store and routes to handle Create, Read, Update, and Delete operations of Items.

_Note: All commands listed below are executed from this directory._

## Pre-requisites

- Node >= 18.17.1

## How to Run

Execute the following commands to run the application:

```bash
# install dependencies
npm install

# run the application in development mode
npm run dev
```

The application will now be accessible at `http://localhost:8080/ping`. See the "How to Test" section below for more information.

## How to Test

Navigate to `http://localhost:8080/ping` in your browser to view the application and confirm that it is running.

You can test the application routes using `curl`:

1. Create an item
```bash
$ curl -X POST http://localhost:8080/items \
    -H "Content-Type: application/json" -d '{"name": "item1"}' 

{"id":1,"name":"item1"}
```

2. Get all items
```sh
$ curl http://localhost:8080/items

[{"id":1,"name":"item1"}]
```

3. Get an item by ID
```sh
$ curl http://localhost:8080/items/1

{"id":1,"name":"item1"}
```

4. Update an item
```sh
$ curl -X PUT http://localhost:8080/items/1 \
    -H "Content-Type: application/json" -d '{"name": "item2"}'

{"id":1,"name":"item2"}
```

5. Delete an item
```sh
$ curl -X DELETE http://localhost:8080/items/1

Item deleted
```