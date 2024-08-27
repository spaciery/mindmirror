# Struct

A *struct*, or *structure*, is a custom data type that lets you package together and name multiple related values that make up a meaningful group. If you’re familiar with an object-oriented language, a *struct* is like an object’s data attributes.

 Unlike with tuples, in a struct you’ll name each piece of data so it’s clear what the values mean. Adding these names means that structs are more flexible than tuples: you don’t have to rely on the order of the data to specify or access the values of an instance.

Here's how you write a struct

```rust
struct User {
    active: bool,
    username: String,
    email: String,
    sign_in_count: u64,
}
```

```rust
fn main() {

		// a mutable instance of User 
    let mut user1 = User {
        active: true,
        username: String::from("someusername123"),
        email: String::from("someone@example.com"),
        sign_in_count: 1,
    };
		// updating the value
    user1.email = String::from("anotheremail@example.com");
}
```

Note that the entire instance must be mutable; Rust doesn’t allow us to mark only certain fields as mutable.

```rust
fn build_user(email: String, username: String) -> User {
    User {
        active: true,
        username: username,
        email: email,
        sign_in_count: 1,
    }
}
```

this can be written in short as (*field init shorthand* syntax)

```rust
fn build_user(email: String, username: String) -> User {
    User {
        active: true,
        username,
        email,
        sign_in_count: 1,
    }
}
```

this is how to create a struct instance from another instance

```rust
let user2 = User {
        active: user1.active,
        username: user1.username,
        email: String::from("another@example.com"),
        sign_in_count: user1.sign_in_count,
    };
```

this is the shorthand way to do it 

```rust
    let user2 = User {
        email: String::from("another@example.com"),
        ..user1
    };
```

Note that the struct update syntax uses `=` like an assignment; this is because it moves the data In this example, we can no longer use`user1` as a whole after creating `user2` because the `String` in the`username` field of `user1` was moved into `user2`. If we had given `user2` new `String` values for both `email` and `username`, and thus only used the`active` and `sign_in_count` values from `user1`, then `user1` would still be valid after creating `user2`. Both `active` and `sign_in_count` are types that implement the `Copy` trait so it is not a move since it is in the stack.

Rust also supports structs that look similar to tuples, called *tuple structs*. Tuple structs have the added meaning the struct name provides but don’t have names associated with their fields; rather, they just have the types of the fields. Tuple structs are useful when you want to give the whole tuple a name and make the tuple a different type from other tuples, and when naming each field as in a regular struct would be verbose or redundant.

```rust
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

fn main() {
    let black = Color(0, 0, 0);
    let origin = Point(0, 0, 0);
}
```

You can also define structs that don’t have any fields! These are called *unit-like structs* because they behave similarly to `()`, the unit type that we mentioned in [“The Tuple Type”](https://doc.rust-lang.org/book/ch03-02-data-types.html#the-tuple-type) section. Unit-like structs can be useful when you need to implement a trait on some type but don’t have any data that you want to store in the type itself. 

```rust
struct AlwaysEqual;

fn main() {
    let subject = AlwaysEqual;
}
```

Normally structs need to own all the fields inside it , so that the fields are valid till struct is valid . inorder to use reference inside struct , we need to use lifetime 

this code wont work , because weve not added lifetime 

```rust
struct User {
    active: bool,
    username: &str,
    email: &str,
    sign_in_count: u64,
}

fn main() {
    let user1 = User {
        active: true,
        username: "someusername123",
        email: "someone@example.com",
        sign_in_count: 1,
    };
}
```

## **Method Syntax**

```java
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

// impl is used to define the area method within the Rectangle type
impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };

    println!(
        "The area of the rectangle is {} square pixels.",
        rect1.area()
    );
}
```

*Methods* are similar to functions: we declare them with the `fn` keyword and a name, they can have parameters and a return value, and they contain some code that’s run when the method is called from somewhere else. Unlike functions, methods are defined within the context of a struct (or an  enum or a trait object and their first parameter is always `self`, which represents the instance of the struct the method is being called on.