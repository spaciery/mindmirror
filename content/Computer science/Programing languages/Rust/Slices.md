# Slices

*Slices* let you reference a contiguous sequence of elements in a [collection](https://doc.rust-lang.org/book/ch08-00-common-collections.html) rather than the whole collection. A slice is a kind of reference, so it does not have ownership.

here’s a way to take a first word out of a string

```rust
fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

// iter is a method that returns each element in a collection and that enumerate wraps the result of iter and returns each element as part of a tuple instead. 

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}
```

this function is returning an integer which hold the index of the first word , but lets say once returning this integer we cleared the string , then the index value is useless. Also lets say we want to return first 2 words , then well have to return a tuple of 2 usize. these methods are very inconvinent. we can use a slice type to solve the problem.

```rust
    let s = String::from("hello world");

    let hello = &s[0..5];
    let world = &s[6..11];
```

working : 

![Untitled](Slices%20bc901b27d9c043bb923319daafa76ce6/Untitled.svg)

this is also valid 

```rust
let s = String::from("hello");
let slice = &s[0..2];
let slice = &s[..2];
```

```rust
let s = String::from("hello");
let len = s.len();
let slice = &s[3..len];
let slice = &s[3..];
```

```rust
let s = String::from("hello");
let len = s.len();
let slice = &s[0..len];
let slice = &s[..];
```

with string slice we can rewrite the first word program as 

```rust
fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}
```

the entire string can be considered as a string slice 

```rust
// this will now work for string and string slices
fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}
```