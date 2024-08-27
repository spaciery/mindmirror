# Set Operations

## Union

The set of elements that are in A or B or both, def : $A \cup B = \{x∣x  \in A \  or \ x \in B\}$ , common elements are not repeated.

example : 

$$
A = \{1,2,3\} \\ B = \{2,3,4\} \\ A \cup B = \{1,2,3,4\}
$$

## Intersection

A set that contains only the elements that are common to both A and C. 

def :  $A \cup B=\{x∣x\ in \  A \  and \  x \in B\}$

## **Difference**

A set that contains the elements that are in AAA but not in BBB.
def : $A - B=\{x∣x\ in \  A \  and \  x \notin B\}$

example : 

$$
A = \{1,2,3\} \\ B = \{3,4,5\} \\ A - B = \{1,2\}
$$

## Symmetric Difference

The symmetric difference of two sets A and B is a set that contains the elements that are in either A or B but not in their intersection. The symmetric difference is denoted by $A \Delta B$.

def : 

$$
A \Delta B= (A−B)\cup(B−A) \\ 
A \Delta B = \{ x \mid x \in A \text{ or } x \in B \text{ but not in both } A \text{ and } B \}
$$

in simple terms commons elements are avoided.

## **Complement**

The complement of a set A, denoted by $A^\complement$, is the set of all elements in the universal set U that are not in A.

def : $A^\complement =\{x∣x \in U \  and \  x\notin A\}$ 

## Cartesian Product

The Cartesian product of two sets A and B is the set of all ordered pairs where the first element is from A and the second element is from B . The Cartesian product is denoted by $A \times B.$

def : $A\times B=\{(a,b) ∣ a \in A \ and \  b \in B\}$

example : 

$$
A = \{1, 2\} \\ B = \{3, 4\} \\
A \times B = \{(1, 3), (1, 4), (2, 3), (2, 4)\}
$$

Law : De morgans law

$$
(A \cup B)^\complement = A^\complement \cap B^\complement \\ (A \cap B)^\complement = A^\complement \cup B^\complement
$$