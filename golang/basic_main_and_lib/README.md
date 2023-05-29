
# golang skeleton workspace

## Important points

- The names of files is totally irrelevant anywhere in the code.

- The name of the directory that a module lives in is irrelevant anywhere in the
  code.  However, it would be common sense to have it match the name of the
  package within the root folder of the module.

- The name of the directory that a package lives in does matter.  It is used in
  the `import` path, after the module path.

- The name of the package within the root folder of the module is the default
  name used in dependent code.  The `import` statement can override it.

- Putting main in a separate module isn't necessary, but it does allow code that
  wants to use the project as a library to do so without having to import the
  parts of the code that support the binary running (command line flag parsing,
  for example).
  - An alternative would be to have a single go.mod file in the root of the
    code repository.  The directory structure could be kept the same as it is
    now with main in a separate folder.

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
  "example.com/foo/bar"
)
...

```
