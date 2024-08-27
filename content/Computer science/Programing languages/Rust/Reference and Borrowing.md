# Reference and Borrowing

Because of the concept of ownership , once we call a function with a variable or reassign to another variable , the ownership of variable is transferred to the function or the new variable , in case of function if we need to get the value back we need to return it and reassign. this is not a cool way to do it . reference is used to solve this problem. references allow you to take the value without taking the ownership

### The Rules of References

- At any given time, you can have *either* one mutable reference *or* any number of immutable references.
- References must always be valid.

![Untitled](Reference%20and%20Borrowing%20009954406ec447b2aad08aaa02f489da/Untitled.svg)

here’s a code without reference : 

```rust
fn main() {
    let s1 = String::from("hello");
    let (s2, len) = calculate_length(s1);
    println!("The length of '{s2}' is {len}.");
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len();
    (s, length)
}
```

using reference we can rewrite it like this,

```rust
fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1);
    println!("The length of '{s1}' is {len}.");
}

fn calculate_length(s: &String) -> usize {
    s.len()
}
```

here we are using `&` to make the variable as reference. note that we pass `&s1` into`calculate_length` and, in its definition, we take `&String` rather than `String`.The opposite of referencing is dereferencing , we use `*` for dereferencing.

The `&s1` syntax lets us create a reference that *refers* to the value of `s1`but does not own it. Because it does not own it, the value it points to will not be dropped when the reference stops being used.
the signature of the function uses `&` to indicate that the type of the parameter `s` is a reference.

When functions have references as parameters instead of the actual values, we won’t need to return the values in order to give back ownership, because we never had ownership. We call the action of creating a reference *borrowing*. Just as variables are immutable by default, so are references. We’re not allowed to modify something we have a reference to.

## **Mutable References**

if we want to modify a referred variable , we need to use mutable reference

```rust
fn main() {
    let mut s = String::from("hello");
    change(&mut s);
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
```

First we change `s` to be `mut`. Then we create a mutable reference with `&mut s` where we call the `change` function, and update the function signature to accept a mutable reference with `some_string: &mut String`. This makes it very clear that the `change` function will mutate the value it borrows.

Mutable references have one big restriction: if you have a mutable reference to a value, you can have no other references to that value. This code that attempts to create two mutable references to `s` will fail. this restriction is to avoid lot of issues like data races.

```rust
let mut s = String::from("hello");

    let r1 = &mut s;
    let r2 = &mut s;

    println!("{}, {}", r1, r2);
    
    // this will fail because both r1 and r2 are mutable reference of the same variable s
```

If we need to use 2 mutable reference we need to do it like this

```rust
let mut s = String::from("hello");

    {
        let r1 = &mut s;
    } // r1 goes out of scope here, so we can make a new reference with no problems.

    let r2 = &mut s;
```

here’s another problem

```rust
 let mut s = String::from("hello");

    let r1 = &s; // no problem
    let r2 = &s; // no problem
    let r3 = &mut s; // BIG PROBLEM

    println!("{}, {}, and {}", r1, r2, r3);
```

Even though borrowing errors may be frustrating at times, remember that it’s the Rust compiler pointing out a potential bug early (at compile time rather than at runtime) and showing you exactly where the problem is. Then you don’t have to track down why your data isn’t what you thought it was.

here’s another kind of problem 

```rust
fn dangle() -> &String { // dangle returns a reference to a String

    let s = String::from("hello"); // s is a new String

    &s // we return a reference to the String, s
} // Here, s goes out of scope, and is dropped. Its memory goes away.
  // Danger!
```

Because `s` is created inside `dangle`, when the code of `dangle` is finished,`s` will be deallocated. But we tried to return a reference to it. That means this reference would be pointing to an invalid `String`. That’s no good! Rust won’t let us do this.