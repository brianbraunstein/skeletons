
## Important points

- The names of files is totally irrelevant anywhere in the code.  Only a build
  system like `make` cares.

- A directory's files must all use the same package name.

- When importing code from a directory `X` containing files using package name
  `Y`, the default import name is `Y`.  `X` is irrelevant, other than in the
  import path.  This is true for both modules and subdirectories of modules.
  However, it would be common sense to make `X == Y`.

- You don't need a package at the root folder of a module.

- If creating a module to be shared on github, then the module name and path
  should match the github URL, like
  `github.com/brianbraunstein/some_repo/some_module`.  Again, the package name
  is competely unrelated, and there doesn't need to be a package at the root,
  but if there is, it would be common sense to make it named `some_module`.
