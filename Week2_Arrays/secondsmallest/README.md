# Smallest and Second Smallest Element in Array

## Question
Given an array `arr[]` of integers, your task is to return the smallest and second smallest element in the array.  
If the smallest and second smallest do not exist, return `-1`.

---

## Examples

**Example 1**  
Input:  
5  
2 4 3 5 6  
Output:  
2 3  
Explanation:  
2 and 3 are respectively the smallest and second smallest elements in the array.

---

**Example 2**  
Input:  
3  
1 1 1  
Output:  
-1  
Explanation:  
Only element is 1 which is smallest, so there is no second smallest element.

---

## Constraints
- 1 < arr.size() < 10⁶  
- 1 ≤ arr[i] ≤ 10⁶  

---

## Expected Complexities
- Time Complexity: O(n)  
- Auxiliary Space: O(1)  

---

## My Approach
1. Initialize `smallest` and `second_smallest` to infinity.
2. Traverse the array:
   - If the current element is smaller than `smallest`, update `second_smallest` to `smallest` and `smallest` to current element.
   - Else if the current element is greater than `smallest` but smaller than `second_smallest`, update `second_smallest`.
3. After traversal:
   - If `second_smallest` remains infinity, return `[-1]` since no second smallest element exists.
   - Otherwise, return `[smallest, second_smallest]`.

---

## Complexity Analysis
- **Time Complexity**: O(n) — Single pass over the array.
- **Space Complexity**: O(1) — Constant extra space.