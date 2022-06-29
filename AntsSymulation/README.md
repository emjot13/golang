A simulation of ants wrriten in go.
At the start certain number of ants and leafs are randomly generated on a rectangular grid - all of these parameters can be adjusted freely.
Then for some number of iterations the grid is being updated with the following rules:
- each ant randomly chooses next square:
    - if that square is empty the ant moves to it
    - if on that squares a leaf is lying:
        - if the ant is not currently currying a leaf then it picks up the leaf
        - else it randomly chooses a adjacent square to put its current leaf down
    - if the ant can not move anywhere it stays in the same square
