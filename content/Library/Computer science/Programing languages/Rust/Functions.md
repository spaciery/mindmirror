# Functions

```rust
fn main() {
    println!("Hello, world!");

    another_function();
}

fn another_function() {
    println!("Another function.");
}
```

function is declared using `fn` keyword.

example of passing arguments 

```rust
fn main() {
    print_labeled_measurement(5, 'h');
}

fn print_labeled_measurement(value: i32, unit_label: char) {
    println!("The measurement is: {value}{unit_label}");
}
```

## Statement and Expressions

- **Statements** are instructions that perform some action and do not return
a value.
- **Expressions** evaluate to a resultant value.

```rust
// this is a statement
// 6 is an expression that evaluates to 6
let y = 6;

// function defenition is also a statement, infact the entire function is a statement
fn main() {
    let y = 6;
}
```

`5 + 6`, is an expression that evaluates to the value `11`. Expressions can be part of statements ex : `let y = 5+6;`

Calling a function is an expression. 

Calling a macro is an expression. 

A new scope block created with curly brackets is an expression

```rust
fn main() {

// the code inside {} is an expression
// Expressions do not include ending semicolons.
// If you add a semicolon to the end of an expression, you turn it into a statement, and it will then not return a value.
    let y = {
        let x = 3;
        x + 1
    };
// println! is an expression
    println!("The value of y is: {y}");
}
```

## Functions with return values

```rust

// -> is used to specify the return value type
// here the last expression is automatically returned , but we can also use return keyword
fn five() -> i32 {
    5
}

fn main() {
    let x = five();

    println!("The value of x is: {x}");
}
```