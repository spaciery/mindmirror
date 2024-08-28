# Discrete mathematics

1.     Book: Kenneth H. Rosen. 2021. *Discrete Mathematics and Its Applications*. 8th Edition. McGraw Hill: McGraw Hill Education (India) Private Limited.

- Proposition - a declarative sentence that is either true or false , but not both. true is demotes by $T$ , false is denoted by $F$ and proposition variables are denoted by $p,q,r...$ or $p_1,p_2,p_3...$
- Compound propositions are new propositions created from excising proposition using logical operators or connectives.

Connectives

- Negation of proposition is denoted by $\neg p$, the truth value of the negation of $p$ is the opposite of truth value of $p$.
- conjunction - lets take 2 proposition say p and q , the conjunction ($\land$) $p \land q$ is *“p and q”* as in logical and. it is true if both p and q are true
- Disjunction ($\lor$) - it is like a logical or, it is false when both p and q are false otherwise true.
- implication or conditional statement, $p \implies q$  (p implies q or p conditional q) is the proposition *“if p, then q” . $p \implies q$* is false when p is true and q is false. here p is called hypothesis and q is called conclusion.
- converse , $q \implies p$ is the converse of $p \implies q$
- contrapositive , $\neg q \implies \neg p$ is the contrapositive of $p \implies q$

### Logical equivalences

a compount proposition that is always true, ie. true for every combination of truth values of its propositional variables is called tautology.

for example take $p \lor \neg p$ , this will always evaluate to true , so it is a tautology .

the opposit of tatutolody , ie always false is called contradiction
for example , $p \land \neg p$ will always evaluate to false , which is a contradiction

compount proposition p and q are called **logicaly equivalent** if their truth table are the same.it is denoted by $ p \equiv q $
example : $ p \implies q $ and $ \neg p \implies \neg q $

DeMorgans laws :
- $\neg(p \land q) \equiv \neg p \lor \neg q$
- $\neg(p \lor q) \equiv \neg p \land \neg q$
- $ p \implies q \equiv \neg p \lor q$

how to state two propositons are not logically equaivalent ?
- find any one row in the truith table where one evaluates to true and other evaluates to false
