# Find All Duplicates in an Array

## Question
Given an array `arr[]` of size `n`, containing elements from the range 1 to n, and each element appears at most twice,  
return an array of all the integers that appear twice.  

**Note**: You can return the elements in any order but the driver code will print them in sorted order.

---

## Examples

**Example 1**  
Input:  
5  
2 3 1 2 3  
Output:  
2 3  
Explanation:  
2 and 3 occur more than once in the given array.

---

**Example 2**  
Input:  
3  
3 1 2  
Output:  
<empty output>  
Explanation:  
There is no repeating element in the array, so the output is empty.

---

## Constraints
- 1 < n ≤ 10⁶  
- 1 ≤ arr[i] ≤ n  

---

## Expected Complexities
- Time Complexity: O(n)  
- Auxiliary Space: O(1)  

---

## My Approach
1. Iterate over the array:
   - For each value, take its absolute value.
   - Map it to an index by subtracting 1 (`val - 1`).
   - If the number at that index is already negative, it means the value has been seen before → it’s a duplicate.
   - Otherwise, mark it as visited by making the value at that index negative.
2. Add duplicates to a result list.
3. Return the result list.

**Why this works**:  
By marking visited indices as negative, we track occurrences without using extra space, and since each number is between `1` and `n`, mapping `value → index` is direct.

---

## Complexity Analysis
- **Time Complexity**: O(n) — Single pass over the array.
- **Space Complexity**: O(1) — Modifies the array in place without extra data structures.