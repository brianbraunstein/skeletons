
# golang skeleton multi-module workspace

## Caveat

Due to bug https://github.com/golang/go/issues/50750 the utility of workspaces
seems limited or fundamentally broken.  Until that's resolved, the code
structure documented here will not work as intended.  You're forced into either
one of the following problems:
1) A combined module for all code which forces a giant list of dependencies in
   the go.mod file, forcing clients to pull down unnecessary dependencies.
2) For a module A which depends on module B both within the same workspace W, if
   there are changes to B which A needs, and A also needs to `go mod tidy`, then:
  - Publish the changes to B in a separate commit.
  - Then run `git mod tidy` in B.

One will note that (1) doesn't use workspaces at all, and (2) is the same
procedure that would be needed as if A and B were not together within W.  This
means workspaces are mostly or entirely useless.

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
