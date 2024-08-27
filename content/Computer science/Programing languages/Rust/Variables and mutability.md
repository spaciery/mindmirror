# Variables and mutability

When a variable is immutable, once a value is bound to a name, you can’t change that value. we will get a compile time error if we try to assign values to an immutable variable. we can use `mut` keyword to make a variable mutable.

Ultimately, deciding whether to use mutability or not is up to you and depends on what you think is clearest in that particular situation.

## Constants

Like immutable variables, *constants* are values that are bound to a name and are not allowed to change, but there are a few differences between constants and variables.

You declare constants using the `const` keyword instead of the `let` keyword and you cannot use `mut` .

ex : 

```rust
const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;
```

## Shadowing

you can declare a new variable with the same name as a previous variable. Rustaceans say that the first variable is *shadowed* by the second, which means that the second variable is what the compiler will see when you use the name of the variable. we need to use `let` keyword for shadowing By using `let`, we can perform a few transformations on a value but have the variable be immutable after those transformations have been completed. We can even change the type with shadowing.

```rust
fn main() {
    let x = 5;
    let x = x + 1;
    
    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}");
    }

    println!("The value of x is: {x}");
    
    let spaces = "   "; // here type is string
    let spaces = spaces.len(); // here type is a number
}
```

This program first binds `x` to a value of `5`. Then it creates a new variable `x` by repeating `let x =`, taking the original value and adding `1` so the value of `x` is then `6`. Then, within an inner scope created with the curly brackets, the third `let` statement also shadows `x` and creates a new variable, multiplying the previous value by `2` to give `x` a value of `12`. When that scope is over, the inner shadowing ends and `x` returns to being `6`. When we run this program, it will output the