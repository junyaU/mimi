# Mimi: Go Dependency Quantifier CLI Tool

Mimi is a command-line interface (CLI) tool written in Go that quantifies the dependencies of Go packages. It helps you manage the complexity of your Go projects by providing detailed information about both direct and indirect dependencies.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
    - [Check Command](#check-command)
    - [Table Command](#table-command)
    - [List Command](#list-command)
    - [Deps Command](#deps-command)
    - [Run Command](#run-command)
- [Configuration File](#configuration-file)
- [Contributing](#contributing)
- [License](#license)

## Features

- Quantify direct and indirect dependencies of a Go package.
- Visualize dependencies in a table.
- Set thresholds for dependencies and get alerts when these thresholds are exceeded.


## Installation
Assuming you have a working Go environment (version 1.1.9 or newer), Mimi can be installed by running:

```sh
$ go get -u github.com/junyaU/mimi

$ go install github.com/junyaU/mimi
```

Make sure that your PATH includes the $GOPATH/bin directory so your commands can be easily used:

```sh
$ export PATH=$PATH:$GOPATH/bin
```

## Usage
Mimi provides several commands to analyze your Go packages.

### Check Command
Checks if the direct and indirect dependencies of a given Go package exceed the given thresholds.

The direct_threshold parameter specifies the maximum number of direct dependencies allowed. Direct dependencies are the packages that the given package directly depends on.

The indirect_threshold parameter specifies the maximum number of indirect dependencies allowed. Indirect dependencies are the packages that the given package depends on through one or more intermediary packages.

The depth parameter specifies the maximum depth of dependencies allowed. Depth is a measure of the farthest distance from the given package to a dependency, with a direct dependency being at a depth of 1, a dependency of a direct dependency being at a depth of 2, and so on.

The lines parameter specifies the maximum number of lines of code allowed in a single Go package. It's a way to keep your Go packages concise and maintainable. If a package's code exceeds this limit, it may indicate that the package is doing too much and may need to be broken down into smaller, more focused packages.


```sh
$ mimi check <package_path> --direct=<direct_threshold> --indirect=<indirect_threshold> --depth=<depth> --lines=<lines>
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

If any of the specified thresholds are exceeded, the check command will return an error.

### Table Command
Generates a table showing the direct and indirect dependencies for a given Go package.

```sh
$ mimi table <package_path> --direct=<direct_threshold> --indirect=<indirect_threshold> --depth=<depth> --lines=<lines> -w
```

ex) Generate a table showing the dependencies of the `github.com/junyaU/mimi/testdata/layer` package.

```sh  
$ mimi table ./testdata/layer/ -w
```

<img width="943" alt="スクリーンショット 2023-06-03 11 50 19" src="https://github.com/junyaU/mimi/assets/61627945/d5eab75c-a883-4e4e-9d07-0c3eb9b5d6da">

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

### Deps Command
Displays the dependents of a given Go package. A dependent is a package that relies on the specified package. This command can be particularly useful for identifying potential issues or impacts before making changes to the package.

```sh
$ mimi deps <package_path>
```

ex) Display the dependents of the `github.com/junyaU/mimi/testdata/layer/domain/model` package.

```sh
$ mimi deps ./testdata/layer/domain/model
github.com/junyaU/mimi/testdata/layer/domain/model/creator
  github.com/junyaU/mimi/testdata/layer/domain/model/recipe

github.com/junyaU/mimi/testdata/layer/domain/model/recipe
  github.com/junyaU/mimi/testdata/layer/domain/model/flow
  github.com/junyaU/mimi/testdata/layer/domain/model/necessity

github.com/junyaU/mimi/testdata/layer/domain/model/flow
  No dependents

github.com/junyaU/mimi/testdata/layer/domain/model/necessity
  No dependents
```

The output above indicates that the `creator` package is used by the `recipe` package, and the `recipe` package is used by both the `flow` and `necessity` packages. Neither the `flow` nor the `necessity` packages are used by any other packages.

### Run Command
You can use the run command by specifying the path to the configuration file as follows:

```sh
$ mimi run <config_file_path>
```

This will read the configuration from the specified YAML file and execute the commands defined in it.

#### Output

Upon successful execution, the run command will display a message indicating that the command was completed successfully along with the number of commands processed. If an error occurs during the execution of any command, it will stop the process and display an error message.

Please note that the actual output will depend on the commands specified in the configuration file. For instance, the list command will print a list of dependencies, whereas the table command will generate a table displaying the dependencies.


## Configuration File
The configuration file is a YAML file that contains a list of commands to be executed. Each command has a name and parameters associated with it. Here is an example of how the configuration file looks like:

```yaml
version: 1.0
commands:
  - name: check
    parameters:
      path: "./"
      directThreshold: 10
      indirectThreshold: 20
      depthThreshold: 6
      linesThreshold: 1000
  - name: list
    parameters:
      path: "./"
  - name: table
    parameters:
      path: "./"
      directThreshold: 10
      indirectThreshold: 20
```

In the above example, three commands check, list, and table will be executed.

## Contributing
Contributions to Mimi are welcome! Feel free to open an issue or submit a pull request if you have a way to improve this tool.

## License
Mimi is released under the MIT License. See the [LICENSE](https://github.com/junyaU/mimi/blob/master/LICENSE) file for more details.




