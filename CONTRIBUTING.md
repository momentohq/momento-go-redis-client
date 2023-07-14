# Welcome to the Momento Go-Redis contributing guide :wave:

Thank you for taking your time to contribute to our Momento `go-redis` wrapper!
<br/>
This guide will provide you information to start your own development and testing.
<br/>
Happy coding :dancer:
<br/>

## Requirements :coffee:

- Go 1.2 or above is required https://go.dev/doc/install
- A Momento auth token is required, you can generate one using the [Momento Console](https://console.gomomento.com)

<br/>

## Build :computer:

```
make build
```

<br/>

## Formatting and Tidy :flashlight:

```
make lint
```

<br/>

## Tests :zap:

### Run integration tests against Momento 

```bash
TEST_AUTH_TOKEN=<YOUR_AUTH_TOKEN>
make test-momento
```
### Run integration tests against Redis

First run Redis either natively, run Redis in a Docker container, or do your development in a devcontainer. Here is an example of running Redis in a Docker container:

```
docker run -it -p 6379:6379 redis
```

Then run the tests against .NET6 and .NET Framework:

```
make test-redis
```

This assumes the Redis server is running on `localhost:6379`.

By running Redis on the local host, you can use the `redis-cli` to inspect the state of the Redis server as well as interactively debug the tests. We have also included a devcontainer config which will mount the entire coding environment into a Docker container, including the Redis server. See `.devcontainer` for more details.

### Run all tests

This will run both the integration tests against Momento and Redis. As above, we assume the Redis server is running on `localhost:6379`.

```
TEST_AUTH_TOKEN=<YOUR_AUTH_TOKEN> make test
```