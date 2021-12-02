# advent-of-code-2021

Solutions for https://adventofcode.com/2021, largely an excuse to learn [golang](https://go.dev/)

Uses a vscode devcontainer for dependencies.

## golang learning notes

### Day 1

- you can pass go channels around as send/recv only types chan<- / <-chan
- vscode go plugin is pretty good, lots of "do what I want" features already
  with no extra config
- defer is pretty sweet
- channels can be used like C# IEnumerable or python yield, but you need to add
  concurrency to do so and be careful about when you close the channel
- a little surprised to have zero problems with closures or pointers
- the stdlib seems a little bare bones, hoping I just haven't found what to
  import yet

### Day 2

- ~~every file in a folder is in the same package~~ depends on the `package` declaration at the top of each file
- PascalCase means "public"
- docstrings go above the element, free form comment. Maybe there's a jsdoc equivalent?
- organization hierarchy:
    1. git repo -> module
    2. subdirectories -> packages in a module
    3. files -> functions/types in a package
- tests can live next to the code being tested, but add a `_test` suffix on the
  package to prevent exporting the tests with the rest of the package
- [`Example*` tests](https://pkg.go.dev/testing#hdr-Examples) exist, and seem to
  satisify my want for something like `Assert.AreEqual`
- you can import symbols from one package into yours with `import . "other"`,
  good for testing but seems iffy elsewhere; basically a pythong `from x import *`
- wow, error propagation is really a thing
- vscode linter is giving some good stuff, e.g. `errors.New(fmt.Sprintf("bad: %v", x))` to `fmt.Errorf("bad: %v", x)`