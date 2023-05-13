# Mimi: Go Dependency Quantifier CLI Tool

Mimi is a command-line interface (CLI) tool written in Go that quantifies the dependencies of Go packages. It helps you manage the complexity of your Go projects by providing detailed information about both direct and indirect dependencies.

## Features

- Quantify direct and indirect dependencies of a Go package.
- Visualize dependencies in a table.
- Set thresholds for dependencies and get alerts when these thresholds are exceeded.


## Installation

Assuming you have a working Go environment, Mimi can be installed by running:

```sh
$ go install github.com/junyaU/mimi
```

## Usage
Mimi provides several commands to analyze your Go packages.

### Check Command
Checks if the direct and indirect dependencies of a given Go package exceed the given thresholds.

```sh
$ mimi check <package_path> --direct=<direct_threshold> --indirect=<indirect_threshold>
```

ex) Check if the direct dependencies of the `github.com/junyaU/mimi/testdata` package exceed 2.

```sh  
$ mimi ./testdata --direct=2
Package github.com/junyaU/mimi/testdata/layer/adapter/data_handler has 3 direct dependencies
Package github.com/junyaU/mimi/testdata/layer/usecase/recipes has 3 direct dependencies
Package github.com/junyaU/mimi/testdata/layer/adapter/presenters has 3 direct dependencies
Package github.com/junyaU/mimi/testdata/layer/infra has 5 direct dependencies
Error: exceeded dependency threshold
exit status 1
```

### Table Command
Generates a table showing the direct and indirect dependencies for a given Go package.

```sh
$ mimi table <package_path> --direct=<direct_threshold> --indirect=<indirect_threshold>
```

ex) Generate a table showing the direct and indirect dependencies of the `github.com/junyaU/mimi/testdata/layer/domain/model` package.

```sh  
$ mimi table ./testdata/layer/domain/model
+--------------------------------------------------------------+-------------+---------------+
|                           PACKAGE                            | DIRECT DEPS | INDIRECT DEPS |
+--------------------------------------------------------------+-------------+---------------+
| github.com/junyaU/mimi/testdata/layer/domain/model/creator   | 0           | 0             |
+--------------------------------------------------------------+-------------+---------------+
| github.com/junyaU/mimi/testdata/layer/domain/model/recipe    | 1           | 0             |
+--------------------------------------------------------------+-------------+---------------+
| github.com/junyaU/mimi/testdata/layer/domain/model/flow      | 2           | 1             |
+--------------------------------------------------------------+-------------+---------------+
| github.com/junyaU/mimi/testdata/layer/domain/model/necessity | 2           | 1             |
+--------------------------------------------------------------+-------------+---------------+
```

### List Command
Lists all the dependencies of a given Go package.

```sh
$ mimi list <package_path>
```

ex) List all the dependencies of the `github.com/junyaU/mimi/testdata/layer/domain/model` package.

```sh
$ mimi list ./testdata/layer/domain/model
github.com/junyaU/mimi/testdata/layer/domain/model/creator
  Direct Deps:
    No direct dependency
  Indirect Deps:
    No indirect dependency

github.com/junyaU/mimi/testdata/layer/domain/model/recipe
  Direct Deps:
    github.com/junyaU/mimi/testdata/layer/domain/model/creator
  Indirect Deps:
    No indirect dependency

github.com/junyaU/mimi/testdata/layer/domain/model/flow
  Direct Deps:
    github.com/junyaU/mimi/testdata/layer/domain/model/recipe
    github.com/junyaU/mimi/testdata/layer/domain
  Indirect Deps:
    github.com/junyaU/mimi/testdata/layer/domain/model/creator

github.com/junyaU/mimi/testdata/layer/domain/model/necessity
  Direct Deps:
    github.com/junyaU/mimi/testdata/layer/domain
    github.com/junyaU/mimi/testdata/layer/domain/model/recipe
  Indirect Deps:
    github.com/junyaU/mimi/testdata/layer/domain/model/creator
```

## Contributing
Contributions to Mimi are welcome! Feel free to open an issue or submit a pull request if you have a way to improve this tool.

## License
MIT License




