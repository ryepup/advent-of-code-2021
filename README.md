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
- PascalCase means "public" _this is weirdly nominal_; a symbol needs to be
  PascalCase to use that _name_ in another file, but you can use values just
  fine if you don't need the type name. i.e. you can use a PascalCase (public)
  function to return a camelCase (private) type, and the calling code and use
  return value just fine, but it can't use the _name_ of the private type.
  `ParseCommands` returns a `<-chan command`. `ParseCommands` can be used just
  fine, it's return value can be used just fine with `var`, but you can't
  declare a variable of type `command`
- docstrings go above the element, free form comment. Maybe there's a jsdoc equivalent for metadata?
- organization hierarchy:
    1. git repo -> module
    2. subdirectories -> packages in a module, usually one package per dir?
    3. files -> functions/types in a package
- tests can live next to the code being tested, but add a `_test` suffix on the
  package to prevent exporting the tests with the rest of the package
- [`Example*` tests](https://pkg.go.dev/testing#hdr-Examples) exist, and seem to
  satisify my want for something like `Assert.AreEqual`
- you can import symbols from one package into yours with `import . "other"`,
  good for testing but seems iffy elsewhere; basically a pythong `from x import *`
- wow, error propagation is really a thing
- vscode linter is giving some good stuff, e.g. `errors.New(fmt.Sprintf("bad: %v", x))` to `fmt.Errorf("bad: %v", x)`
- enums exist, kinda odd to hang out as bare global symbols

### Day 3

- external libs are easy, `go.sum` is a lock file and it's nice that `go.mod`
  shows all the indirect dependencies. Would be very obvious if you fall into
  npm-style thousands of deps
- there's no direct
  [devDependencies](https://docs.npmjs.com/specifying-dependencies-and-devdependencies-in-a-package-json-file)
  equivalent, but there's [module graph
  pruning](https://go.dev/ref/mod#graph-pruning) so consumers of a library will
  only download modules that their code depends on - this might be a good reason
  to put tests in a different package
- type definitions are nice, there's something really subtle there with type
  definitions (`type meters int`) vs type aliases (`type meters = int`). aliases
  don't create a new type (i.e. [Rust
  `newtype`](https://doc.rust-lang.org/rust-by-example/generics/new_types.html));
  the compiler will be happy assigning any `int` to a `meters`. Seems like a
  nice way to get some compiler guarantees on correctness.
- `make` can handle user-defined array types
- I think golang wants me to stop thinking too much

### Day 4

- mutable data wants to be a pointer; there is more default immutability via
  implicit copies than I realized. This can get a little weird; for a while I
  was copying an array of pointers, and mutating the structs being pointed at. I
  can reduce memory by passing around a pointer to my array
- vscode plugin has some profiling features via `pprof`, but the tests run fast
  enough that I don't get much data. The profiler runs as an http server in my
  go process, and it's allocation/cpu dwarfs what my stupid code does
- `pprof` uses graphviz, which is nice to see in the wild
- still feeling the lack of map/filter/reduce, but the use cases are all just
  different enough that I didn't write a `map` helper. In the long run that
  probably saves a lot of computation; allocating lambdas, etc
- OO style still feels kinda wrong in golang. I'm used to starting with an
  interface to work out the shape of the solution, and then jump into
  implementation, ending with a lot of interfaces with only one implementation. in golang it seems like dropping the interface is a better approach.
- labels on for loops are a nice way to reduce accounting, but seems like it
  could get iffy fast

### Day 5

- raw strings exist! perfect for regex
- struct equality is nice, mostly DWIM for basic things
- default zero values are a little magical; can write less code, but why it
  works is a little less clear without a "everything starts empty" mindset
- symbol visibility for tests is still a thing, a few options:

    1. make it public for fine-grained testing
    1. put the tests in the package (potentially making test libs a runtime dep
      for library consumers)
    1. make it public in a special [`internal` package](https://go.dev/doc/go1.4#internalpackages)

- tests have naming conventions to match up the test with the code being tested
- no ternary operator, "clever" workarounds feels worse than if/else
- stdlib is a little bare bones; feels pretty silly writing `abs` functions

### Day 6

- went down a rabbithole of scaffolding to get some practice with IO and
  templating (inspired by @smsutherland, might port some of his work w/
  downloading inputs later)
- managing newlines in templates is a PITA
- data-driven tests are nice, but error reporting gets a little worse (thanks
  @unwashedmeme)
- `assert` doesn't actually stop the run, but `require` does; this might be nice
  if there are multiple things that could go wrong and want to see all the
  errors in one test run. (thanks @unwashedmeme)
- I'm used to trying to write declarative code to focus on the "what" more than
  the "how". Feels like golang wants me to stay thinking about the "how", and
  just give enough primitives that the "how" also feels simple/obvious. There
  were a few times I wanted something like `sum(map.Values())` and ended up with
  a `for` loop. Writing a `.Values` function can't be done generically (yet!),
  and isn't really worth writing over the obvious `for` loop. If it were there I'd use it, but adding it doesn't feel right