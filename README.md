## Intro
This repository exposes an API endpoint to send SMS's via MessageBird.

## Before running

* You need to have an API key from message bird.

* Setting environment variables -

  Create a file called .env and see "example_env" for what variables need to be set.

## How to run (hassle free with docker)

Navigate to the root directory and 

```
make docker-run
```

Note - if you have not created a file called .env and put it in the root directory with the right values, the service will not run.

## How to build/run locally (if you have Go installed)

You need to set the environment variables in the .env file and then execute

```
. .env
```

or 

```
source .env
```

in order to source the env file

Navigate to the root directory and 

```
make run
```

In order to do this you need Go installed. In case you don't have go or don't want to install it, use make docker-run as mentioned above.

## Run tests

```
make test
```

## Sample SMS request -

```
POST http://localhost:8081/api/sms
```

*Body*

```
{
 "recipient" : "+4917636504146",
 "originator" : "Balaji",
 "message" : "He had come to that moment in his age when there occurred to him, with increasing intensity, a question of such overwhelming simplicity that he had no means to face it. He found himself wondering if his life were worth the living; if it had ever been. It was a question, he suspected, that came to all men at one time or another;"
}
```

## Code

/cmd/api/main.go is the entrypoint. Start exploring from there.
