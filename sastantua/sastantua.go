package main

import (
    "fmt"
    "flag"
    "strings"
    "strconv"
)

func sastantua(size int) {
	// first, get the total row count
    // and determine at the same time
    // the max blank space count
    blankCount := 0;
    rowCount := 0;
    // we will need an offset to increment
    // for even tiers. At tier 0 its value is 2 then, it will be incremented.
    offset := 2;
    // the offset doesn't exist for the last tier iteration
    // so we can already take it off from the blank space count.
    // (to avoid if statement in the loop)
    blankCount -= offset;
    for i := 0; i < size; i++ {
        // the row count for a tier is defined by n+3 (if n starts at 0)
        rowCount += i + 3;
        // blank space take in account the number of row per tier (which
        // creates a blank space each), then the offset applied per tier.
        // (please note the -1 because spaces= n-1 rows).
        blankCount += i + 3 - 1 + offset;
        if i % 2 == 0 {
			offset++;
		}
	}
    // now we can draw it
    row := 0;  // global row index
    offset = 2;   // reset the offset
    wallWidth := 1; // wall width (initial value is 1)
    isEven := 0
    if size % 2 == 0 { isEven = 1; } // check if size is even, useful later
    halfSize := size / 2;
    for i :=0; i < size; i++ {
        // floor = row index inside the current tier
        for floor := 0; floor < i + 3; floor++ {
            // draw blank spaces
            fmt.Print(strings.Repeat(" ", blankCount));
            // draw side
            fmt.Print("/");
            // check if we are in a floor without door parts
            if (row < rowCount - size) {
                fmt.Print(strings.Repeat("*", wallWidth));
            } else {
                // draw only left wall
                halfWall := (wallWidth / 2) - halfSize;
                fmt.Print(strings.Repeat("*", halfWall));
                for door := 0; door < size; door++ {
					// check if we need to draw the lock
					// 1 - the lock is only visible on size >= 3
					// 2 - check horizontally (second last column)
					// 3 - check vertically (centered)
                    if size >= 3 && door == size - 2 && rowCount - 1 - row == halfSize {
						fmt.Print("$");
					} else {
						fmt.Print("|");
					}
                }
                // to print the remaining wall, we check also if the size of
                // sastantua is an even number. In this case the door can't
                // be centered and is shifted to the left, so we need an extra
                // star to compensate this assymetry.
                fmt.Print(strings.Repeat("*", halfWall + isEven));
            }
            // draw other side and line feed
            fmt.Print("\\");
            // funny, you could fmt.Print(strings.Repeat(' ', blankCount));
            fmt.Print("\n");
            // decrement the blank space count for each floor
            // but not if we have processed the last one,
            // because a special decrement using the offset will be used
            // later. Same for the wall increment.
            if floor + 1 < i + 3 {
                blankCount--;
                // for each floor in the tier, the wall width is increased by 2
                // (one star at each side)
                wallWidth += 2;
            }
            row++;
        }
        // check if the tier index is even to increment the offset
        if i % 2 == 0 {
			offset++;
		}
        // decrease the blank count using the current offset
        blankCount -= offset;
        // add the offset to the wall width too for next tier
        // (for symmetry we need the offset twice)
        wallWidth += offset + offset;
    }
}

func main() {

    flag.Parse()
    args := flag.Args()
    for _,arg := range args {
        if size,err := strconv.Atoi(arg); err != nil {
            fmt.Printf("***Argument \"%s\" Invalid***\n", arg)
        } else {
            fmt.Printf("Size %s\n", arg)
            sastantua(size);
        }
        fmt.Printf("\n\n\n")
    }
}