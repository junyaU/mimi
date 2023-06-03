# Mimi: Go Dependency Quantifier CLI Tool

Mimi is a command-line interface (CLI) tool written in Go. It provides quantitative information about the dependencies of Go packages such as the number of direct and indirect dependencies, the depth of these dependencies, and their weight. This detailed knowledge can help you understand and manage the complexity of your Go projects better.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Weight of Dependencies](#weight-of-dependencies)
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

- **Quantify Direct and Indirect Dependencies** : Mimi counts the number of direct and indirect dependencies of your Go packages. This information can be instrumental in understanding the structure and complexity of your projects. For example, a high number of indirect dependencies may indicate a complex package structure that could be difficult to maintain.

- **Visualize Dependencies in a Table** : Mimi can display dependencies in a table format, giving you a visual overview of your project's structure. By using the -w option in the table command, the table can be sorted based on the weight of the dependencies.

- **Set Thresholds for Dependencies** : With Mimi, you can set thresholds for dependencies and receive alerts when these thresholds are exceeded. This feature can help enforce good coding practices in large projects and ensure that packages don't become overly complex or dependent on too many other packages.

- **Calculate the Weight of Dependencies** : Mimi calculates a "weight" for each dependency, reflecting its significance within the project. This feature can help you identify key dependencies that could be a focus for optimization or refactoring.

- **Get Dependency Alerts** : With the -w option in the check command, you can receive alerts when the weight of a dependency exceeds a specified threshold. This can be a valuable tool for maintaining code quality over time.

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

## Weight of Dependencies
Weight is a measure of the dependency's significance and is calculated based on various factors like the number of dependent packages, how deeply nested the dependency is, etc. Weight provides a quantitative way to evaluate the impact of a dependency on your project.

The weight of a dependency ranges from 0 to 1, with 0 being the least significant and 1 being the most significant.

The calculation of weight is done as follows:

Weight is calculated using four factors: direct dependencies, indirect dependencies, dependents, and dependency depth. Each of these factors is normalized to a value between 0 and 1 and then weighted according to the following percentages:

- Direct dependencies: 30%
- Indirect dependencies: 30%
- Dependents: 20%
- Dependency depth: 20%

The weight is then calculated as follows:

```
Weight = (DirectScore * 0.3) + (IndirectScore * 0.3) + (DependentScore * 0.2) + (DepthScore * 0.2)
```

Each of the scores (DirectScore, IndirectScore, DependentScore, DepthScore) is calculated by normalizing the corresponding value (the number of direct dependencies, indirect dependencies, dependents, and dependency depth respectively) to a range between 0 and 1. The normalization is done based on the minimum and maximum limits specified for each value.

The normalization formula is:

```
Score = (Value - MinValue) / (MaxValue - MinValue)
```

## Usage
Mimi provides several commands to analyze your Go packages.

### Check Command
The check command evaluates the dependencies of a Go package against defined thresholds, checking for five key parameters: direct dependencies, indirect dependencies, dependency depth, lines of code in the package, and weight of dependencies.

- direct - Maximum permissible direct dependencies.
- indirect - Maximum permissible indirect dependencies.
- depth - Maximum permissible dependency depth.
- lines - Maximum permissible lines of code in a package.
- weight - Maximum permissible weight of dependencies.

```sh
$ mimi check <package_path> --direct=<direct> --indirect=<indirect> --depth=<depth> --lines=<lines> --weight=<weight>
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

The --weight option allows the table to be sorted based on the weight of dependencies. Dependencies with a weight of 0.3 or less are color-coded green, between 0.3 and 0.7 yellow, and between 0.7 and 1.0 red.

```sh
$ mimi table <package_path> --direct=<direct> --indirect=<indirect> --depth=<depth> --lines=<lines> --weight
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
      enableWeight: true
```

In the above example, three commands check, list, and table will be executed.

## Contributing
Contributions to Mimi are welcome! Feel free to open an issue or submit a pull request if you have a way to improve this tool.

## License
Mimi is released under the MIT License. See the [LICENSE](https://github.com/junyaU/mimi/blob/master/LICENSE) file for more details.




