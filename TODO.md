# TODO

- add actual CLI with options and flags
  - [x] add commands for creating the maze with different algorithms
  - [x] add flag to allow specifying size for the maze when creating it
  - 

- allow for arbitrary sizes of the maze
  - [ ] change cell content from hexadecimal to just decimal numbers
  - [ ] adjust cell width based on the number of digits in the largest number

-  add maze coloring
  - [ ] darken color as the distance from starting location increases

-  draw solutions to the maze graphically, instead of numerically
  - [ ] add option to either specify numbers in the cells for distances or draw arrows from start to terminal locations


-  animate dijkstra spread from starting cell

-  clean up terminal display of the maze
  - [ ] clean up characters, don't use `+` except for intersections, etc...

-  parameter to change the starting cell

# Refactor/Cleanup:

- [ ] Change grid to use uints instead of ints
- [x] Extract width and height flags into higher level create options
