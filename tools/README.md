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