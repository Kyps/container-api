# KSI Container API

KSI Container API is an API that uses Guardtime KSI Blockchain to create containers from user files and providing it with data hierarchy and signatures.
The APi allows for creation of containers, adding signatures to containers and removing signatures from created containers.

At the moment there is no database implementation because task file did not specify it. Basic functionality that the task required could be done without it.

## Installation

For building and installing use the standard go tools.

To get the server up and running, first you need to edit the configuration file `conf.json`, provide the server with a temporary files directory (default value is `./tmp/`) and database files directory ( default value is `./database/`). If the directories do not exist, they will be created at initial run.

To get correct signature files for manifest files you also need to provide Guardtime Signing service URL and credentials. These can be aquired from [Guardtime developers page](https://guardtime.com/technology/blockchain-developers).

## Usage

By default the server is running on port 4000. This can be changed in the configuration file.

To create new container use `/` endpoint with `POST` method. Add multipart form to the body of the request with `data` key that includes the files you want to upload.

For adding new signatures to a container use `/sign` endpoint with `PUT` method. Add query parameter name "key" and the UUID of the container you want to add a signature to. You do not need to add file extension to the value.

Deleting signatures from a container uses `/sign` endpoint with `DELETE` method. Add query parameter name "key" and the UUID of the container you want to delete a signature from. You do not need to add file extension to the value.

For testing purposes there has been exposed an endpoint `/getall` with a `GET` method. The request returns a JSON object containing an array of database files.

## Tests

To run tests just run `go test` to test main project directory, `go test ./...` to test whole project. API endpoints also can be tested with curl or Postman.

## Dependencies

The project depends on [guardtime/goksi](https://github.com/guardtime/goksi).
No other 3rd party dependencies used.

## Compatibility

Go 1.10 or newer.

## Todo

-   Code refactor
-   Add actual database
-   More tests
