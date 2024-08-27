# C

C is a System programming language used for writing operating systems, compilers, and other performance-critical applications.

- **Low-level access to memory**: Direct manipulation of hardware and memory.
- **Simple and efficient**: Fast execution and minimalistic syntax.
- **Portable**: C code can be compiled and run on many different types of computers.
- **Rich set of operators and libraries**: Extensive collection of functions for performing various tasks.

Structure : 

```c
#include <stdio.h>  // Preprocessor directive for standard input-output header file

int main() {        // Main function where program execution begins
    printf("Hello, World!\n");  // Function to print text to the console
    return 0;       // Return statement to indicate the program ended successfully
}
```

The first line `#include <stdio.h>`  is a preprocessor line, `stdio.h` is a header file where the code for io is already written , when the processor executes the code the preprocessor will replace the line with required code.

Below is how gcc works, here instead of directly converting code to machine code , an assembler converts it to assembly language first.

![screenshot_2024-08-25_11-24-25.png](C%2071f5da72448042e8bf542aa14be31781/screenshot_2024-08-25_11-24-25.png)

Types of errors : 

1. Compile time error
2. Run time error
3. Semantic error

**Where and How Are Program Data Stored?**

Once a program is executed , the OS loads the program into the RAM and execute it line by line , in RAM there will be a space allocated to the entire program. within the allocated space , main function will have a set of space allocated. variables will be allocated space in this memory

Variables will be allocated space based on their type. main memory is made of several units of 8bits (1byte) (typically)

lets say we want to store a datatype of int which requires 32bits , 4 units of 8 bit memory will be allocated with an address to that location.

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled.png)

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled%201.png)

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled%202.png)

**Characters** 
 

Everything in memory is stored as a byte , so in order to store a character in memory we will use character encoding schemes  ex ASCII, which is like a lookup table . 

for ascii there will be characters associated with 0 to 127 . which is an 8 but number , so one character will be 1 byte.

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled%203.png)

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled.gif)

**Type conversion**

Implicit conversion - lets say we are trying to add a `short int` and a normal `int` , ALU in CPU can only add numbers with same byte values , so the lower type will be updated to the higher one , i.e. `short int` will be updated to `int` by adding trailing zeros and sum will be calculated as an `int`

lets say we are dividing an `int` of 100 with an `int` of 3 , and assigning it to a `float` then the result will be 33.000.. , if we change `int` of 3 to `double` of 3 then the `int` of 100 will be promoted , hence the final value as `float` will be 33.333..

**Explicit conversion (type casting)**

here we are telling the compiler explicitly to change the type.

example : 

```java
int a = 65;
// here 65 will be converted to 'A' before assigning to ch
char ch = (char) a;

```

**Operators**

There are 2 types of operators 

- unary  , ex : `-a` here `-` is the operator

-binary , ex : `a+b` here `+` is the operator

Expressions are created using operation , even `25` is an expression without any operator.

Another kind of type

- Relational (`<`, `>` ..)
- Arithmetic (`+`, `-` ..)
- Logical (`&&` , `||` ..)
- Assignment operator (`=` , `==` ..)

> Lot of information is available in the reading of week 2 in bits course
> 

## Statements and Blocks

Statement is a single expression/statement that ends with a `;`

Block is a set of statement

## Switch

switch is like an efficient way to write several if-else , we add breaks in each case so that once a case is matched it can execute the code and come out of the switch . in C if there is not break ,once a case is matched it will go to the next case and even if the case is not matching , it will execute.

## Function

A self contained program segment that carry out some specific well defined tasks.Functions make the code modular.

declaration ⇒ <return_type> <function_name> (list_of_params_comma_seperated_with_type)

body ⇒ defined within curly brackets

in C we can declare function and define it later , while doing so the name of parameters don't have to match

while functions have lots of benefits , it will reduce the speed of execution because of additional overheads

in c if we are not declareing functions we need to write the function before the function where it is called

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled%204.png)

## Scope and Storage class

The scope of a variable is the section of a program where that variable can be accessed

```c
#include <stdio.h>

void call_func(int a);

int main() {

  int a = 5;
  call_func(a);
  printf("%d", a);
}

void call_func(int a) {
  printf("%d", a);
  a = 10;
  printf("%d", a);
}
```

in above code output will be 5, 10 ,5 
means variable a is not referenced.

There are many types of storage classes like , static, registrar

## Memory allocated to program

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled%205.png)

- The **Stack** segment stores the memory allocated to the functions and variables declared inside the functions.
- The **Heap** segment stores dynamically allocated memory, i.e., variables declared using **malloc**, **calloc**, and **realloc** functions.
- The **Data** segment stores global and static variables declared in the program.
- The **Text** segment stores the executable code that we are executing.

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled%206.png)

## Auto variables or Local Variables

variables that are declared within `{ }` is called auto variables or local variables. all the local variables are stored in stack. these variables will be available in memory as long as the scope exists.

Initial value on declaration is a garbage value

## Global variables

These are the variables that are declared outside all the functions, it is accessible everywhere and the initial value is 0, it is stored in the data segment memory.

## Static variables

it is declared using static. the initial value will be 0 , and stored in the data segment.They reside in the memory as long as the program is running. all static variables are initialized only once.

Static variable is only accessible within the scope.

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled%207.png)

| **Storage Classes** | **Storage Place** | **Default Value** | **Scope** | **Lifetime** |
| --- | --- | --- | --- | --- |
| auto | stack | garbage value | Local to a block or a function | Until the block or the function in which it is declared is executing. |
| global | data | zero | Global to all functions | Until the entire program finishes execution. |
| static | data | zero | Local to a block or a function | Until the entire program finishes execution, they also retain values between function calls. |

## Arrays

Arrays are collection of elements , it allows easier storage and access of data. its useful of easier searching , organizing data elements ,databases and performing matrix operations

syntax <Storage-class> <datatype> <array_name>[<size>]

```java
// declaration without storage-class
int arr[5];

// initialization
int intArr[6] = {1,2,3,4,5,6};

// partial initialization , here rest of the values will be 0.00
float floatArr[10] = {1.3,1.45,5.67}

// without size , here the size will be 3
int arr[] = {10,15,20}

// access
arr[3]
```

Ordinary values are passed by value while arrays are passed by reference.

![Untitled](C%2071f5da72448042e8bf542aa14be31781/Untitled%208.png)

## Linear search

it is searching through array for an element by checking each element one by one until we find it

## Selection sort

we will iterate the array and move the smallest element to the beginning of the array