# firead

Create Go microservices that read Firestore documents.

Very easy to use, create a working service by adding just one line of code to your `main` method.

Either read documents with a `string` ID:
```
firead.RunStringId()
```
Or read documents with an `int` ID:
```
firead.RunIntId()
```

## Examples

Try working examples of firead [here](https://github.com/JonnyOrman/firead-examples)

## Getting started

Create a new Go project
```
mkdir firead-example
cd firead-example
go mod init firead/example
```

Get `firead`
```
go get github.com/jonnyorman/firead
```

Add a `main.go` file with the following
```
package main

import "github.com/jonnyorman/firead"

func main() {
	firead.RunStringId()
}
```

Add a `firead-config.json` file with the following
```
{
    "projectID": "your-firebase-project",
    "collectionName": "FirestoreCollection"
}
```

Tidy and run with access to a Firebase project or emulator
```
    go mod tidy
    go run .
```

Submit a `GET` to the service's `/{id}` path, replaceing `{id}` with the ID of the document you want to read. You will get a `200` response with the document in the response body.

You can also create a struct with the data you want to read from Firestore. Create a struct and use it:
```
package main

import "github.com/jonnyorman/firead"

type DocumentModel struct {
	Prop1 string
	Prop2 int
}

func main() {
	firead.RunTypedStringId[DocumentModel]()
}
```

## Environment configuration

The configuration can also be provided by the environment with the following keys:
- `projectID` - `PROJECT_ID`
- `collectionName` - `COLLECTION_NAME`

A combination of the `firead-config.json` file and environment variables can be used. For example, the project ID could be provided as the `PROJECT_ID` environment variable, while the collection name is provided with the following configuration file:
```
{
    "collectionName": "FirestoreCollection"
}
```

If a configuration value is provided in both `firead-config.json` and the environment, then the configuration file with take priority. For example, if the `PROJECT_ID` envronment variable has value "env-project-id" and the following `firead-config.json` file is provided:
```
{
    "projectID": "config-project-id",
    "collectionName": "FirestoreCollection"
}
```
then the project ID will be "config-project-id".