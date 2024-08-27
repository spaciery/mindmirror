# DataType

There are 2 types of datatypes in rust , scalar and compound. A *scalar* type represents a single value. Rust has four primary scalar types: integers, floating-point numbers, Booleans, and characters. *Compound types* can group multiple values into one type. Rust has two primitive compound types: tuples and arrays.

more details : [**Data Types**](https://doc.rust-lang.org/book/ch03-02-data-types.html#data-types)

### The Tuple Type

A *tuple* is a general way of grouping together a number of values with a variety of types into one compound type. Tuples have a fixed length: once declared, they cannot grow or shrink in size.

```rust
fn main() {
    let tup_variable: (i32, f64, u8) = (500, 6.4, 1);
}
```

```rust
fn main() {
    let tup_variable = (500, 6.4, 1);

    // this is called deconstructing
    let (x, y, z) = tup_variable

    println!("The value of y is: {y}");
}
```

we can also access elements in tuple via index.

## The Array Type

Unlike a tuple, every element of an array must have the same type. Unlike arrays in some other languages, arrays in Rust have a fixed length. if non fixed length is reqired use vector.Arrays are useful when you want your data allocated on the stack rather than the heap

```rust
fn main() {
    let a = [1, 2, 3, 4, 5];
}
```

```rust
let months = ["January", "February", "March", "April", "May", "June", "July",
              "August", "September", "October", "November", "December"];
```

```rust
let a: [i32; 5] = [1, 2, 3, 4, 5];
```

You can also initialize an array to contain the same value for each element by specifying the initial value, followed by a semicolon, and then the length of the array in square brackets, as shown here:

```rust
let a = [3; 5];
```