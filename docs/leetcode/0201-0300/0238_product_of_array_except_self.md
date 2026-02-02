---
link: https://leetcode.com/problems/product-of-array-except-self/
tags:
  - Medium
  - Array
  - Prefix_Sum
---
## Description
Given an integer array `nums`, return *an array* `answer` *such that* `answer[i]` *is equal to the product of all the elements of* `nums` *except* `nums[i]`.

The product of any prefix or suffix of `nums` is **guaranteed** to fit in a **32-bit** integer.

You must write an algorithm that runs in `O(n)` time and without using the division operation.

**Example 1:**

```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
```

**Example 2:**

```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
```

**Constraints:**

- `2 <= nums.length <= 105`
- `-30 <= nums[i] <= 30`
- The input is generated such that `answer[i]` is **guaranteed** to fit in a **32-bit** integer.

**Follow up:** Can you solve the problem in `O(1)` extra space complexity? (The output array **does not** count as extra space for space complexity analysis.)

## Solution

### Approach 1

**Prefix and Suffix Product Arrays**: Pre-compute prefix product and suffix product for each position, multiply them to get the result.

**Principle:**
For position i, the product excluding itself = product of all elements to the left × product of all elements to the right. By pre-computing prefix product array and suffix product array, we can get the left and right products for any position in O(1).

**Steps:**
1. Create `leftAccProd` array for prefix products, `leftAccProd[i]` = nums[0] × ... × nums[i]
2. Create `rightAccProd` array for suffix products, `rightAccProd[i]` = nums[i] × ... × nums[n-1]
3. Result `res[i]` = `leftAccProd[i-1]` × `rightAccProd[i+1]`, handle boundary cases specially

**Example:**
- Input: `nums = [1,2,3,4]`
- Prefix products: `[1, 2, 6, 24]`
- Suffix products: `[24, 24, 12, 4]`
- Calculation: `res[1] = leftAccProd[0] × rightAccProd[2] = 1 × 12 = 12`

**Complexity:** Time O(n), Space O(n)

```embed-go
PATH: "vault://leetcode/0201-0300/0238_product_of_array_except_self/solution.go"
TITLE: "leetcode 238. Product of Array Except Self"
```

### Approach 2

**Prefix Product Array + Suffix Product Variable**: Use one array to store prefix products, use a variable to dynamically compute suffix products, reducing space usage.

**Principle:**
Suffix products can be dynamically accumulated during right-to-left traversal, no need to pre-store the entire array.

**Steps:**
1. First forward traverse to build prefix product array `accProd`
2. Use variable `rightAcc` to accumulate from right to left, compute result `res[i] = accProd[i-1] × rightAcc`
3. Handle boundary positions specially

**Example:**
- Input: `nums = [1,2,3,4]`
- Prefix products: `[1, 2, 6, 24]`
- Right traversal: at i=3, `res[3]=accProd[2]=6`, rightAcc=4
- At i=2, `res[2]=accProd[1]×4=2×4=8`, rightAcc=12
- At i=1, `res[1]=accProd[0]×12=1×12=12`, rightAcc=24
- Finally `res[0]=rightAcc=24`

**Complexity:** Time O(n), Space O(n)

```embed-go
PATH: "vault://leetcode/0201-0300/0238_product_of_array_except_self/solution2.go"
TITLE: "leetcode 238. Product of Array Except Self"
```

### Approach 3

**In-Place Computation (O(1) Extra Space)**: Reuse the result array to store prefix products, then use a variable to compute suffix products.

**Principle:**
Store prefix products directly in the result array, then traverse from right to left, using a suffix product variable to update the result array.

**Steps:**
1. First store prefix products in `res` array
2. Use variable `rightAcc` to traverse from right to left, update `res[i]` to `res[i-1] × rightAcc`
3. Handle boundary positions specially

**Example:**
- Input: `nums = [1,2,3,4]`
- Initialize res as prefix products: `[1, 2, 6, 24]`
- Right traversal update:
  - i=3: `res[3]=res[2]=6`, rightAcc=4
  - i=2: `res[2]=res[1]×4=8`, rightAcc=12
  - i=1: `res[1]=res[0]×12=12`, rightAcc=24
  - Finally: `res[0]=24`
- Output: `[24, 12, 8, 6]`

**Complexity:** Time O(n), Space O(1) (excluding output array)

```embed-go
PATH: "vault://leetcode/0201-0300/0238_product_of_array_except_self/solution3.go"
TITLE: "leetcode 238. Product of Array Except Self"
```
