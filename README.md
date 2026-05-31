# NeetCode Solutions — chiranjeet-baruah

Personal DSA practice repository. Solutions auto-synced from [NeetCode.io](https://neetcode.io). All solutions written in Go.

---

## Progress

| Category | Problems Solved |
|---|---|
| Data Structures & Algorithms | 21 |

**Total:** 21 problems

---

## Problem Index

| Problem | Submissions | Topic |
|---|---|---|
| [anagram-groups](Data%20Structures%20%26%20Algorithms/anagram-groups/) | 4 | Arrays & Hashing |
| [concatenation-of-array](Data%20Structures%20%26%20Algorithms/concatenation-of-array/) | 1 | Arrays & Hashing |
| [daily-temperatures](Data%20Structures%20%26%20Algorithms/daily-temperatures/) | 4 | Stack |
| [duplicate-integer](Data%20Structures%20%26%20Algorithms/duplicate-integer/) | 5 | Arrays & Hashing |
| [evaluate-reverse-polish-notation](Data%20Structures%20%26%20Algorithms/evaluate-reverse-polish-notation/) | 2 | Stack |
| [is-anagram](Data%20Structures%20%26%20Algorithms/is-anagram/) | 7 | Arrays & Hashing |
| [is-palindrome](Data%20Structures%20%26%20Algorithms/is-palindrome/) | 1 | Two Pointers |
| [longest-common-prefix](Data%20Structures%20%26%20Algorithms/longest-common-prefix/) | 8 | Arrays & Hashing |
| [longest-consecutive-sequence](Data%20Structures%20%26%20Algorithms/longest-consecutive-sequence/) | 2 | Arrays & Hashing |
| [max-water-container](Data%20Structures%20%26%20Algorithms/max-water-container/) | 2 | Two Pointers |
| [minimum-stack](Data%20Structures%20%26%20Algorithms/minimum-stack/) | 6 | Stack |
| [products-of-array-discluding-self](Data%20Structures%20%26%20Algorithms/products-of-array-discluding-self/) | 2 | Arrays & Hashing |
| [remove-element](Data%20Structures%20%26%20Algorithms/remove-element/) | 1 | Arrays & Hashing |
| [string-encode-and-decode](Data%20Structures%20%26%20Algorithms/string-encode-and-decode/) | 1 | Arrays & Hashing |
| [three-integer-sum](Data%20Structures%20%26%20Algorithms/three-integer-sum/) | 1 | Two Pointers |
| [top-k-elements-in-list](Data%20Structures%20%26%20Algorithms/top-k-elements-in-list/) | 3 | Heap / Priority Queue |
| [trapping-rain-water](Data%20Structures%20%26%20Algorithms/trapping-rain-water/) | 2 | Two Pointers |
| [two-integer-sum](Data%20Structures%20%26%20Algorithms/two-integer-sum/) | 2 | Arrays & Hashing |
| [two-integer-sum-ii](Data%20Structures%20%26%20Algorithms/two-integer-sum-ii/) | 2 | Two Pointers |
| [valid-sudoku](Data%20Structures%20%26%20Algorithms/valid-sudoku/) | 1 | Arrays & Hashing |
| [validate-parentheses](Data%20Structures%20%26%20Algorithms/validate-parentheses/) | 3 | Stack |

Submission count = number of iterations attempted (includes optimizations, not just accepted solutions).

---

## Repository Structure

```
Data Structures & Algorithms/
  <problem-slug>/
    submission-0.go   ← first attempt
    submission-1.go   ← subsequent iterations
    ...
```

Each submission file is a standalone Go source. Files are numbered in commit order — the highest-numbered file is the most recent iteration.

---

## Running a Solution Locally

Prerequisites: Go 1.21+

```bash
# Copy a submission into a runnable file
cp "Data Structures & Algorithms/two-integer-sum/submission-1.go" /tmp/solution.go

# Add a main() if none exists, then run
cd /tmp && go run solution.go
```

Most submissions define only the solution function (no `main`). To test:

```bash
cat > /tmp/main_test.go << 'EOF'
package main

import "testing"

func TestTwoSum(t *testing.T) {
    result := twoSum([]int{2, 7, 11, 15}, 9)
    if result[0] != 0 || result[1] != 1 {
        t.Fatalf("expected [0 1], got %v", result)
    }
}
EOF

cd /tmp && go test .
```

---

## GitHub Sync

Solutions sync automatically from [NeetCode.io](https://neetcode.io) on every accepted submission.

- Manage sync settings: [neetcode.io/profile/github](https://neetcode.io/profile/github)
- Commit format: `Add: <problem-slug> - submission-<N>`
- Bulk sync available from the GitHub settings page

---

*Auto-synced via [NeetCode GitHub Integration](https://neetcode.io)*
