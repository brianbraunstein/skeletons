
# golang skeleton multi-module workspace

## Motivation

While more complex than
[`../single_main_and_lib_module`](../single_main_and_lib_module) The motiviation
for this example is the following:

- This demostrates how to develop multiple modules in the same code repository.

- Putting main in a separate modules allows code that wants to use the project
  as a library to do so without having to import the parts of the code that
  support the binary running (command line flag parsing libraries, for example).

## Adding a new library module to a workspace

```
cd $workspace_root_dir
mkdir some_dir_name
(cd some_dir_name; go mod init example.com/foo/bar)
# automatically update the go.work file (consider having the Makefile do this).
go work use -r .
```

import it in other modules in this workspace via
```
import (
  "example.com/foo/bar"
)
...
```
