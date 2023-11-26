# InstaGo (WIP)

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/tempor1s/notiongo)
[![Test Status](https://github.com/google/go-github/workflows/tests/badge.svg)](https://github.com/google/go-github/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/google/go-github/branch/master/graph/badge.svg)](https://codecov.io/gh/google/go-github)


InstaGo is a minimalist Instagram API wrapper.


## Installation ##

InstaGo is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/imthaghost/instago/
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/imthaghost/instago/instagram"
```

## Usage ##

```go
import "github.com/imthaghost/instago/v1/instagram"	 // with go modules enabled (GO111MODULE=on or outside GOPATH)
import "github.com/imthaghost/instago/instagram"   // with go modules disabled
```

Construct a new Instagram client, then use the various services on the client to
access different parts of the Instagram API. For example:

```go
client := instagram.NewClient(nil)

```

### Integration Tests ###

You can run integration tests from the `test` directory. See the integration tests [README](test/README.md).

## Contributing ##
I would like to cover the entire GitHub API and contributions are of course always welcome. The
calling pattern is pretty well established, so adding new methods is relatively
straightforward. See [`CONTRIBUTING.md`](CONTRIBUTING.md) for details.


## License ##

Library distributed under the Apache-2.0 license found in the [LICENSE](./LICENSE)
file.
