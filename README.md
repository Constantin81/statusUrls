# Urls check service

## Project description

This is a grpc microservice designed to check the url status.

## Creating and running a project in Docker

Create new docker image
<mark>docker build -t status-url .</mark>

Run app in container interactive mode
<mark>docker run -it -p 8080:8080 status-url</mark>

Run container in background
<mark>docker run -d -p 8080:8080 status-url</mark>

## Testing service
The Evans utility is used to test the service.
<mark>https://github.com/ktr0731/evans</mark>

Run utility in new terminal
<mark>evans api/status.proto -p 8080</mark>

## Endpoins
The service accepts url and the number of times it is checked per day, at least once a day (test url - http://pinterest.com)
<mark>call AddUrl</mark>
strUrl (TYPE_STRING) => http://pinterest.com
countPointCheckUrl (TYPE_INT32) => 200

Getting information on url. The endpoint accepts the url and gives information about the latest checks (time and status codes) (test url - http://google.com)
<mark>call GetStatusUrl</mark>
strUrl (TYPE_STRING) => http://google.com

Endpoint accepts url. The status of this url is not checked but the history remains available (test url - http://google.com)
<mark>call DeleteUrl</mark>
strUrl (TYPE_STRING) => http://google.com
