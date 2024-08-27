# Datastructures

Lists

Arrays
linkedLists

Set

Stack

Queue

Tree

HashTable

Graph

Map

Sort and search algoritm

## Arrays

Collection of elements where size is fixed. It is very fast

in java arrays start with index 0.Once array is defined we cannot change it.We cannot delete an element directly from an array

```java
public class ColorsArrays {
	
	public static void main(String[] args) {
			
			//syntax 1
			String[] colors = new String[5];
			colors[0] = "purple";
			colors[1] = "red";
			colors[2] = "blue";
			
			// printing the array will only give u the object
			// so wrap it in Arrays.toString
			System.out.println(Arrays.toString(colors));
			
			// syntax 2 
			String[] monoColors = {"Black","White"};
			
			// ierations
			for(int i = 0; i < colors.length; i++) {
			 // do something
			}
			
			// iteration in reverse
			for(int i = (colors.length -1); i >= 0; i--) {
			 // do something
			}
			
			// enhanced for
			for(String color : colors) {
					// do something
			}
			
			// using streams
			Arrays.stream(colors).forEach(System.out::println);
			
			// a nice utility
			// Arrays

			
	}
	
}
```

## 2D Arrays

The elements in a 2D array is arranged like a matrix. (0,0) is a the first element.

```java
public class ColorsTwoDArrays {

	public static void main(String[] args) {
	
			char[][] board = new char[5][5];

			board[0][0] = 'X';
			board[0][1] = 'O';
			board[2][0] = 'X';
			
			char[][] boardTwo = new char[][] {
				new char[] {'a','b','c'},
				new char[] {'a','b','c'},
				new char[] {'a','b','c'}
			};

			
		
		  // for looping we need 1 innerloops within a loop
		
		  // printing the array will only give u the object
			// so wrap it in Arrays.deepToString
			System.out.println(Arrays.deepToString(board));
		
	}

}
```

## List

List is an ordered collection.List can grow and shrink in size. some implementation of list are 

- 
- Stack
- Vector
- Others

We can delete elements from array

unlike arrays primitives cannot be added to a list , instead we need to use wrappers

### ArrayList

```java
public class WorkingWithLists {

	public static void main(String[] args) {
			
			// default size of an arrayList is 10
			List<String> colors = new ArrayList<>();
			colors.add("abv");
	}
}
```

### Stack

Stack is a LIFO (last in first out) list.

```java
Stack<String> inteStack = new Stack<>();

stack.push("1");
stack.push("2");
stack.push("3");

// size
stack.size();

// get the top element and remove it 
stack.pop();

// get the top element
stack.peek();

// check if empty
stack.empty();
```

### Vector

Vector is similar to an ArrayList , the only diffrence is it is a synchronized list , so if we need to use thread safe implementations we can use vector

### Queue

This is FIFO list

```java
Queue<Person> supermarket = new PriorityQueue<>();
// also 
Queue<Person> supamkt = new LinkedList<>();

// to add an element
supamkt.add(new Person());

// or we can do this , this is the better way
supamkt.offer(new Person());

// shows the element that we can take out of queue
supamkt.peek();
// get the element that we can take out of queue and removes it from the list
supamkt.poll();

// or we can do this aswell
supamkt.remove();
```

### LinkedList

linked list is made of multiple nodes where we have an head and tail. each node contain a reference to next and previous node. heads prev reference is null and tail’s next reference is null.

since the reference is bidirectional this is a double linked list.

if we need to insert an element  in between we need to pass the index in the add method , internaly it will update the previous and next pointer

### Sets

Sets do not contain duplicates, usually doesn't guarantees order aswell. we can use TreeSet to get the order.

## Map

its a collection of key value pairs.A map cannot contain a duplicate key

there are many implementations like , hashmap , hashtable , linkedHashMap etc .

hashmap allows us to store null values while hashtable dont and hashtables are synchronized