# Closest Number Divisible by m

## Problem Statement

Given two integers `n` and `m` (where `m != 0`), find the number closest to `n` that is divisible by `m`.

If there are two such numbers that are equally close to `n`, return the one that has the **greater absolute value** (i.e., the one farther from zero).

---

## Input

- Two integers `n` and `m`  
  (where: -10⁹ ≤ `n` ≤ 10⁹ and `m ≠ 0`)

## Output

- A single integer: the number closest to `n` and divisible by `m`.

---

## Examples

### Example 1
**Input:**  
`n = 13`, `m = 4`  
**Output:**  
`12`

**Explanation:**  
Multiples of 4 near 13 are 12 and 16.  
- `|13 - 12| = 1`  
- `|13 - 16| = 3`  
→ Closest is `12`

---

### Example 2  
**Input:**  
`n = -15`, `m = 6`  
**Output:**  
`-18`

**Explanation:**  
Multiples of 6 near -15: `-12`, `-18`  
- Both are 3 units away from -15  
- But `|-18| > |-12|`, so answer is `-18`

---

## Approach

1. Start by calculating the closest multiple of `m` less than or equal to `n`:
   \[
   x = \left\lfloor \frac{n}{m} \right\rfloor \times m
   \]

2. In Go, integer division truncates toward zero, so for negative numbers, we adjust:
   - If `n < 0` and `n % m != 0`, subtract 1 from the quotient before multiplying.

3. Calculate the next higher multiple:
   \[
   y = x + m
   \]

4. Compare distances of `x` and `y` from `n`.  
   - If one is closer, return it.  
   - If equally close, return the one with greater absolute value.

---

## Constraints

- `m` is never zero.
- `n` and `m` can be positive or negative.