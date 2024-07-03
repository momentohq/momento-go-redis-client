<img src="https://docs.momentohq.com/img/momento-logo-forest.svg" alt="logo" width="400"/>

[![project status](https://momentohq.github.io/standards-and-practices/badges/project-status-official.svg)](https://github.com/momentohq/standards-and-practices/blob/main/docs/momento-on-github.md)
[![project stability](https://momentohq.github.io/standards-and-practices/badges/project-stability-beta.svg)](https://github.com/momentohq/standards-and-practices/blob/main/docs/momento-on-github.md)


# Welcome to the Momento Go-Redis contributing guide :wave:

Thank you for taking your time to contribute to our Momento `go-redis` wrapper!
<br/>
This guide will provide you information to start your own development and testing.
<br/>
Happy coding :dancer:
<br/>

## Submitting

If you've found a bug or have a suggestion, please [open an issue in our project](https://github.com/momentohq/terraform-provider-momento/issues).

If you want to submit a change, please [submit a pull request to our project](https://github.com/momentohq/terraform-provider-momento/pulls). Use the normal [Github pull request process](https://docs.github.com/en/pull-requests). 

## Requirements :coffee:

- [Go 1.18 or above](https://go.dev/doc/install) is required 
- A Momento API key is required, you can generate one using the [Momento Console](https://console.gomomento.com)

## Developing :computer:

### Build 

```bash
make build
```

### Formatting and Tidy :flashlight:

```bash
make lint
```

## Tests :zap:

### Run integration tests against Momento 

```bash
export TEST_AUTH_TOKEN=<YOUR_AUTH_TOKEN>
make test-momento
```

### Run integration tests against Redis

First run Redis either natively, run Redis in a Docker container, or do your development in a devcontainer. Here is an example of running Redis in a Docker container:

```bash
docker run -it -p 6379:6379 redis
```

Then run the tests

```bash
make test-redis
```

This assumes the Redis server is running on `localhost:6379`.

By running Redis on the local host, you can use the `redis-cli` to inspect the state of the Redis server as well as interactively debug the tests. 

### Run all tests

This will run both the integration tests against Momento and Redis. As above, we assume the Redis server is running on `localhost:6379`.

```bash
export TEST_AUTH_TOKEN=<YOUR_AUTH_TOKEN> 
make test
```

----------------------------------------------------------------------------------------
For more info, visit our website at [https://gomomento.com](https://gomomento.com)!
