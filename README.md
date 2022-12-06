# solutions
This repo stores my problem-solving related code 

leetcode: [@dimaglushkov](https://leetcode.com/dimaglushkov/) <br>
codeforces: [@dimaglushkov](https://codeforces.com/profile/dimaglushkov)



 ## Table of Contents
[1. tools](#tools) <br>
[2. contests](#contests) <br>
[3. leetcode](#leetcode) <br>
[4. codeforces](#codeforces) <br>
[5. acm.timus](#acmtimus) <br>


## tools 
This repo comes with simple yet useful tools to simplify solutions management. This set of utilities allows one to:
1. Automatically pull problems from leetcode and codeforces and then generate solution templates.
2. Automatically update list of one's solutions
3. Get some simple stats on how solved problems are distributed 

Below are few examples of what these tools can do:
```
# cf - codeforces, lc - leetcode
sol cf/lc p <URLs> - for every given problem url create solution template,
                     pull meta data, update list of solved problems. For leetcode
                     golang solutions create auto tests based on the problem examples.
                              
sol cf/lc d <URL>  - for every given problem url delete solution and 
                     remove correspoding entriy from the list of solved problems
                            
sol readme g .     - generate this README.md file

sol con pre  <cf/lc contest URL> - generate template for contest solutions, update list of 
                                   contests one participated at
                                   
sol con post <cf contest URL>    - for every successfull solution from a given contest
                                   create a solution for corresponding problem from this contest
```

In order to get more information consider checking [tools](https://github.com/dimaglushkov/solutions/tree/master/tools).

## contests

|Top % places by days|
|-|
|![](https://github.com/dimaglushkov/solutions/blob/master/contests/.stats.svg)|


| Name | Rank | Solutions | Date |
|-|-|-|-|
| [Leetcode Weekly Contest 290](https://leetcode.com/contest/weekly-contest-290/) | 4828 / 19491 | [solutions](/contests/lc-weekly-290) | 24 apr 2022 |
| [Leetcode Biweekly Contest 77](https://leetcode.com/contest/biweekly-contest-77) | 6827 / 15519 | [solutions](/contests/lc-biweekly-77) | 30 apr 2022 |
| [Leetcode Weekly Contest 291](https://leetcode.com/contest/weekly-contest-291/) | 6183 / 19491 | [solutions](/contests/lc-weekly-291) | 1 may 2022 |
| [Codeforces Round #790 (Div. 4)](https://codeforces.com/contests/1676) | 7449 / 29363 | [solutions](/contests/cf-round-790) | 10 may 2022 | 
| [Codeforces Round #799 (Div. 4)](https://codeforces.com/contests/1692) | 3515 / 20029 | [solutions](/contests/cf-round-799) | 14 jun 2022 | 
| [Codeforces Round #811 (Div. 3)](https://codeforces.com/contests/1714) | 3935 / 30969 | [solutions](/contests/cf-round-811) | 1 aug 2022 |
| [Leetcode Biweekly Contest 85](https://leetcode.com/contest/biweekly-contest-85) | 2962 / 23603 | [solutions](/contests/lc-biweekly-85) | 20 aug 2022 |
| [Codeforces Round #820 (Div. 3)](https://codeforces.com/contest/1729) | 4833 / 31271 | [solutions](/contests/cf-round-820) | 13 sep 2022 |
| [Codeforces Round #826 (Div. 3)](https://codeforces.com/contest/1741) | 3102 / 17198 | [solutions](/contests/cf-round-826) | 11 oct 2022 |
| [Codeforces Round #828 (Div. 3)](https://codeforces.com/contest/1744) | 2398 / 14451 | [solutions](/contests/cf-round-828) | 16 oct 2022 |
| [Codeforces Round #832 (Div. 2)](https://codeforces.com/contest/1747) | 3708 / 13803 | [solutions](/contests/cf-round-832) | 4 nov 2022 |
| [Leetcode Weekly Contest 320](https://leetcode.com/contest/weekly-contest-320/) | 4955 / 15183 | [solutions](/contests/lc-weekly-320) | 20 nov 2022 |
| [Codeforces Round #835 (Div. 4)](https://codeforces.com/contests/1760) | 702 / 19245 | [solutions](/contests/cf-round-835) | 22 nov 2022 |
| [Leetcode Weekly Contest 321](https://leetcode.com/contest/weekly-contest-321) | 6122 / 17826 | [solutions](/contests/lc-weekly-321) | 27 nov 2022 |
| [Leetcode Weekly Contest 322](https://leetcode.com/contest/weekly-contest-322) | 547 / 19626 | [solutions](/contests/lc-weekly-322) | 4 dec 2022 |


## leetcode

Problems solved in total: 279

|Solutions by difficulty|Solutions by tags|
|-|-|
|![by_difficulty](https://github.com/dimaglushkov/solutions/blob/master/leetcode/.by_difficulty.svg)|![by_tags](https://github.com/dimaglushkov/solutions/blob/master/leetcode/.by_tags.svg)|

|Problem|Solution|Difficulty|Tags|
|-|-|-|-|
| [1. Two sum](https://leetcode.com/problems/two-sum/) | [golang](/leetcode/two-sum/two-sum.go) | Easy | Array, Hash Table |
| [9. Palindrome number](https://leetcode.com/problems/palindrome-number/) | [golang](/leetcode/palindrome-number/palindrome-number.go) | Easy | Math |
| [2. Add two numbers](https://leetcode.com/problems/add-two-numbers/) | [golang](/leetcode/add-two-numbers/add-two-numbers.go) | Medium | Linked List, Math, Recursion |
| [3. Longest substring without repeating characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | [golang](/leetcode/longest-substring-without-repeating-characters/longest-substring-without-repeating-characters.go) | Medium | Hash Table, String, Sliding Window |
| [4. Median of two sorted arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | [golang](/leetcode/median-of-two-sorted-arrays/median-of-two-sorted-arrays.go) | Hard | Array, Binary Search, Divide and Conquer |
| [14. Longest common prefix](https://leetcode.com/problems/longest-common-prefix/) | [golang](/leetcode/longest-common-prefix/longest-common-prefix.go) | Easy | String |
| [13. Roman to integer](https://leetcode.com/problems/roman-to-integer/) | [golang](/leetcode/roman-to-integer/roman-to-integer.go) | Easy | Hash Table, Math, String |
| [28. Implement strstr](https://leetcode.com/problems/implement-strstr/) | [golang](/leetcode/implement-strstr/implement-strstr.go) | Easy | Two Pointers, String, String Matching |
| [459. Repeated substring pattern](https://leetcode.com/problems/repeated-substring-pattern/) | [golang](/leetcode/repeated-substring-pattern/repeated-substring-pattern.go) | Easy | String, String Matching |
| [37. Sudoku solver](https://leetcode.com/problems/sudoku-solver/) | [golang](/leetcode/sudoku-solver/sudoku-solver.go) | Hard | Array, Backtracking, Matrix |
| [36. Valid sudoku](https://leetcode.com/problems/valid-sudoku/) | [golang](/leetcode/valid-sudoku/valid-sudoku.go) | Medium | Array, Hash Table, Matrix |
| [8. String to integer atoi](https://leetcode.com/problems/string-to-integer-atoi/) | [golang](/leetcode/string-to-integer-atoi/string-to-integer-atoi.go) | Medium | String |
| [2042. Check if numbers are ascending in a sentence](https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/) | [golang](/leetcode/check-if-numbers-are-ascending-in-a-sentence/check-if-numbers-are-ascending-in-a-sentence.go) | Easy | String |
| [53. Maximum subarray](https://leetcode.com/problems/maximum-subarray/) | [golang](/leetcode/maximum-subarray/maximum-subarray.go) | Easy | Array, Divide and Conquer, Dynamic Programming |
| [192. Word frequency](https://leetcode.com/problems/word-frequency/) | [bash](/leetcode/word-frequency/word-frequency.sh) | Medium | Shell |
| [338. Counting bits](https://leetcode.com/problems/counting-bits/) | [golang](/leetcode/counting-bits/counting-bits.go) | Easy | Dynamic Programming, Bit Manipulation |
| [33. Search in rotated sorted array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | [golang](/leetcode/search-in-rotated-sorted-array/search-in-rotated-sorted-array.go) | Medium | Array, Binary Search |
| [70. Climbing stairs](https://leetcode.com/problems/climbing-stairs/) | [golang](/leetcode/climbing-stairs/climbing-stairs.go) | Easy | Math, Dynamic Programming, Memoization |
| [118. Pascals triangle](https://leetcode.com/problems/pascals-triangle/) | [golang](/leetcode/pascals-triangle/pascals-triangle.go) | Easy | Array, Dynamic Programming |
| [75. Sort colors](https://leetcode.com/problems/sort-colors/) | [golang](/leetcode/sort-colors/sort-colors.go) | Medium | Array, Two Pointers, Sorting |
| [15. 3sum](https://leetcode.com/problems/3sum/) | [golang](/leetcode/3sum/3sum.go) | Medium | Array, Two Pointers, Sorting |
| [392. Is subsequence](https://leetcode.com/problems/is-subsequence/) | [golang](/leetcode/is-subsequence/is-subsequence.go) | Easy | Two Pointers, String, Dynamic Programming |
| [20. Valid parentheses](https://leetcode.com/problems/valid-parentheses/) | [golang](/leetcode/valid-parentheses/valid-parentheses.go) | Easy | String, Stack |
| [71. Simplify path](https://leetcode.com/problems/simplify-path/) | [golang](/leetcode/simplify-path/simplify-path.go) | Medium | String, Stack |
| [413. Arithmetic slices](https://leetcode.com/problems/arithmetic-slices/) | [golang](/leetcode/arithmetic-slices/arithmetic-slices.go) | Medium | Array, Dynamic Programming |
| [21. Merge two sorted lists](https://leetcode.com/problems/merge-two-sorted-lists/) | [golang](/leetcode/merge-two-sorted-lists/merge-two-sorted-lists.go) | Easy | Linked List, Recursion |
| [141. Linked list cycle](https://leetcode.com/problems/linked-list-cycle/) | [golang](/leetcode/linked-list-cycle/linked-list-cycle.go) | Easy | Hash Table, Linked List, Two Pointers |
| [344. Reverse string](https://leetcode.com/problems/reverse-string/) | [golang](/leetcode/reverse-string/reverse-string.go) | Easy | Two Pointers, String, Recursion |
| [680. Valid palindrome ii](https://leetcode.com/problems/valid-palindrome-ii/) | [golang](/leetcode/valid-palindrome-ii/valid-palindrome-ii.go) | Easy | Two Pointers, String, Greedy |
| [1791. Find center of star graph](https://leetcode.com/problems/find-center-of-star-graph/) | [golang](/leetcode/find-center-of-star-graph/find-center-of-star-graph.go) | Easy | Graph |
| [797. All paths from source to target](https://leetcode.com/problems/all-paths-from-source-to-target/) | [golang](/leetcode/all-paths-from-source-to-target/all-paths-from-source-to-target.go) | Medium | Backtracking, Depth-First Search, Breadth-First Search, Graph |
| [31. Next permutation](https://leetcode.com/problems/next-permutation/) | [golang](/leetcode/next-permutation/next-permutation.go) | Medium | Array, Two Pointers |
| [1920. Build array from permutation](https://leetcode.com/problems/build-array-from-permutation/) | [golang](/leetcode/build-array-from-permutation/build-array-from-permutation.go) | Easy | Array, Simulation |
| [1929. Concatenation of array](https://leetcode.com/problems/concatenation-of-array/) | [golang](/leetcode/concatenation-of-array/concatenation-of-array.go) | Easy | Array |
| [1721. Swapping nodes in a linked list](https://leetcode.com/problems/swapping-nodes-in-a-linked-list/) | [golang](/leetcode/swapping-nodes-in-a-linked-list/swapping-nodes-in-a-linked-list.go) | Medium | Linked List, Two Pointers |
| [11. Container with most water](https://leetcode.com/problems/container-with-most-water/) | [golang](/leetcode/container-with-most-water/container-with-most-water.go) | Medium | Array, Two Pointers, Greedy |
| [923. 3sum with multiplicity](https://leetcode.com/problems/3sum-with-multiplicity/) | [golang](/leetcode/3sum-with-multiplicity/3sum-with-multiplicity.go) | Medium | Array, Hash Table, Two Pointers, Sorting, Counting |
| [1046. Last stone weight](https://leetcode.com/problems/last-stone-weight/) | [golang](/leetcode/last-stone-weight/last-stone-weight.go) | Easy | Array, Heap (Priority Queue) |
| [146. Lru cache](https://leetcode.com/problems/lru-cache/) | [golang](/leetcode/lru-cache/lru-cache.go) | Medium | Hash Table, Linked List, Design, Doubly-Linked List |
| [703. Kth largest element in a stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) | [golang](/leetcode/kth-largest-element-in-a-stream/kth-largest-element-in-a-stream.go) | Easy | Tree, Design, Binary Search Tree, Heap (Priority Queue), Binary Tree, Data Stream |
| [460. Lfu cache](https://leetcode.com/problems/lfu-cache/) | [golang](/leetcode/lfu-cache/lfu-cache.go) | Hard | Hash Table, Linked List, Design, Doubly-Linked List |
| [1480. Running sum of 1d array](https://leetcode.com/problems/running-sum-of-1d-array/) | [golang](/leetcode/running-sum-of-1d-array/running-sum-of-1d-array.go) | Easy | Array, Prefix Sum |
| [1672. Richest customer wealth](https://leetcode.com/problems/richest-customer-wealth/) | [golang](/leetcode/richest-customer-wealth/richest-customer-wealth.go) | Easy | Array, Matrix |
| [2011. Final value of variable after performing operations](https://leetcode.com/problems/final-value-of-variable-after-performing-operations/) | [golang](/leetcode/final-value-of-variable-after-performing-operations/final-value-of-variable-after-performing-operations.go) | Easy | Array, String, Simulation |
| [2114. Maximum number of words found in sentences](https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/) | [golang](/leetcode/maximum-number-of-words-found-in-sentences/maximum-number-of-words-found-in-sentences.go) | Easy | Array, String |
| [1108. Defanging an ip address](https://leetcode.com/problems/defanging-an-ip-address/) | [golang](/leetcode/defanging-an-ip-address/defanging-an-ip-address.go) | Easy | String |
| [1476. Subrectangle queries](https://leetcode.com/problems/subrectangle-queries/) | [golang](/leetcode/subrectangle-queries/subrectangle-queries.go) | Medium | Array, Design, Matrix |
| [347. Top k frequent elements](https://leetcode.com/problems/top-k-frequent-elements/) | [golang](/leetcode/top-k-frequent-elements/top-k-frequent-elements.go) | Medium | Array, Hash Table, Divide and Conquer, Sorting, Heap (Priority Queue), Bucket Sort, Counting, Quickselect |
| [17. Letter combinations of a phone number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | [golang](/leetcode/letter-combinations-of-a-phone-number/letter-combinations-of-a-phone-number.go) | Medium | Hash Table, String, Backtracking |
| [682. Baseball game](https://leetcode.com/problems/baseball-game/) | [golang](/leetcode/baseball-game/baseball-game.go) | Easy | Array, Stack, Simulation |
| [980. Unique paths iii](https://leetcode.com/problems/unique-paths-iii/) | [golang](/leetcode/unique-paths-iii/unique-paths-iii.go), [python](/leetcode/unique-paths-iii/unique-paths-iii.py) | Hard | Array, Backtracking, Bit Manipulation, Matrix |
| [1603. Design parking system](https://leetcode.com/problems/design-parking-system/) | [golang](/leetcode/design-parking-system/design-parking-system.go) | Easy | Design, Simulation, Counting |
| [1260. Shift 2d grid](https://leetcode.com/problems/shift-2d-grid/) | [golang](/leetcode/shift-2d-grid/shift-2d-grid.go) | Easy | Array, Matrix, Simulation |
| [289. Game of life](https://leetcode.com/problems/game-of-life/) | [golang](/leetcode/game-of-life/game-of-life.go) | Medium | Array, Matrix, Simulation |
| [59. Spiral matrix ii](https://leetcode.com/problems/spiral-matrix-ii/) | [golang](/leetcode/spiral-matrix-ii/spiral-matrix-ii.go) | Medium | Array, Matrix, Simulation |
| [1557. Minimum number of vertices to reach all nodes](https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/) | [golang](/leetcode/minimum-number-of-vertices-to-reach-all-nodes/minimum-number-of-vertices-to-reach-all-nodes.go) | Medium | Graph |
| [700. Search in a binary search tree](https://leetcode.com/problems/search-in-a-binary-search-tree/) | [golang](/leetcode/search-in-a-binary-search-tree/search-in-a-binary-search-tree.go) | Easy | Tree, Binary Search Tree, Binary Tree |
| [669. Trim a binary search tree](https://leetcode.com/problems/trim-a-binary-search-tree/) | [golang](/leetcode/trim-a-binary-search-tree/trim-a-binary-search-tree.go) | Medium | Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| [538. Convert bst to greater tree](https://leetcode.com/problems/convert-bst-to-greater-tree/) | [golang](/leetcode/convert-bst-to-greater-tree/convert-bst-to-greater-tree.go) | Medium | Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| [2236. Root equals sum of children](https://leetcode.com/problems/root-equals-sum-of-children/) | [golang](/leetcode/root-equals-sum-of-children/root-equals-sum-of-children.go) | Easy | Tree, Binary Tree |
| [2235. Add two integers](https://leetcode.com/problems/add-two-integers/) | [golang](/leetcode/add-two-integers/add-two-integers.go) | Easy | Math |
| [1470. Shuffle the array](https://leetcode.com/problems/shuffle-the-array/) | [golang](/leetcode/shuffle-the-array/shuffle-the-array.go) | Easy | Array |
| [912. Sort an array](https://leetcode.com/problems/sort-an-array/) | [golang](/leetcode/sort-an-array/sort-an-array.go) | Medium | Array, Divide and Conquer, Sorting, Heap (Priority Queue), Merge Sort, Bucket Sort, Radix Sort, Counting Sort |
| [897. Increasing order search tree](https://leetcode.com/problems/increasing-order-search-tree/) | [golang](/leetcode/increasing-order-search-tree/increasing-order-search-tree.go) | Easy | Stack, Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| [230. Kth smallest element in a bst](https://leetcode.com/problems/kth-smallest-element-in-a-bst/) | [golang](/leetcode/kth-smallest-element-in-a-bst/kth-smallest-element-in-a-bst.go) | Medium | Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| [1769. Minimum number of operations to move all balls to each box](https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/) | [golang](/leetcode/minimum-number-of-operations-to-move-all-balls-to-each-box/minimum-number-of-operations-to-move-all-balls-to-each-box.go) | Medium | Array, String |
| [1137. N th tribonacci number](https://leetcode.com/problems/n-th-tribonacci-number/) | [golang](/leetcode/n-th-tribonacci-number/n-th-tribonacci-number.go) | Easy | Math, Dynamic Programming, Memoization |
| [139. Word break](https://leetcode.com/problems/word-break/) | [golang](/leetcode/word-break/word-break.go) | Medium | Hash Table, String, Dynamic Programming, Trie, Memoization |
| [509. Fibonacci number](https://leetcode.com/problems/fibonacci-number/) | [golang](/leetcode/fibonacci-number/fibonacci-number.go) | Easy | Math, Dynamic Programming, Recursion, Memoization |
| [1614. Maximum nesting depth of the parentheses](https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/) | [golang](/leetcode/maximum-nesting-depth-of-the-parentheses/maximum-nesting-depth-of-the-parentheses.go) | Easy | String, Stack |
| [99. Recover binary search tree](https://leetcode.com/problems/recover-binary-search-tree/) | [golang](/leetcode/recover-binary-search-tree/recover-binary-search-tree.go) | Medium | Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| [173. Binary search tree iterator](https://leetcode.com/problems/binary-search-tree-iterator/) | [golang](/leetcode/binary-search-tree-iterator/binary-search-tree-iterator.go) | Medium | Stack, Tree, Design, Binary Search Tree, Binary Tree, Iterator |
| [69. Sqrtx](https://leetcode.com/problems/sqrtx/) | [golang](/leetcode/sqrtx/sqrtx.go) | Easy | Math, Binary Search |
| [7. Reverse integer](https://leetcode.com/problems/reverse-integer/) | [golang](/leetcode/reverse-integer/reverse-integer.go) | Medium | Math |
| [18. 4sum](https://leetcode.com/problems/4sum/) | [golang](/leetcode/4sum/4sum.go) | Medium | Array, Two Pointers, Sorting |
| [695. Max area of island](https://leetcode.com/problems/max-area-of-island/) | [golang](/leetcode/max-area-of-island/max-area-of-island.go) | Medium | Array, Depth-First Search, Breadth-First Search, Union Find, Matrix |
| [153. Find minimum in rotated sorted array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) | [golang](/leetcode/find-minimum-in-rotated-sorted-array/find-minimum-in-rotated-sorted-array.go) | Medium | Array, Binary Search |
| [705. Design hashset](https://leetcode.com/problems/design-hashset/) | [golang](/leetcode/design-hashset/design-hashset.go) | Easy | Array, Hash Table, Linked List, Design, Hash Function |
| [706. Design hashmap](https://leetcode.com/problems/design-hashmap/) | [golang](/leetcode/design-hashmap/design-hashmap.go) | Easy | Array, Hash Table, Linked List, Design, Hash Function |
| [535. Encode and decode tinyurl](https://leetcode.com/problems/encode-and-decode-tinyurl/) | [golang](/leetcode/encode-and-decode-tinyurl/encode-and-decode-tinyurl.go) | Medium | Hash Table, String, Design, Hash Function |
| [187. Repeated dna sequences](https://leetcode.com/problems/repeated-dna-sequences/) | [golang](/leetcode/repeated-dna-sequences/repeated-dna-sequences.go) | Medium | Hash Table, String, Bit Manipulation, Sliding Window, Rolling Hash, Hash Function |
| [6. Zigzag conversion](https://leetcode.com/problems/zigzag-conversion/) | [golang](/leetcode/zigzag-conversion/zigzag-conversion.go) | Medium | String |
| [5. Longest palindromic substring](https://leetcode.com/problems/longest-palindromic-substring/) | [golang](/leetcode/longest-palindromic-substring/longest-palindromic-substring.go) | Medium | String, Dynamic Programming |
| [2181. Merge nodes in between zeros](https://leetcode.com/problems/merge-nodes-in-between-zeros/) | [golang](/leetcode/merge-nodes-in-between-zeros/merge-nodes-in-between-zeros.go) | Medium | Linked List, Simulation |
| [1396. Design underground system](https://leetcode.com/problems/design-underground-system/) | [golang](/leetcode/design-underground-system/design-underground-system.go) | Medium | Hash Table, String, Design |
| [23. Merge k sorted lists](https://leetcode.com/problems/merge-k-sorted-lists/) | [golang](/leetcode/merge-k-sorted-lists/merge-k-sorted-lists.go) | Hard | Linked List, Divide and Conquer, Heap (Priority Queue), Merge Sort |
| [16. 3sum closest](https://leetcode.com/problems/3sum-closest/) | [golang](/leetcode/3sum-closest/3sum-closest.go) | Medium | Array, Two Pointers, Sorting |
| [2248. Intersection of multiple arrays](https://leetcode.com/problems/intersection-of-multiple-arrays/) | [golang](/leetcode/intersection-of-multiple-arrays/intersection-of-multiple-arrays.go) | Easy | Array, Hash Table, Counting |
| [2249. Count lattice points inside a circle](https://leetcode.com/problems/count-lattice-points-inside-a-circle/) | [golang](/leetcode/count-lattice-points-inside-a-circle/count-lattice-points-inside-a-circle.go) | Medium | Array, Hash Table, Math, Geometry, Enumeration |
| [284. Peeking iterator](https://leetcode.com/problems/peeking-iterator/) | [golang](/leetcode/peeking-iterator/peeking-iterator.go) | Medium | Array, Design, Iterator |
| [19. Remove nth node from end of list](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | [golang](/leetcode/remove-nth-node-from-end-of-list/remove-nth-node-from-end-of-list.go) | Medium | Linked List, Two Pointers |
| [22. Generate parentheses](https://leetcode.com/problems/generate-parentheses/) | [golang](/leetcode/generate-parentheses/generate-parentheses.go) | Medium | String, Dynamic Programming, Backtracking |
| [24. Swap nodes in pairs](https://leetcode.com/problems/swap-nodes-in-pairs/) | [golang](/leetcode/swap-nodes-in-pairs/swap-nodes-in-pairs.go) | Medium | Linked List, Recursion |
| [1584. Min cost to connect all points](https://leetcode.com/problems/min-cost-to-connect-all-points/) | [golang](/leetcode/min-cost-to-connect-all-points/min-cost-to-connect-all-points.go) | Medium | Array, Union Find, Minimum Spanning Tree |
| [46. Permutations](https://leetcode.com/problems/permutations/) | [golang](/leetcode/permutations/permutations.go) | Medium | Array, Backtracking |
| [39. Combination sum](https://leetcode.com/problems/combination-sum/) | [golang](/leetcode/combination-sum/combination-sum.go) | Medium | Array, Backtracking |
| [34. Find first and last position of element in sorted array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | [golang](/leetcode/find-first-and-last-position-of-element-in-sorted-array/find-first-and-last-position-of-element-in-sorted-array.go) | Medium | Array, Binary Search |
| [1202. Smallest string with swaps](https://leetcode.com/problems/smallest-string-with-swaps/) | [golang](/leetcode/smallest-string-with-swaps/smallest-string-with-swaps.go) | Medium | Hash Table, String, Depth-First Search, Breadth-First Search, Union Find |
| [128. Longest consecutive sequence](https://leetcode.com/problems/longest-consecutive-sequence/) | [golang](/leetcode/longest-consecutive-sequence/longest-consecutive-sequence.go) | Medium | Array, Hash Table, Union Find |
| [547. Number of provinces](https://leetcode.com/problems/number-of-provinces/) | [golang](/leetcode/number-of-provinces/number-of-provinces.go) | Medium | Depth-First Search, Breadth-First Search, Union Find, Graph |
| [200. Number of islands](https://leetcode.com/problems/number-of-islands/) | [golang](/leetcode/number-of-islands/number-of-islands.go) | Medium | Array, Depth-First Search, Breadth-First Search, Union Find, Matrix |
| [104. Maximum depth of binary tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | [golang](/leetcode/maximum-depth-of-binary-tree/maximum-depth-of-binary-tree.go) | Easy | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [1631. Path with minimum effort](https://leetcode.com/problems/path-with-minimum-effort/) | [golang](/leetcode/path-with-minimum-effort/path-with-minimum-effort.go) | Medium | Array, Binary Search, Depth-First Search, Breadth-First Search, Union Find, Heap (Priority Queue), Matrix |
| [1512. Number of good pairs](https://leetcode.com/problems/number-of-good-pairs/) | [golang](/leetcode/number-of-good-pairs/number-of-good-pairs.go) | Easy | Array, Hash Table, Math, Counting |
| [1431. Kids with the greatest number of candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/) | [golang](/leetcode/kids-with-the-greatest-number-of-candies/kids-with-the-greatest-number-of-candies.go) | Easy | Array |
| [1281. Subtract the product and sum of digits of an integer](https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/) | [golang](/leetcode/subtract-the-product-and-sum-of-digits-of-an-integer/subtract-the-product-and-sum-of-digits-of-an-integer.go) | Easy | Math |
| [1155. Number of dice rolls with target sum](https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/) | [golang](/leetcode/number-of-dice-rolls-with-target-sum/number-of-dice-rolls-with-target-sum.go) | Medium | Dynamic Programming |
| [785. Is graph bipartite](https://leetcode.com/problems/is-graph-bipartite/) | [golang](/leetcode/is-graph-bipartite/is-graph-bipartite.go) | Medium | Depth-First Search, Breadth-First Search, Union Find, Graph |
| [1573. Number of ways to split a string](https://leetcode.com/problems/number-of-ways-to-split-a-string/) | [golang](/leetcode/number-of-ways-to-split-a-string/number-of-ways-to-split-a-string.go) | Medium | Math, String |
| [399. Evaluate division](https://leetcode.com/problems/evaluate-division/) | [golang](/leetcode/evaluate-division/evaluate-division.go) | Medium | Array, Depth-First Search, Breadth-First Search, Union Find, Graph, Shortest Path |
| [844. Backspace string compare](https://leetcode.com/problems/backspace-string-compare/) | [golang](/leetcode/backspace-string-compare/backspace-string-compare.go) | Easy | Two Pointers, String, Stack, Simulation |
| [2255. Count prefixes of a given string](https://leetcode.com/problems/count-prefixes-of-a-given-string/) | [golang](/leetcode/count-prefixes-of-a-given-string/count-prefixes-of-a-given-string.go) | Easy | Array, String |
| [2256. Minimum average difference](https://leetcode.com/problems/minimum-average-difference/) | [golang](/leetcode/minimum-average-difference/minimum-average-difference.go) | Medium | Array, Prefix Sum |
| [2259. Remove digit from number to maximize result](https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/) | [golang](/leetcode/remove-digit-from-number-to-maximize-result/remove-digit-from-number-to-maximize-result.go) | Easy | String, Greedy, Enumeration |
| [2260. Minimum consecutive cards to pick up](https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/) | [golang](/leetcode/minimum-consecutive-cards-to-pick-up/minimum-consecutive-cards-to-pick-up.go) | Medium | Array, Hash Table, Sliding Window |
| [48. Rotate image](https://leetcode.com/problems/rotate-image/) | [golang](/leetcode/rotate-image/rotate-image.go) | Medium | Array, Math, Matrix |
| [905. Sort array by parity](https://leetcode.com/problems/sort-array-by-parity/) | [golang](/leetcode/sort-array-by-parity/sort-array-by-parity.go) | Easy | Array, Two Pointers, Sorting |
| [581. Shortest unsorted continuous subarray](https://leetcode.com/problems/shortest-unsorted-continuous-subarray/) | [golang](/leetcode/shortest-unsorted-continuous-subarray/shortest-unsorted-continuous-subarray.go) | Medium | Array, Two Pointers, Stack, Greedy, Sorting, Monotonic Stack |
| [1679. Max number of k sum pairs](https://leetcode.com/problems/max-number-of-k-sum-pairs/) | [golang](/leetcode/max-number-of-k-sum-pairs/max-number-of-k-sum-pairs.go) | Medium | Array, Hash Table, Two Pointers, Sorting |
| [225. Implement stack using queues](https://leetcode.com/problems/implement-stack-using-queues/) | [golang](/leetcode/implement-stack-using-queues/implement-stack-using-queues.go) | Easy | Stack, Design, Queue |
| [27. Remove element](https://leetcode.com/problems/remove-element/) | [golang](/leetcode/remove-element/remove-element.go) | Easy | Array, Two Pointers |
| [1209. Remove all adjacent duplicates in string ii](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/) | [golang](/leetcode/remove-all-adjacent-duplicates-in-string-ii/remove-all-adjacent-duplicates-in-string-ii.go) | Medium | String, Stack |
| [456. 132 pattern](https://leetcode.com/problems/132-pattern/) | [golang](/leetcode/132-pattern/132-pattern.go) | Medium | Array, Binary Search, Stack, Monotonic Stack, Ordered Set |
| [409. Longest palindrome](https://leetcode.com/problems/longest-palindrome/) | [golang](/leetcode/longest-palindrome/longest-palindrome.go) | Easy | Hash Table, String, Greedy |
| [561. Array partition i](https://leetcode.com/problems/array-partition-i/) | [golang](/leetcode/array-partition-i/array-partition-i.go) | Easy | Array, Greedy, Sorting, Counting Sort |
| [341. Flatten nested list iterator](https://leetcode.com/problems/flatten-nested-list-iterator/) | [golang](/leetcode/flatten-nested-list-iterator/flatten-nested-list-iterator.go) | Medium | Stack, Tree, Depth-First Search, Design, Queue, Iterator |
| [216. Combination sum iii](https://leetcode.com/problems/combination-sum-iii/) | [golang](/leetcode/combination-sum-iii/combination-sum-iii.go) | Medium | Array, Backtracking |
| [47. Permutations ii](https://leetcode.com/problems/permutations-ii/) | [golang](/leetcode/permutations-ii/permutations-ii.go) | Medium | Array, Backtracking |
| [1641. Count sorted vowel strings](https://leetcode.com/problems/count-sorted-vowel-strings/) | [golang](/leetcode/count-sorted-vowel-strings/count-sorted-vowel-strings.go) | Medium | Dynamic Programming |
| [117. Populating next right pointers in each node ii](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/) | [golang](/leetcode/populating-next-right-pointers-in-each-node-ii/populating-next-right-pointers-in-each-node-ii.go) | Medium | Linked List, Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [1302. Deepest leaves sum](https://leetcode.com/problems/deepest-leaves-sum/) | [golang](/leetcode/deepest-leaves-sum/deepest-leaves-sum.go) | Medium | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [191. Number of 1 bits](https://leetcode.com/problems/number-of-1-bits/) | [golang](/leetcode/number-of-1-bits/number-of-1-bits.go) | Easy | Bit Manipulation |
| [1342. Number of steps to reduce a number to zero](https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/) | [golang](/leetcode/number-of-steps-to-reduce-a-number-to-zero/number-of-steps-to-reduce-a-number-to-zero.go) | Easy | Math, Bit Manipulation |
| [268. Missing number](https://leetcode.com/problems/missing-number/) | [golang](/leetcode/missing-number/missing-number.go) | Easy | Array, Hash Table, Math, Bit Manipulation, Sorting |
| [318. Maximum product of word lengths](https://leetcode.com/problems/maximum-product-of-word-lengths/) | [golang](/leetcode/maximum-product-of-word-lengths/maximum-product-of-word-lengths.go) | Medium | Array, String, Bit Manipulation |
| [29. Divide two integers](https://leetcode.com/problems/divide-two-integers/) | [golang](/leetcode/divide-two-integers/divide-two-integers.go) | Medium | Math, Bit Manipulation |
| [1461. Check if a string contains all binary codes of size k](https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/) | [golang](/leetcode/check-if-a-string-contains-all-binary-codes-of-size-k/check-if-a-string-contains-all-binary-codes-of-size-k.go) | Medium | Hash Table, String, Bit Manipulation, Rolling Hash, Hash Function |
| [867. Transpose matrix](https://leetcode.com/problems/transpose-matrix/) | [golang](/leetcode/transpose-matrix/transpose-matrix.go) | Easy | Array, Matrix, Simulation |
| [304. Range sum query 2d immutable](https://leetcode.com/problems/range-sum-query-2d-immutable/) | [golang](/leetcode/range-sum-query-2d-immutable/range-sum-query-2d-immutable.go) | Medium | Array, Design, Matrix, Prefix Sum |
| [160. Intersection of two linked lists](https://leetcode.com/problems/intersection-of-two-linked-lists/) | [golang](/leetcode/intersection-of-two-linked-lists/intersection-of-two-linked-lists.go) | Easy | Hash Table, Linked List, Two Pointers |
| [88. Merge sorted array](https://leetcode.com/problems/merge-sorted-array/) | [golang](/leetcode/merge-sorted-array/merge-sorted-array.go) | Easy | Array, Two Pointers, Sorting |
| [1332. Remove palindromic subsequences](https://leetcode.com/problems/remove-palindromic-subsequences/) | [golang](/leetcode/remove-palindromic-subsequences/remove-palindromic-subsequences.go) | Easy | Two Pointers, String |
| [51. N queens](https://leetcode.com/problems/n-queens/) | [golang](/leetcode/n-queens/n-queens.go) | Hard | Array, Backtracking |
| [52. N queens ii](https://leetcode.com/problems/n-queens-ii/) | [golang](/leetcode/n-queens-ii/n-queens-ii.go) | Hard | Backtracking |
| [167. Two sum ii input array is sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | [golang](/leetcode/two-sum-ii-input-array-is-sorted/two-sum-ii-input-array-is-sorted.go) | Medium | Array, Two Pointers, Binary Search |
| [1695. Maximum erasure value](https://leetcode.com/problems/maximum-erasure-value/) | [golang](/leetcode/maximum-erasure-value/maximum-erasure-value.go) | Medium | Array, Hash Table, Sliding Window |
| [120. Triangle](https://leetcode.com/problems/triangle/) | [golang](/leetcode/triangle/triangle.go) | Medium | Array, Dynamic Programming |
| [583. Delete operation for two strings](https://leetcode.com/problems/delete-operation-for-two-strings/) | [golang](/leetcode/delete-operation-for-two-strings/delete-operation-for-two-strings.go) | Medium | String, Dynamic Programming |
| [1048. Longest string chain](https://leetcode.com/problems/longest-string-chain/) | [golang](/leetcode/longest-string-chain/longest-string-chain.go) | Medium | Array, Hash Table, Two Pointers, String, Dynamic Programming |
| [242. Valid anagram](https://leetcode.com/problems/valid-anagram/) | [golang](/leetcode/valid-anagram/valid-anagram.go) | Easy | Hash Table, String, Sorting |
| [62. Unique paths](https://leetcode.com/problems/unique-paths/) | [golang](/leetcode/unique-paths/unique-paths.go) | Medium | Math, Dynamic Programming, Combinatorics |
| [378. Kth smallest element in a sorted matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/) | [golang](/leetcode/kth-smallest-element-in-a-sorted-matrix/kth-smallest-element-in-a-sorted-matrix.go) | Medium | Array, Binary Search, Sorting, Heap (Priority Queue), Matrix |
| [729. My calendar i](https://leetcode.com/problems/my-calendar-i/) | [golang](/leetcode/my-calendar-i/my-calendar-i.go) | Medium | Binary Search, Design, Segment Tree, Ordered Set |
| [858. Mirror reflection](https://leetcode.com/problems/mirror-reflection/) | [golang](/leetcode/mirror-reflection/mirror-reflection.go) | Medium | Math, Geometry |
| [377. Combination sum iv](https://leetcode.com/problems/combination-sum-iv/) | [golang](/leetcode/combination-sum-iv/combination-sum-iv.go) | Medium | Array, Dynamic Programming |
| [458. Poor pigs](https://leetcode.com/problems/poor-pigs/) | [golang](/leetcode/poor-pigs/poor-pigs.go) | Hard | Math, Dynamic Programming, Combinatorics |
| [300. Longest increasing subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) | [golang](/leetcode/longest-increasing-subsequence/longest-increasing-subsequence.go) | Medium | Array, Binary Search, Dynamic Programming |
| [1220. Count vowels permutation](https://leetcode.com/problems/count-vowels-permutation/) | [golang](/leetcode/count-vowels-permutation/count-vowels-permutation.go) | Hard | Dynamic Programming |
| [823. Binary trees with factors](https://leetcode.com/problems/binary-trees-with-factors/) | [golang](/leetcode/binary-trees-with-factors/binary-trees-with-factors.go) | Medium | Array, Hash Table, Dynamic Programming |
| [108. Convert sorted array to binary search tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/) | [golang](/leetcode/convert-sorted-array-to-binary-search-tree/convert-sorted-array-to-binary-search-tree.go) | Easy | Array, Divide and Conquer, Tree, Binary Search Tree, Binary Tree |
| [98. Validate binary search tree](https://leetcode.com/problems/validate-binary-search-tree/) | [golang](/leetcode/validate-binary-search-tree/validate-binary-search-tree.go) | Medium | Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| [235. Lowest common ancestor of a binary search tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | [golang](/leetcode/lowest-common-ancestor-of-a-binary-search-tree/lowest-common-ancestor-of-a-binary-search-tree.go) | Easy | Tree, Depth-First Search, Binary Search Tree, Binary Tree |
| [387. First unique character in a string](https://leetcode.com/problems/first-unique-character-in-a-string/) | [golang](/leetcode/first-unique-character-in-a-string/first-unique-character-in-a-string.go) | Easy | Hash Table, String, Queue, Counting |
| [804. Unique morse code words](https://leetcode.com/problems/unique-morse-code-words/) | [golang](/leetcode/unique-morse-code-words/unique-morse-code-words.go) | Easy | Array, Hash Table, String |
| [1338. Reduce array size to the half](https://leetcode.com/problems/reduce-array-size-to-the-half/) | [golang](/leetcode/reduce-array-size-to-the-half/reduce-array-size-to-the-half.go) | Medium | Array, Hash Table, Greedy, Sorting, Heap (Priority Queue) |
| [659. Split array into consecutive subsequences](https://leetcode.com/problems/split-array-into-consecutive-subsequences/) | [golang](/leetcode/split-array-into-consecutive-subsequences/split-array-into-consecutive-subsequences.go) | Medium | Array, Hash Table, Greedy, Heap (Priority Queue) |
| [342. Power of four](https://leetcode.com/problems/power-of-four/) | [golang](/leetcode/power-of-four/power-of-four.go) | Easy | Math, Bit Manipulation, Recursion |
| [234. Palindrome linked list](https://leetcode.com/problems/palindrome-linked-list/) | [golang](/leetcode/palindrome-linked-list/palindrome-linked-list.go) | Easy | Linked List, Two Pointers, Stack, Recursion |
| [326. Power of three](https://leetcode.com/problems/power-of-three/) | [golang](/leetcode/power-of-three/power-of-three.go) | Easy | Math, Recursion |
| [383. Ransom note](https://leetcode.com/problems/ransom-note/) | [golang](/leetcode/ransom-note/ransom-note.go) | Easy | Hash Table, String, Counting |
| [869. Reordered power of 2](https://leetcode.com/problems/reordered-power-of-2/) | [golang](/leetcode/reordered-power-of-2/reordered-power-of-2.go) | Medium | Math, Sorting, Counting, Enumeration |
| [2379. Minimum recolors to get k consecutive black blocks](https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/) | [golang](/leetcode/minimum-recolors-to-get-k-consecutive-black-blocks/minimum-recolors-to-get-k-consecutive-black-blocks.go) | Easy | String, Sliding Window |
| [2380. Time needed to rearrange a binary string](https://leetcode.com/problems/time-needed-to-rearrange-a-binary-string/) | [golang](/leetcode/time-needed-to-rearrange-a-binary-string/time-needed-to-rearrange-a-binary-string.go) | Medium | String, Dynamic Programming, Simulation |
| [2381. Shifting letters ii](https://leetcode.com/problems/shifting-letters-ii/) | [golang](/leetcode/shifting-letters-ii/shifting-letters-ii.go) | Medium | Array, String, Prefix Sum |
| [126. Word ladder ii](https://leetcode.com/problems/word-ladder-ii/) | [golang](/leetcode/word-ladder-ii/word-ladder-ii.go) | Hard | Hash Table, String, Backtracking, Breadth-First Search |
| [363. Max sum of rectangle no larger than k](https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/) | [golang](/leetcode/max-sum-of-rectangle-no-larger-than-k/max-sum-of-rectangle-no-larger-than-k.go) | Hard | Array, Binary Search, Matrix, Prefix Sum, Ordered Set |
| [1578. Minimum time to make rope colorful](https://leetcode.com/problems/minimum-time-to-make-rope-colorful/) | [golang](/leetcode/minimum-time-to-make-rope-colorful/minimum-time-to-make-rope-colorful.go) | Medium | Array, String, Dynamic Programming, Greedy |
| [91. Decode ways](https://leetcode.com/problems/decode-ways/) | [golang](/leetcode/decode-ways/decode-ways.go) | Medium | String, Dynamic Programming |
| [112. Path sum](https://leetcode.com/problems/path-sum/) | [golang](/leetcode/path-sum/path-sum.go) | Easy | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [623. Add one row to tree](https://leetcode.com/problems/add-one-row-to-tree/) | [golang](/leetcode/add-one-row-to-tree/add-one-row-to-tree.go) | Medium | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [981. Time based key value store](https://leetcode.com/problems/time-based-key-value-store/) | [golang](/leetcode/time-based-key-value-store/time-based-key-value-store.go) | Medium | Hash Table, String, Binary Search, Design |
| [732. My calendar iii](https://leetcode.com/problems/my-calendar-iii/) | [golang](/leetcode/my-calendar-iii/my-calendar-iii.go) | Hard | Binary Search, Design, Segment Tree, Ordered Set |
| [653. Two sum iv input is a bst](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/) | [golang](/leetcode/two-sum-iv-input-is-a-bst/two-sum-iv-input-is-a-bst.go) | Easy | Hash Table, Two Pointers, Tree, Depth-First Search, Breadth-First Search, Binary Search Tree, Binary Tree |
| [1328. Break a palindrome](https://leetcode.com/problems/break-a-palindrome/) | [golang](/leetcode/break-a-palindrome/break-a-palindrome.go) | Medium | String, Greedy |
| [334. Increasing triplet subsequence](https://leetcode.com/problems/increasing-triplet-subsequence/) | [golang](/leetcode/increasing-triplet-subsequence/increasing-triplet-subsequence.go) | Medium | Array, Greedy |
| [121. Best time to buy and sell stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [golang](/leetcode/best-time-to-buy-and-sell-stock/best-time-to-buy-and-sell-stock.go) | Easy | Array, Dynamic Programming |
| [125. Valid palindrome](https://leetcode.com/problems/valid-palindrome/) | [golang](/leetcode/valid-palindrome/valid-palindrome.go) | Easy | Two Pointers, String |
| [226. Invert binary tree](https://leetcode.com/problems/invert-binary-tree/) | [golang](/leetcode/invert-binary-tree/invert-binary-tree.go) | Easy | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [704. Binary search](https://leetcode.com/problems/binary-search/) | [golang](/leetcode/binary-search/binary-search.go) | Easy | Array, Binary Search |
| [733. Flood fill](https://leetcode.com/problems/flood-fill/) | [golang](/leetcode/flood-fill/flood-fill.go) | Easy | Array, Depth-First Search, Breadth-First Search, Matrix |
| [976. Largest perimeter triangle](https://leetcode.com/problems/largest-perimeter-triangle/) | [golang](/leetcode/largest-perimeter-triangle/largest-perimeter-triangle.go) | Easy | Array, Math, Greedy, Sorting |
| [237. Delete node in a linked list](https://leetcode.com/problems/delete-node-in-a-linked-list/) | [golang](/leetcode/delete-node-in-a-linked-list/delete-node-in-a-linked-list.go) | Medium | Linked List |
| [2095. Delete the middle node of a linked list](https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/) | [golang](/leetcode/delete-the-middle-node-of-a-linked-list/delete-the-middle-node-of-a-linked-list.go) | Medium | Linked List, Two Pointers |
| [1531. String compression ii](https://leetcode.com/problems/string-compression-ii/) | [golang](/leetcode/string-compression-ii/string-compression-ii.go) | Hard | String, Dynamic Programming |
| [110. Balanced binary tree](https://leetcode.com/problems/balanced-binary-tree/) | [golang](/leetcode/balanced-binary-tree/balanced-binary-tree.go) | Easy | Tree, Depth-First Search, Binary Tree |
| [232. Implement queue using stacks](https://leetcode.com/problems/implement-queue-using-stacks/) | [golang](/leetcode/implement-queue-using-stacks/implement-queue-using-stacks.go) | Easy | Stack, Design, Queue |
| [278. First bad version](https://leetcode.com/problems/first-bad-version/) | [golang](/leetcode/first-bad-version/first-bad-version.go) | Easy | Binary Search, Interactive |
| [206. Reverse linked list](https://leetcode.com/problems/reverse-linked-list/) | [golang](/leetcode/reverse-linked-list/reverse-linked-list.go) | Easy | Linked List, Recursion |
| [169. Majority element](https://leetcode.com/problems/majority-element/) | [golang](/leetcode/majority-element/majority-element.go) | Easy | Array, Hash Table, Divide and Conquer, Sorting, Counting |
| [67. Add binary](https://leetcode.com/problems/add-binary/) | [golang](/leetcode/add-binary/add-binary.go) | Easy | Math, String, Bit Manipulation, Simulation |
| [1832. Check if the sentence is pangram](https://leetcode.com/problems/check-if-the-sentence-is-pangram/) | [golang](/leetcode/check-if-the-sentence-is-pangram/check-if-the-sentence-is-pangram.go) | Easy | Hash Table, String |
| [543. Diameter of binary tree](https://leetcode.com/problems/diameter-of-binary-tree/) | [golang](/leetcode/diameter-of-binary-tree/diameter-of-binary-tree.go) | Easy | Tree, Depth-First Search, Binary Tree |
| [876. Middle of the linked list](https://leetcode.com/problems/middle-of-the-linked-list/) | [golang](/leetcode/middle-of-the-linked-list/middle-of-the-linked-list.go) | Easy | Linked List, Two Pointers |
| [217. Contains duplicate](https://leetcode.com/problems/contains-duplicate/) | [golang](/leetcode/contains-duplicate/contains-duplicate.go) | Easy | Array, Hash Table, Sorting |
| [38. Count and say](https://leetcode.com/problems/count-and-say/) | [golang](/leetcode/count-and-say/count-and-say.go) | Medium | String |
| [542. 01 matrix](https://leetcode.com/problems/01-matrix/) | [golang](/leetcode/01-matrix/01-matrix.go) | Medium | Array, Dynamic Programming, Breadth-First Search, Matrix |
| [57. Insert interval](https://leetcode.com/problems/insert-interval/) | [golang](/leetcode/insert-interval/insert-interval.go) | Medium | Array |
| [692. Top k frequent words](https://leetcode.com/problems/top-k-frequent-words/) | [golang](/leetcode/top-k-frequent-words/top-k-frequent-words.go) | Medium | Hash Table, String, Trie, Sorting, Heap (Priority Queue), Bucket Sort, Counting |
| [12. Integer to roman](https://leetcode.com/problems/integer-to-roman/) | [golang](/leetcode/integer-to-roman/integer-to-roman.go) | Medium | Hash Table, Math, String |
| [973. K closest points to origin](https://leetcode.com/problems/k-closest-points-to-origin/) | [golang](/leetcode/k-closest-points-to-origin/k-closest-points-to-origin.go) | Medium | Array, Math, Divide and Conquer, Geometry, Sorting, Heap (Priority Queue), Quickselect |
| [433. Minimum genetic mutation](https://leetcode.com/problems/minimum-genetic-mutation/) | [golang](/leetcode/minimum-genetic-mutation/minimum-genetic-mutation.go) | Medium | Hash Table, String, Breadth-First Search |
| [2131. Longest palindrome by concatenating two letter words](https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/) | [golang](/leetcode/longest-palindrome-by-concatenating-two-letter-words/longest-palindrome-by-concatenating-two-letter-words.go) | Medium | Array, Hash Table, String, Greedy, Counting |
| [345. Reverse vowels of a string](https://leetcode.com/problems/reverse-vowels-of-a-string/) | [golang](/leetcode/reverse-vowels-of-a-string/reverse-vowels-of-a-string.go) | Easy | Two Pointers, String |
| [212. Word search ii](https://leetcode.com/problems/word-search-ii/) | [golang](/leetcode/word-search-ii/word-search-ii.go) | Hard | Array, String, Backtracking, Trie, Matrix |
| [899. Orderly queue](https://leetcode.com/problems/orderly-queue/) | [golang](/leetcode/orderly-queue/orderly-queue.go) | Hard | Math, String, Sorting |
| [208. Implement trie prefix tree](https://leetcode.com/problems/implement-trie-prefix-tree/) | [golang](/leetcode/implement-trie-prefix-tree/implement-trie-prefix-tree.go) | Medium | Hash Table, String, Design, Trie |
| [211. Design add and search words data structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/) | [golang](/leetcode/design-add-and-search-words-data-structure/design-add-and-search-words-data-structure.go) | Medium | String, Depth-First Search, Design, Trie |
| [1323. Maximum 69 number](https://leetcode.com/problems/maximum-69-number/) | [golang](/leetcode/maximum-69-number/maximum-69-number.go) | Easy | Math, Greedy |
| [102. Binary tree level order traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) | [golang](/leetcode/binary-tree-level-order-traversal/binary-tree-level-order-traversal.go) | Medium | Tree, Breadth-First Search, Binary Tree |
| [133. Clone graph](https://leetcode.com/problems/clone-graph/) | [golang](/leetcode/clone-graph/clone-graph.go) | Medium | Hash Table, Depth-First Search, Breadth-First Search, Graph |
| [150. Evaluate reverse polish notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/) | [golang](/leetcode/evaluate-reverse-polish-notation/evaluate-reverse-polish-notation.go) | Medium | Array, Math, Stack |
| [207. Course schedule](https://leetcode.com/problems/course-schedule/) | [golang](/leetcode/course-schedule/course-schedule.go) | Medium | Depth-First Search, Breadth-First Search, Graph, Topological Sort |
| [1544. Make the string great](https://leetcode.com/problems/make-the-string-great/) | [golang](/leetcode/make-the-string-great/make-the-string-great.go) | Easy | String, Stack |
| [322. Coin change](https://leetcode.com/problems/coin-change/) | [golang](/leetcode/coin-change/coin-change.go) | Medium | Array, Dynamic Programming, Breadth-First Search |
| [238. Product of array except self](https://leetcode.com/problems/product-of-array-except-self/) | [golang](/leetcode/product-of-array-except-self/product-of-array-except-self.go) | Medium | Array, Prefix Sum |
| [155. Min stack](https://leetcode.com/problems/min-stack/) | [golang](/leetcode/min-stack/min-stack.go) | Medium | Stack, Design |
| [994. Rotting oranges](https://leetcode.com/problems/rotting-oranges/) | [golang](/leetcode/rotting-oranges/rotting-oranges.go) | Medium | Array, Breadth-First Search, Matrix |
| [901. Online stock span](https://leetcode.com/problems/online-stock-span/) | [golang](/leetcode/online-stock-span/online-stock-span.go) | Medium | Stack, Design, Monotonic Stack, Data Stream |
| [1047. Remove all adjacent duplicates in string](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/) | [golang](/leetcode/remove-all-adjacent-duplicates-in-string/remove-all-adjacent-duplicates-in-string.go) | Easy | String, Stack |
| [26. Remove duplicates from sorted array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | [golang](/leetcode/remove-duplicates-from-sorted-array/remove-duplicates-from-sorted-array.go) | Easy | Array, Two Pointers |
| [151. Reverse words in a string](https://leetcode.com/problems/reverse-words-in-a-string/) | [golang](/leetcode/reverse-words-in-a-string/reverse-words-in-a-string.go) | Medium | Two Pointers, String |
| [947. Most stones removed with same row or column](https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/) | [golang](/leetcode/most-stones-removed-with-same-row-or-column/most-stones-removed-with-same-row-or-column.go) | Medium | Depth-First Search, Union Find, Graph |
| [222. Count complete tree nodes](https://leetcode.com/problems/count-complete-tree-nodes/) | [golang](/leetcode/count-complete-tree-nodes/count-complete-tree-nodes.go) | Medium | Binary Search, Tree, Depth-First Search, Binary Tree |
| [56. Merge intervals](https://leetcode.com/problems/merge-intervals/) | [golang](/leetcode/merge-intervals/merge-intervals.go) | Medium | Array, Sorting |
| [374. Guess number higher or lower](https://leetcode.com/problems/guess-number-higher-or-lower/) | [golang](/leetcode/guess-number-higher-or-lower/guess-number-higher-or-lower.go) | Easy | Binary Search, Interactive |
| [236. Lowest common ancestor of a binary tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) | [golang](/leetcode/lowest-common-ancestor-of-a-binary-tree/lowest-common-ancestor-of-a-binary-tree.go) | Medium | Tree, Depth-First Search, Binary Tree |
| [721. Accounts merge](https://leetcode.com/problems/accounts-merge/) | [golang](/leetcode/accounts-merge/accounts-merge.go) | Medium | Array, String, Depth-First Search, Breadth-First Search, Union Find |
| [223. Rectangle area](https://leetcode.com/problems/rectangle-area/) | [golang](/leetcode/rectangle-area/rectangle-area.go) | Medium | Math, Geometry |
| [416. Partition equal subset sum](https://leetcode.com/problems/partition-equal-subset-sum/) | [golang](/leetcode/partition-equal-subset-sum/partition-equal-subset-sum.go) | Medium | Array, Dynamic Programming |
| [263. Ugly number](https://leetcode.com/problems/ugly-number/) | [golang](/leetcode/ugly-number/ugly-number.go) | Easy | Math |
| [2395. Find subarrays with equal sum](https://leetcode.com/problems/find-subarrays-with-equal-sum/) | [golang](/leetcode/find-subarrays-with-equal-sum/find-subarrays-with-equal-sum.go) | Easy | Array, Hash Table |
| [224. Basic calculator](https://leetcode.com/problems/basic-calculator/) | [golang](/leetcode/basic-calculator/basic-calculator.go) | Hard | Math, String, Stack, Recursion |
| [279. Perfect squares](https://leetcode.com/problems/perfect-squares/) | [golang](/leetcode/perfect-squares/perfect-squares.go) | Medium | Math, Dynamic Programming, Breadth-First Search |
| [2475. Number of unequal triplets in array](https://leetcode.com/problems/number-of-unequal-triplets-in-array/) | [golang](/leetcode/number-of-unequal-triplets-in-array/number-of-unequal-triplets-in-array.go) | Easy | Array, Hash Table |
| [2476. Closest nodes queries in a binary search tree](https://leetcode.com/problems/closest-nodes-queries-in-a-binary-search-tree/) | [golang](/leetcode/closest-nodes-queries-in-a-binary-search-tree/closest-nodes-queries-in-a-binary-search-tree.go) | Medium | Array, Binary Search, Tree, Depth-First Search, Binary Tree |
| [79. Word search](https://leetcode.com/problems/word-search/) | [golang](/leetcode/word-search/word-search.go) | Medium | Array, Backtracking, Matrix |
| [907. Sum of subarray minimums](https://leetcode.com/problems/sum-of-subarray-minimums/) | [golang](/leetcode/sum-of-subarray-minimums/sum-of-subarray-minimums.go) | Medium | Array, Dynamic Programming, Stack, Monotonic Stack |
| [2485. Find the pivot integer](https://leetcode.com/problems/find-the-pivot-integer/) | [golang](/leetcode/find-the-pivot-integer/find-the-pivot-integer.go) | Easy |  |
| [2486. Append characters to string to make subsequence](https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/) | [golang](/leetcode/append-characters-to-string-to-make-subsequence/append-characters-to-string-to-make-subsequence.go) | Medium |  |
| [2487. Remove nodes from linked list](https://leetcode.com/problems/remove-nodes-from-linked-list/) | [golang](/leetcode/remove-nodes-from-linked-list/remove-nodes-from-linked-list.go) | Medium |  |
| [2225. Find players with zero or one losses](https://leetcode.com/problems/find-players-with-zero-or-one-losses/) | [golang](/leetcode/find-players-with-zero-or-one-losses/find-players-with-zero-or-one-losses.go) | Medium | Array, Hash Table, Sorting, Counting |
| [587. Erect the fence](https://leetcode.com/problems/erect-the-fence/) | [golang](/leetcode/erect-the-fence/erect-the-fence.go) | Hard | Array, Math, Geometry |
| [380. Insert delete getrandom o1](https://leetcode.com/problems/insert-delete-getrandom-o1/) | [golang](/leetcode/insert-delete-getrandom-o1/insert-delete-getrandom-o1.go) | Medium | Array, Hash Table, Math, Design, Randomized |
| [1207. Unique number of occurrences](https://leetcode.com/problems/unique-number-of-occurrences/) | [golang](/leetcode/unique-number-of-occurrences/unique-number-of-occurrences.go) | Easy | Array, Hash Table |
| [1704. Determine if string halves are alike](https://leetcode.com/problems/determine-if-string-halves-are-alike/) | [golang](/leetcode/determine-if-string-halves-are-alike/determine-if-string-halves-are-alike.go) | Easy | String, Counting |
| [1657. Determine if two strings are close](https://leetcode.com/problems/determine-if-two-strings-are-close/) | [golang](/leetcode/determine-if-two-strings-are-close/determine-if-two-strings-are-close.go) | Medium | Hash Table, String, Sorting |
| [451. Sort characters by frequency](https://leetcode.com/problems/sort-characters-by-frequency/) | [golang](/leetcode/sort-characters-by-frequency/sort-characters-by-frequency.go) | Medium | Hash Table, String, Sorting, Heap (Priority Queue), Bucket Sort, Counting |
| [1752. Check if array is sorted and rotated](https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/) |  | Easy | Array |
| [35. Search insert position](https://leetcode.com/problems/search-insert-position/) | [golang](/leetcode/search-insert-position/search-insert-position.go) | Easy | Array, Binary Search |
| [58. Length of last word](https://leetcode.com/problems/length-of-last-word/) | [golang](/leetcode/length-of-last-word/length-of-last-word.go) | Easy | String |
| [66. Plus one](https://leetcode.com/problems/plus-one/) | [golang](/leetcode/plus-one/plus-one.go) | Easy | Array, Math |
| [83. Remove duplicates from sorted list](https://leetcode.com/problems/remove-duplicates-from-sorted-list/) | [golang](/leetcode/remove-duplicates-from-sorted-list/remove-duplicates-from-sorted-list.go) | Easy | Linked List |
| [94. Binary tree inorder traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | [golang](/leetcode/binary-tree-inorder-traversal/binary-tree-inorder-traversal.go) | Easy | Stack, Tree, Depth-First Search, Binary Tree |
| [100. Same tree](https://leetcode.com/problems/same-tree/) | [golang](/leetcode/same-tree/same-tree.go) | Easy | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [101. Symmetric tree](https://leetcode.com/problems/symmetric-tree/) | [golang](/leetcode/symmetric-tree/symmetric-tree.go) | Easy | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [111. Minimum depth of binary tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/) | [golang](/leetcode/minimum-depth-of-binary-tree/minimum-depth-of-binary-tree.go) | Easy | Tree, Depth-First Search, Breadth-First Search, Binary Tree |
| [119. Pascals triangle ii](https://leetcode.com/problems/pascals-triangle-ii/) | [golang](/leetcode/pascals-triangle-ii/pascals-triangle-ii.go) | Easy | Array, Dynamic Programming |
| [136. Single number](https://leetcode.com/problems/single-number/) | [golang](/leetcode/single-number/single-number.go) | Easy | Array, Bit Manipulation |
| [144. Binary tree preorder traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/) | [golang](/leetcode/binary-tree-preorder-traversal/binary-tree-preorder-traversal.go) | Easy | Stack, Tree, Depth-First Search, Binary Tree |
| [145. Binary tree postorder traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/) | [golang](/leetcode/binary-tree-postorder-traversal/binary-tree-postorder-traversal.go) | Easy | Stack, Tree, Depth-First Search, Binary Tree |
| [168. Excel sheet column title](https://leetcode.com/problems/excel-sheet-column-title/) | [golang](/leetcode/excel-sheet-column-title/excel-sheet-column-title.go) | Easy | Math, String |
| [171. Excel sheet column number](https://leetcode.com/problems/excel-sheet-column-number/) | [golang](/leetcode/excel-sheet-column-number/excel-sheet-column-number.go) | Easy | Math, String |
| [190. Reverse bits](https://leetcode.com/problems/reverse-bits/) | [golang](/leetcode/reverse-bits/reverse-bits.go) | Easy | Divide and Conquer, Bit Manipulation |
| [202. Happy number](https://leetcode.com/problems/happy-number/) | [golang](/leetcode/happy-number/happy-number.go) | Easy | Hash Table, Math, Two Pointers |
| [203. Remove linked list elements](https://leetcode.com/problems/remove-linked-list-elements/) | [golang](/leetcode/remove-linked-list-elements/remove-linked-list-elements.go) | Easy | Linked List, Recursion |
| [205. Isomorphic strings](https://leetcode.com/problems/isomorphic-strings/) | [golang](/leetcode/isomorphic-strings/isomorphic-strings.go) | Easy | Hash Table, String |
| [228. Summary ranges](https://leetcode.com/problems/summary-ranges/) | [golang](/leetcode/summary-ranges/summary-ranges.go) | Easy | Array |
| [258. Add digits](https://leetcode.com/problems/add-digits/) | [golang](/leetcode/add-digits/add-digits.go) | Easy | Math, Simulation, Number Theory |
| [328. Odd even linked list](https://leetcode.com/problems/odd-even-linked-list/) | [golang](/leetcode/odd-even-linked-list/odd-even-linked-list.go) | Medium | Linked List |


## codeforces

Problems solved in total: 69

|Solutions by difficulty|Solutions by tags|
|-|-|
|![by_difficulty](https://github.com/dimaglushkov/solutions/blob/master/codeforces/.by_difficulty.svg)|![by_tags](https://github.com/dimaglushkov/solutions/blob/master/codeforces/.by_tags.svg)|

| Problem | Solution | Difficulty | Tags |
|-|-|-|-|
| [112A. Petya and Strings](https://codeforces.com/contest/112/problem/A) | [python](/codeforces/112A/112A.py) | 800 | implementation, strings |
| [158A. Next Round](https://codeforces.com/contest/158/problem/A) | [python](/codeforces/158A/158A.py) | 800 | *special problem, implementation |
| [231A. Team](https://codeforces.com/contest/231/problem/A) | [python](/codeforces/231A/231A.py) | 800 | bruteforce, greedy |
| [236A. Boy or Girl](https://codeforces.com/contest/236/problem/A) | [python](/codeforces/236A/236A.py) | 800 | bruteforce, implementation, strings |
| [263A. Beautiful Matrix](https://codeforces.com/contest/263/problem/A) | [python](/codeforces/263A/263A.py) | 800 | implementation |
| [266A. Stones on the Table](https://codeforces.com/contest/266/problem/A) | [python](/codeforces/266A/266A.py) | 800 | implementation |
| [281A. Word Capitalization](https://codeforces.com/contest/281/problem/A) | [python](/codeforces/281A/281A.py) | 800 | implementation, strings |
| [282A. Bit++](https://codeforces.com/contest/282/problem/A) | [python](/codeforces/282A/282A.py) | 800 | implementation |
| [339A. Helpful Maths](https://codeforces.com/contest/339/problem/A) | [python](/codeforces/339A/339A.py) | 800 | greedy, implementation, sortings, strings |
| [4A. Watermelon](https://codeforces.com/contest/4/problem/A) | [python](/codeforces/4A/4A.py) | 800 | bruteforce, math |
| [50A. Domino piling](https://codeforces.com/contest/50/problem/A) | [python](/codeforces/50A/50A.py) | 800 | greedy, math |
| [546A. Soldier and Bananas](https://codeforces.com/contest/546/problem/A) | [python](/codeforces/546A/546A.py) | 800 | bruteforce, implementation, math |
| [71A. Way Too Long Words](https://codeforces.com/contest/71/problem/A) | [python](/codeforces/71A/71A.py) | 800 | strings |
| [133A. HQ9+](https://codeforces.com/contest/133/problem/A) | [python](/codeforces/133A/133A.py) | 900 | implementation |
| [160A. Twins](https://codeforces.com/contest/160/problem/A) | [python](/codeforces/160A/160A.py) | 900 | greedy, sortings |
| [318A. Even Odds](https://codeforces.com/contest/318/problem/A) | [python](/codeforces/318A/318A.py) | 900 | math |
| [580A. Kefa and First Steps](https://codeforces.com/contest/580/problem/A) | [python](/codeforces/580A/580A.py) | 900 | bruteforce, dp, implementation |
| [96A. Football](https://codeforces.com/contest/96/problem/A) | [python](/codeforces/96A/96A.py) | 900 | implementation, strings |
| [118A. String Task](https://codeforces.com/contest/118/problem/A) | [python](/codeforces/118A/118A.py) | 1000 | implementation, strings |
| [122A. Lucky Division](https://codeforces.com/contest/122/problem/A) | [python](/codeforces/122A/122A.py) | 1000 | bruteforce, number theory |
| [1A. Theatre Square](https://codeforces.com/contest/1/problem/A) | [python](/codeforces/1A/1A.py) | 1000 | math |
| [58A. Chat room](https://codeforces.com/contest/58/problem/A) | [python](/codeforces/58A/58A.py) | 1000 | greedy, strings |
| [69A. Young Physicist](https://codeforces.com/contest/69/problem/A) | [python](/codeforces/69A/69A.py) | 1000 | implementation, math |
| [1669A. Division?](https://codeforces.com/contest/1669/problem/A) | [golang](/codeforces/1669A/1669A.go) | 800 | implementation |
| [1660B. Vlad and Candies](https://codeforces.com/contest/1660/problem/B) | [golang](/codeforces/1660B/1660B.go) | 800 | math |
| [1674A. Number Transformation](https://codeforces.com/contest/1674/problem/A) | [golang](/codeforces/1674A/1674A.go) | 800 | constructive algorithms, math |
| [1674B. Dictionary](https://codeforces.com/contest/1674/problem/B) | [golang](/codeforces/1674B/1674B.go) | 800 | combinatorics, math |
| [1674C. Infinite Replacement](https://codeforces.com/contest/1674/problem/C) | [golang](/codeforces/1674C/1674C.go) | 1000 | combinatorics, implementation, strings |
| [1674D. A-B-C Sort](https://codeforces.com/contest/1674/problem/D) | [golang](/codeforces/1674D/1674D.go) | 800 | constructive algorithms, implementation, sortings |
| [1692A. Marathon](https://codeforces.com/contest/1692/problem/A) | [golang](/codeforces/1692A/1692A.go) | 800 | implementation |
| [1692B. All Distinct](https://codeforces.com/contest/1692/problem/B) | [golang](/codeforces/1692B/1692B.go) | 800 | greedy, sortings |
| [1692C. Where's the Bishop?](https://codeforces.com/contest/1692/problem/C) | [golang](/codeforces/1692C/1692C.go) | 800 | implementation |
| [1692D. The Clock](https://codeforces.com/contest/1692/problem/D) | [golang](/codeforces/1692D/1692D.go) | 1100 | brute force, implementation |
| [1692F. 3SUM](https://codeforces.com/contest/1692/problem/F) | [golang](/codeforces/1692F/1692F.go) | 1300 | brute force, math |
| [1714A. Everyone Loves to Sleep](https://codeforces.com/contest/1714/problem/A) | [golang](/codeforces/1714A/1714A.go) | 900 | math |
| [1714B. Remove Prefix](https://codeforces.com/contest/1714/problem/B) | [golang](/codeforces/1714B/1714B.go) | 800 | implementation |
| [1714C. Minimum  Varied Number](https://codeforces.com/contest/1714/problem/C) | [golang](/codeforces/1714C/1714C.go) | 800 | greedy |
| [1714E. Add Modulo 10](https://codeforces.com/contest/1714/problem/E) | [golang](/codeforces/1714E/1714E.go) | 1400 | math, number theory |
| [1729A. Two Elevators](https://codeforces.com/contest/1729/problem/A) | [golang](/codeforces/1729A/1729A.go) | 800 | math |
| [1729B. Decode String](https://codeforces.com/contest/1729/problem/B) | [golang](/codeforces/1729B/1729B.go) | 800 | greedy, strings |
| [1729C. Jumping on Tiles](https://codeforces.com/contest/1729/problem/C) | [golang](/codeforces/1729C/1729C.go) | 1100 | constructive algorithms, strings |
| [1703A. YES or YES?](https://codeforces.com/contest/1703/problem/A) | [golang](/codeforces/1703A/1703A.go) | 800 | brute force, implementation, strings |
| [1703B. ICPC Balloons](https://codeforces.com/contest/1703/problem/B) | [golang](/codeforces/1703B/1703B.go) | 800 | data structures, implementation |
| [1741A. Compare T-Shirt Sizes](https://codeforces.com/contest/1741/problem/A) | [golang](/codeforces/1741A/1741A.go) | 800 | implementation, implementation, strings |
| [1741B. Funny Permutation](https://codeforces.com/contest/1741/problem/B) | [golang](/codeforces/1741B/1741B.go) | 800 | constructive algorithms, math |
| [1741C. Minimize the Thickness](https://codeforces.com/contest/1741/problem/C) | [golang](/codeforces/1741C/1741C.go) | 1100 | brute force, greedy, math, two pointers |
| [1741D. Masha and a Beautiful Tree](https://codeforces.com/contest/1741/problem/D) | [golang](/codeforces/1741D/1741D.go) | 1300 | dfs and similar, divide and conquer, graphs, sortings, trees |
| [1744A. Number Replacement](https://codeforces.com/contest/1744/problem/A) | [golang](/codeforces/1744A/1744A.go) | 800 | greedy, implementation |
| [1744B. Even-Odd Increments ](https://codeforces.com/contest/1744/problem/B) | [golang](/codeforces/1744B/1744B.go) | 800 | implementation, math |
| [1744C. Traffic Light](https://codeforces.com/contest/1744/problem/C) | [golang](/codeforces/1744C/1744C.go) | 1000 | binary search, implementation, two pointers |
| [1744D. Divisibility by 2^n](https://codeforces.com/contest/1744/problem/D) | [golang](/codeforces/1744D/1744D.go) | 1200 | greedy, math, sortings |
| [1741E. Sending a Sequence Over the Network](https://codeforces.com/contest/1741/problem/E) | [golang](/codeforces/1741E/1741E.go) | 1600 | dp |
| [1749A. Cowardly Rooks](https://codeforces.com/contest/1749/problem/A) | [golang](/codeforces/1749A/1749A.go) | 800 | greedy, implementation |
| [1749B. Death's Blessing](https://codeforces.com/contest/1749/problem/B) | [golang](/codeforces/1749B/1749B.go) | 900 | greedy |
| [1747A. Two Groups](https://codeforces.com/contest/1747/problem/A) | [golang](/codeforces/1747A/1747A.go) | 800 | constructive algorithms |
| [1747B. BAN BAN](https://codeforces.com/contest/1747/problem/B) | [golang](/codeforces/1747B/1747B.go) | 900 | constructive algorithms |
| [1747C. Swap Game](https://codeforces.com/contest/1747/problem/C) | [golang](/codeforces/1747C/1747C.go) | 1200 | games |
| [1760A. Medium Number](https://codeforces.com/contest/1760/problem/A) | [golang](/codeforces/1760A/1760A.go) | 800 | implementation, sortings |
| [1760B. Atilla's Favorite Problem](https://codeforces.com/contest/1760/problem/B) | [golang](/codeforces/1760B/1760B.go) | 800 | greedy, implementation, strings |
| [1760C. Advantage](https://codeforces.com/contest/1760/problem/C) | [golang](/codeforces/1760C/1760C.go) | 800 | data structures, implementation, sortings |
| [1760D. Challenging Valleys](https://codeforces.com/contest/1760/problem/D) | [golang](/codeforces/1760D/1760D.go) | 1000 | implementation, two pointers |
| [1760E. Binary Inversions](https://codeforces.com/contest/1760/problem/E) | [golang](/codeforces/1760E/1760E.go) | 1100 | data structures, greedy, math |
| [1760F. Quests](https://codeforces.com/contest/1760/problem/F) | [golang](/codeforces/1760F/1760F.go) | 1500 | binary search, greedy, sortings |
| [1676A. Lucky?](https://codeforces.com/contest/1676/problem/A) | [golang](/codeforces/1676A/1676A.go) | 800 | implementation |
| [1676B. Equal Candies](https://codeforces.com/contest/1676/problem/B) | [golang](/codeforces/1676B/1676B.go) | 800 | greedy, math, sortings |
| [1676C. Most Similar Words](https://codeforces.com/contest/1676/problem/C) | [golang](/codeforces/1676C/1676C.go) | 800 | brute force, greedy, implementation, implementation, math, strings |
| [1676D. X-Sum](https://codeforces.com/contest/1676/problem/D) | [golang](/codeforces/1676D/1676D.go) | 1000 | brute force, greedy, implementation |
| [1760G. SlavicG's Favorite Problem](https://codeforces.com/contest/1760/problem/G) | [golang](/codeforces/1760G/1760G.go) | 1700 | bitmasks, dfs and similar, graphs |
| [1690D. Black and White Stripe](https://codeforces.com/contest/1690/problem/D) | [golang](/codeforces/1690D/1690D.go) | 1000 | implementation, two pointers |


## acm.timus
| Problem | Solution | Difficulty | Tags | University course module |
|-|-|-|-|-|
| [1005. Stone Pile](https://acm.timus.ru/problem.aspx?num=1005) | [c++](/acm.timus/1005.cpp) | 84 | problem for beginners | 1 |
| [1155. Troubleduons](https://acm.timus.ru/problem.aspx?num=1155) | [c++](/acm.timus/1155.cpp) | 305 |  | 1 |
| [1296. Hyperjump](https://acm.timus.ru/problem.aspx?num=1296) | [c++](/acm.timus/1296.cpp) | 80 |  | 1 |
| [1401. Gamers](https://acm.timus.ru/problem.aspx?num=1401) | [c++](/acm.timus/1401.cpp) | 321 |  | 1 |
| [2025. Line Fighting](https://acm.timus.ru/problem.aspx?num=2025) | [c++](/acm.timus/2025.cpp) | 77 |  | 1 |
| [1207. Median on the Plane](https://acm.timus.ru/problem.aspx?num=1207) | [c++](/acm.timus/1207.cpp) | 177 | geometry | 2 |
| [1322. Spy](https://acm.timus.ru/problem.aspx?num=1322) | [c++](/acm.timus/1322.cpp) | 324 |  | 2 |
| [1444. Elephpotamus](https://acm.timus.ru/problem.aspx?num=1444) | [c++](/acm.timus/1444.cpp) | 639 | geometry | 2 |
| [1604. Country of Fools](https://acm.timus.ru/problem.aspx?num=1604) | [c++](/acm.timus/1604.cpp) | 220 |  | 2 |
| [1726. Visits](https://acm.timus.ru/problem.aspx?num=1726) | [c++](/acm.timus/1726.cpp) | 197 |  | 2 |
| [1067. Disk Tree](https://acm.timus.ru/problem.aspx?num=1067) | [c++](/acm.timus/1067.cpp) | 486 | data structures | 3 |
| [1494. Monobilliards](https://acm.timus.ru/problem.aspx?num=1494) | [c++](/acm.timus/1494.cpp) | 131 | data structures | 3 |
| [1521. War Games 2](https://acm.timus.ru/problem.aspx?num=1521) | [c++](/acm.timus/1521.cpp) | 192 | data structures | 3 |
| [1628. White Streaks](https://acm.timus.ru/problem.aspx?num=1628) | [c++](/acm.timus/1628.cpp) | 293 | data structures | 3 |
| [1650. Billionaires](https://acm.timus.ru/problem.aspx?num=1650) | [c++](/acm.timus/1650.cpp) | 339 | data structures | 3 |
| [1080. Map Coloring](https://acm.timus.ru/problem.aspx?num=1080) | [c++](/acm.timus/1080.cpp) | 229 | graph theory | 4 |
| [1160. Network](https://acm.timus.ru/problem.aspx?num=1160) | [c++](/acm.timus/1160.cpp) | 226 | graph theory | 4 |
| [1162. Currency Exchange](https://acm.timus.ru/problem.aspx?num=1162) | [c++](/acm.timus/1162.cpp) | 342 | graph theory | 4 |
| [1450. Russian Pipelines](https://acm.timus.ru/problem.aspx?num=1450) | [c++](/acm.timus/1450.cpp) | 225 | graph theory | 4 |
| [1806. Mobile Telegraphs](https://acm.timus.ru/problem.aspx?num=1806) | [c++](/acm.timus/1806.cpp) | 402 | graph theory | 4 |


