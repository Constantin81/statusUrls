# Urls check service

## Project description

This is a grpc microservice designed to check the url status.

## Creating and running a project in Docker

Create new docker image: 
``docker build -t status-url .``

Run app in container interactive mode:
``docker run -it -p 8080:8080 status-url``

Run container in background:
``docker run -d -p 8080:8080 status-url``

## Testing service
The Evans utility is used to test the service.
``https://github.com/ktr0731/evans``

Run utility in new terminal:
``evans api/status.proto -p 8080``

Function call:
``call (name function)``

## Endpoints
The service accepts url and the number of times it is checked per day, at least once a day (test url - ``http://pinterest.com``)

``AddUrl``

``strUrl - http://pinterest.com``

``countPointCheckUrl - 200``

Getting information on url. The endpoint accepts the url and gives information about the latest checks (time and status codes) (test url - ``http://google.com``)

``GetStatusUrl``

``strUrl - http://google.com``

Response:
``` 
{
  "containers": [
    {
      "timeCheckUrl": "2021-01-31 00:00:00 +0000 UTC",
      "statusUrl": 200
    }
  ]
}
```

Endpoint accepts url. The status of this url is not checked but the history remains available (test url - ``http://google.com``)

``DeleteUrl``

``strUrl - http://google.com``
