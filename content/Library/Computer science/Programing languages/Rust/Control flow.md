# Control flow

```rust
fn main() {
    let number = 6;
		
		// condition must return a bool
    if number % 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3, or 2");
    }
    
    // if there are too many if else , we might want to refactor the code to use pattern matching
}
```

## [**Using `if` in a `let`](https://doc.rust-lang.org/book/ch03-05-control-flow.html#using-if-in-a-let-statement) Statement**

Because `if` is an expression, we can use it on the right side of a `let` statement to assign the outcome to a variable

```rust
fn main() {
    let condition = true;
    let number = if condition { 5 } else { 6 };

    println!("The value of number is: {number}");
}
```

```rust
fn main() {
    let condition = true;

		// here the data type is diffrent in if block , so we will get an error
    let number = if condition { 5 } else { "six" };

    println!("The value of number is: {number}");
}
```

## Loop

```rust
fn main() {
    loop {
        println!("again!");
    }
}
```

```rust
fn main() {
    let mut counter = 0;

// loop is also an expression
    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The result is {result}");
}
```

```rust
fn main() {
    let mut count = 0;
    // labelling loop
    'counting_up: loop {
        println!("count = {count}");
        let mut remaining = 10;

        loop {
            println!("remaining = {remaining}");
            if remaining == 9 {
                break;
            }
            if count == 2 {
            // using label to break outer loop
                break 'counting_up;
            }
            remaining -= 1;
        }

        count += 1;
    }
    println!("End count = {count}");
}
```