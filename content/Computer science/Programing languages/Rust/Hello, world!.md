# Hello, world!

Rust files always end with the `*.rs*` extension. If you’re using more than one word in your filename, the convention is to use an underscore to separate them.

Filename: main.rs

```rust
fn main() {
    println!("Hello, world!");
}
```

In this code  `println!` is a macro ,macros don’t always follow the same rules as functions.

enter the following commands to compile and run the file:

```bash
rustc main.rs
./main
```