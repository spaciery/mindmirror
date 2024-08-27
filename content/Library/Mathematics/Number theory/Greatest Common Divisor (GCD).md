# Greatest Common Divisor (GCD)

**Greatest Common Divisor (GCD)**: The GCD or greatest common factor (GCF) of two or more integers is the largest positive integer that divides each of the numbers without leaving a remainder. For example, the GCD of 12 and 18 is  6

Methods to derive GCD : 

1. By listing out all factors and taking the largest common one .The process may be split up using the method of factor pairs: once one determines a factor $a$ of a number $n$ , the quotient $\frac{n}{a}$ is necessarily a factor as well .this is very ineffective
    
    ![source brilliant.org](Greatest%20Common%20Divisor%20(GCD)%20c96e9824f1d44280acf2096557dae976/Untitled.png)
    
    source brilliant.org
    
2. A somewhat more efficient method is to first compute the [prime factorization](https://brilliant.org/wiki/prime-factorization/) of each number in the set. The resulting GCD is the product of the primes that appear in every factorization, to the smallest exponent seen in the factorizations.
    
    ![Untitled](Greatest%20Common%20Divisor%20(GCD)%20c96e9824f1d44280acf2096557dae976/Untitled%201.png)
    
    ie . if 
    
    $$
    a=p_1^{α1}p_2^{α2}…p_k^{αk} \\b=p_1^{β1}p_2^{β2}…p_k^{βk}
    $$
    
    then 
    
    $$
    gcd(a,b) = p_1^{min(α1,β1)}p_2^{min(α2,β2)}...p_k^{min(αk,βk)}
    $$
    
3. Euclidian algorithm