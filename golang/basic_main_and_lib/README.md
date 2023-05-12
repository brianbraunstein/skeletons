
# golang skeleton workspace

## Important points

- The tinylib directory name and package name are irrelevant when using the lib
  in code.  Only the module name is used.

- Putting main in a separate directory cleanly and clearly separates the main
  modules go.mod file from the workspace.

## Adding a new library module

```
cd $workspace_root_dir
mkdir some_dir_name_irrelevant
(cd some_dir_name_irrelevant; go mod init example.com/foo/bar)
# automatically update the go.work file (Makefile can do this too).
go work use -r .

```

import it in other modules in this workspace via

```
import (
  bar "example.com/foo/bar"
)
```
