# Mean And Median of Array

## Question
Difficulty: Easy  
Accuracy: 44.05%  
Submissions: 90K+  
Points: 2  

Given an array of positive integers `arr[]`, find the **mean** and **median** of the array.  

- **Mean** is the average value of the given array (floor value).  
- **Median** is the middle element in sorted order if the size is odd, or the floor of the average of two middle elements if the size is even.  

### Note:
Return the **floor value** of both mean and median.

---

## Examples

**Example 1**  
Input:  
5  
1 2 19 28 5  
Output:  
11 5  
**Explanation:**  
Mean = floor((1 + 2 + 19 + 28 + 5) / 5) = floor(55 / 5) = 11  
Median = 5 (middle element after sorting → [1, 2, 5, 19, 28])  

---

**Example 2**  
Input:  
4  
2 8 3 4  
Output:  
4 3  
**Explanation:**  
Mean = floor((2 + 8 + 3 + 4) / 4) = floor(17 / 4) = 4  
Median = floor((3 + 4) / 2) = floor(3.5) = 3 (sorted array → [2, 3, 4, 8])  

---

## Constraints
- 1 < arr.size() ≤ 10⁶  
- 1 ≤ arr[i] ≤ 10⁶  

---

## My Approach
1. **For Mean**:
   - Sum all elements of the array.
   - Divide the sum by the number of elements using integer division (`//`) to get the floor value.
2. **For Median**:
   - Sort the array.
   - If the number of elements is odd, pick the middle element.
   - If even, take the average of the two middle elements using integer division (`//`) for floor value.
3. Sorting ensures correct median calculation, and integer division ensures we get the floor value without using extra functions.

---

## Complexity Analysis
- **Time Complexity**:
  - Sorting: O(n log n)
  - Mean calculation: O(n)  
  Overall: **O(n log n)**  
- **Space Complexity**:
  - Sorting in-place: O(1) additional space (apart from input array storage).
