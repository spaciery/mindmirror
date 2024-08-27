# Development env setup

`rustup` is a command line tool for managing Rust versions and associated tools. we can install `rustup` using 

```bash
curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
```

After installing `rustup` restart shell and verify using 

```bash
rustc --version
```

To update/uninstall rust

```bash
rustup update 
rustup self uninstall
```

Install gcc or clang or build-essential if faced by linker error

To format rust code  use `rustfmt` . 

Rust style is to indent with four spaces, not a tab.