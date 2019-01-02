# dependency
[![Build Status](https://travis-ci.org/joaosoft/dependency.svg?branch=master)](https://travis-ci.org/joaosoft/dependency) | [![codecov](https://codecov.io/gh/joaosoft/dependency/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/dependency) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/dependency)](https://goreportcard.com/report/github.com/joaosoft/dependency) | [![GoDoc](https://godoc.org/github.com/joaosoft/dependency?status.svg)](https://godoc.org/github.com/joaosoft/dependency)

A simple dependency manager with a internal vcs.
This dependency manager is being used in all my personal projects :)

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* Get, to get the dependencies
* Reset, to delete the user locked dependencies and Get dependencies
* Add <dependency>, to add a new dependency
* Remove <dependency>, to remove a existing dependency

>### Go
```
go get github.com/joaosoft/dependency
```

## Usage 
> Commands
```
// generate dependencies
dependency get (takes in consideration the import-gen.yml (is exists) and the import-lock.yml)

// update dependencies (just takes in consideration the import-lock.yml)
dependency update

// delete lock configuration
dependency reset

// add a new dependency (just takes in consideration the import-lock.yml)
dependency add github/joaosoft/web

// remove a dependency (just takes in consideration the import-lock.yml)
dependency remove github/joaosoft/web
```

> Files
* import-gen.yml, generated files with dependencies
* import-lock.yml, user dependencies lock

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
