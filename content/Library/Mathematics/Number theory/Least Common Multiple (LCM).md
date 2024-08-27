# Least Common Multiple (LCM)

The LCM of two or more integers is the smallest positive integer that is divisible by each of the numbers. For example, the LCM of 12 and 18 is 36.

![Untitled](Least%20Common%20Multiple%20(LCM)%2058ffd9f4f95c47a59800923d25065042/Untitled.png)

### **Computing the LCM**

1. The LCM in the introduction was computed by listing the multiples of each number and searching for the first integer on every list. (The LCM always exists because the product of a list of numbers is divisible by each of the numbers in the list.) This is a very ineffective way of computing the LCM in general.
2. Using prime factorization , The primes in the factorization of the LCM are the primes that appear in the factorizations of at least one member of the list, and their exponent is the maximum of the exponents that appear in the individual factorizations.
    
    ![Untitled](Least%20Common%20Multiple%20(LCM)%2058ffd9f4f95c47a59800923d25065042/Untitled%201.png)
    
    i.e. . if 
    
    $$
    a=p_1^{α1}p_2^{α2}…p_k^{αk} \\b=p_1^{β1}p_2^{β2}…p_k^{βk}
    $$
    
    then 
    
    $$
    lcm(a,b) = p_1^{max(α1,β1)}p_2^{max(α2,β2)}...p_k^{max(αk,βk)}
    $$
    
    ### **Connection with the GCD**
    
    The two concepts are intimately related; in particular, they satisfy the following theorem
    
    $$
    gcd(a,b)×lcm(a,b)=ab
    $$