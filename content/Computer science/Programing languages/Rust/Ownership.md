# Ownership

*Ownership* is a set of rules that govern how a Rust program manages memory.

## stack and heap

stack is a last in first out way of memory , putting something into the stack is very fast. heap is less organized , while putting something into the heap , the system need to find space big enough and do book keeping for future allocations , this is slow. also system need to store the address of the objects in heap memory.

Keeping track of what parts of code are using what data on the heap, minimizing the amount of duplicate data on the heap, and cleaning up unused data on the heap so you don’t run out of space are all problems that ownership addresses. Once you understand ownership, you won’t need to think about the stack and the heap very often, but knowing that the main purpose of ownership is to manage heap data can help explain why it works the way it does.

## **Ownership Rules**

- Each value in Rust has an *owner*.
- There can only be one owner at a time.
- When the owner goes out of scope, the value will be dropped.

## Memory and Allocation

For immutable variables like String that can grow and hold big values, we need to allocate the memory in heap during the runtime , and once its use is over we need to free up memory.

in other languages we use GC, or do this manually which is a big headache . In rust the memory is automatically returned once the variable that owns it goes out of scope.

```rust
{
        let s = String::from("hello"); // s is valid from this point forward
        // do stuff with s
}                                  
// this scope is now over, and s is no longer valid
```

There is a natural point at which we can return the memory our `String` needs to the allocator: when `s` goes out of scope. When a variable goes out of scope, Rust calls a special function for us. This function is called [`drop`](https://doc.rust-lang.org/std/ops/trait.Drop.html#tymethod.drop), and it’s where the author of `String` can put the code to return the memory. Rust calls `drop` automatically at the closing curly bracket.

## **Variables and Data Interacting with Move**

```rust
    let x = 5;
    let y = x;
```

in this code the type is integers with a fixed size , here y will have a value of 5 as expected 

```rust
    let s1 = String::from("hello");
    let s2 = s1;
```

here the same wont happen

for s1 this is what happens while initialization

![Untitled](Ownership%203c93335f8f3c4dcfb694b5fcfd6093a1/Untitled.svg)

s1 will be in stack and there will be a pointer which points to the value in the heap.

when we copy s1 to s2 the stack data is getting copied , that is the pointer of the s1 and s2 will be same

![Untitled](Ownership%203c93335f8f3c4dcfb694b5fcfd6093a1/Untitled%201.svg)

Rust do it this way because it is more memory friendly.

Now another problem arises , when s1 and s2 goes out of scope , rust will call drop function to de allocate memory , but both variable have same pointer . this will result typically result in a double free error (Freeing memory twice can lead to memory corruption) but rust will consider s1 as invalid once the data is copied to s2. we will get an error if we try to use s1 once it is invalid.
the concept of copying the pointer, length, and capacity without copying the data probably sounds like making a shallow copy. But because Rust also invalidates the first variable, instead of being called a shallow copy, it’s known as a *move*.

If we *do* want to deeply copy the heap data of the `String`, not just the stack data, we can use a common method called `clone`.

## **Stack-Only Data: Copy**

```rust
let x = 5;
    let y = x;

    println!("x = {x}, y = {y}");
```

`x` is still valid and wasn’t moved into `y`.

The reason is that types such as integers that have a known size at compile time are stored entirely on the stack, so copies of the actual values are quick to make. That means there’s no reason we would want to prevent `x` from being valid after we create the variable `y`. In other words, there’s no difference between deep and shallow copying here, so calling `clone` wouldn’t do anything different from the usual shallow copying, and we can leave it out.

Rust has a special annotation called the `Copy` trait that we can place on types that are stored on the stack, as integers are . If a type implements the `Copy` trait, variables that use it do not move, but rather are trivially copied ,making them still valid after assignment to another variable.

Rust won’t let us annotate a type with `Copy` if the type, or any of its parts, has implemented the `Drop` trait. If the type needs something special to happen when the value goes out of scope and we add the `Copy` annotation to that type, we’ll get a compile-time error.

## In context of functions

The mechanics of passing a value to a function are similar to those when assigning a value to a variable. Passing a variable to a function will move or copy, just as assignment does.

```rust
fn main() {
    let s = String::from("hello");  // s comes into scope

    takes_ownership(s);             // s's value moves into the function...
                                    // ... and so is no longer valid here

    let x = 5;                      // x comes into scope

    makes_copy(x);                  // x would move into the function,
                                    // but i32 is Copy, so it's okay to still
                                    // use x afterward

} // Here, x goes out of scope, then s. But because s's value was moved, nothing
  // special happens.

fn takes_ownership(some_string: String) { // some_string comes into scope
    println!("{some_string}");
} // Here, some_string goes out of scope and `drop` is called. The backing
  // memory is freed.

fn makes_copy(some_integer: i32) { // some_integer comes into scope
    println!("{some_integer}");
} // Here, some_integer goes out of scope. Nothing special happens.
```

If we tried to use `s` after the call to `takes_ownership`, Rust would throw a compile-time error. 

```rust
fn main() {
    let s1 = gives_ownership();
	  // gives_ownership moves its return value into s1
    let s2 = String::from("hello");
    // s2 comes into scope
    let s3 = takes_and_gives_back(s2);
    // s2 is moved into takes_and_gives_back, which also moves its return value into s3
    
} 
// Here, s3 goes out of scope and is dropped. s2 was moved, so nothing
// happens. s1 goes out of scope and is dropped.

fn gives_ownership() -> String {             // gives_ownership will move its
                                             // return value into the function
                                             // that calls it

    let some_string = String::from("yours"); // some_string comes into scope

    some_string                              // some_string is returned and
                                             // moves out to the calling
                                             // function
}

// This function takes a String and returns one
fn takes_and_gives_back(a_string: String) -> String { // a_string comes into
                                                      // scope

    a_string  // a_string is returned and moves out to the calling function
}
```

Concept of ownership will be too tedious in some cases , what if we don't want the function to take the ownership . Luckly ,Rust has a feature for using a value without transferring ownership, called *references*.