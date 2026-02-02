---
link: https://leetcode.com/problems/valid-sudoku/
tags:
  - Medium
  - Array
  - Hash_Table
  - Matrix
---
## Description
Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

1. Each row must contain the digits `1-9` without repetition.
2. Each column must contain the digits `1-9` without repetition.
3. Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9` without repetition.

**Note:**

- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

**Example 1:**

![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

```
Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
```

**Example 2:**

```
Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
```

**Constraints:**

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` is a digit `1-9` or `'.'`.

## Solution

### Approach 1

**Hash Table Method**: Use hash tables to record digits that have appeared in each row, column, and 3x3 sub-box.

**Principle:**
Check the validity of rows, columns, and 3x3 sub-boxes in a single traversal. The outer loop variable `i` controls checking the i-th row, i-th column, and i-th sub-box, while the inner loop variable `j` iterates through 9 positions in that row/column/sub-box.

**Steps:**
1. Create three hash tables `row`, `col`, `squ` to record digits in the current row, column, and sub-box
2. Outer loop `i` from 0 to 8, representing the i-th row, i-th column, and i-th sub-box
3. Inner loop `j` from 0 to 8:
   - Check if `board[i][j]` (j-th element in row i) is duplicate in `row`
   - Check if `board[j][i]` (j-th element in column i) is duplicate in `col`
   - Check if `board[j/3+3*(i/3)][j%3+3*(i%3)]` (j-th element in sub-box i) is duplicate in `squ`
4. Clear all three hash tables after each outer loop iteration

**Example:**
- When `i=0`: check row 0, column 0, and top-left 3x3 sub-box
- Sub-box index mapping: `i=4` corresponds to center sub-box, `j` iterates through 9 positions
  - `j=0` → `board[1][3]`, `j=4` → `board[2][4]`, `j=8` → `board[3][5]`

```embed-go
PATH: "vault://leetcode/0001-0100/0036_valid_sudoku/solution.go"
TITLE: "leetcode 36. Valid Sudoku"
```

### Approach 2

**Bit Manipulation Method**: Use integer bits to mark whether a digit has appeared, replacing hash tables.

**Principle:**
Use different bits of an integer to represent whether digits 1-9 have appeared. For example, bit 1 represents digit 1, bit 2 represents digit 2, and so on. Use bitwise AND `&` to check and bitwise OR `|` to set the occurrence of digits.

**Steps:**
1. Use three integers `row`, `col`, `squ` as bitmaps
2. For digit `d`, calculate the corresponding bitmask using `1 << (d - '.')`
3. Check duplicate: `row & (1 << (d - '.'))` non-zero means the digit has appeared
4. Record occurrence: `row |= (1 << (d - '.'))` sets the corresponding bit to 1
5. Reset all three integers to 0 after each outer loop iteration

**Example:**
- Digit `'5'` corresponds to bitmask `1 << ('5' - '.')` = `1 << 7` = `128`
- If `row = 0b10000000`, encountering `'5'` again gives `row & 128 = 128 ≠ 0`, duplicate detected

> **Note**: Using `'.'` as the base for offset calculation, `'.'` has ASCII value 46, `'1'` is 49, so digits 1-9 correspond to bits 3-11.

```embed-go
PATH: "vault://leetcode/0001-0100/0036_valid_sudoku/solution2.go"
TITLE: "leetcode 36. Valid Sudoku"
```
