## Projector Cli
A simple cli app that stores key value based on the path 

### Setup
```
git clone git@github.com:ombudhiraja/projector-cli.git
cd projector

go build -o projector main.go
```

### Usage
```
./projector
- List all the keys in the current path and its parents.

./projector foo
- Get the value of foo in the current path and its parents.

./projector add foo bar
- Add foo bar to the given path or current path

./projector rm foo
- Remove foo from the given path or current path
```
