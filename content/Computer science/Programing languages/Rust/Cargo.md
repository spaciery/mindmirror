# Cargo

Cargo is Rust’s build system and package manager.

to check cargo version 

```bash
cargo --version
```

To create a project using cargo 

```bash
cargo new project_name
cd project_name
```

cargo will create `src` , `cargo.toml` , `.gitignore` and initialize a Git repository.Git files won’t be generated if you run `cargo new` within an existing Git repository.

Cargo expects your source files to live inside the *src* directory. The
top-level project directory is just for README files, license information,
configuration files, and anything else not related to your code.

`[package]` is a section heading that indicates that the following statements are configuring a package. 

`[dependencies]`, is the start of a section for you to list any
of your project’s dependencies.

To build the project

```bash
# normal build (debug)
cargo build
# when ready for release (compiles with optimization)
cargo build --release
```

the default build is a debug build, Cargo puts the binary in a directory named *debug* inside *target* folde*r*.

Running `cargo build` for the first time also causes Cargo to create a new file at the top level: *Cargo.lock*. This file keeps track of the exact versions of dependencies in your project.

instead if building and running the binary manually , we can use this to do it in one command.

```bash
cargo run
```

cargo will only rebuild if there is any change in the code.

to check the code , this is faster than cargo build , so to check for errors this is more suitable

```bash
cargo check
```