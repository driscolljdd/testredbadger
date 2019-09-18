package main

import (
	"fmt"
	"github.com/driscolljdd/testredbadger/Structures"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

		// Grab our sample data
		raw, err := ioutil.ReadFile("sample.data")

		if(err != nil) {

			fmt.Println("Could not read from sample.data")

			// Exit with a unix fault status
			os.Exit(1)
		}

		// Break this up into lines
		lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

		// As we're getting towards the end and time is running short I'll skip extensive error checking here but of course there's plenty that could be done with checking number of characters, length etc.
		// First line is the max X and max Y for our grid; fetch those
		parts := strings.Split(strings.TrimSpace(lines[0])," ")

		// What the hell, let's do a bit of validation
		if(len(parts) != 2) {

			// Doing this makes me think this would look better in a go test file. Were this ever to be a 'production' thing a whole test file could be dedicated to just testing the file itself
			fmt.Println("First line is not valid; should have two integers with a space between")
			os.Exit(0)
		}

		// Parse our strings
		worldX, err := strconv.Atoi(parts[0])

		if(err != nil) {

			fmt.Println("First line is not valid; should have two integers with a space between")
			os.Exit(0)
		}

		worldY, err := strconv.Atoi(parts[1])

		if(err != nil) {

			fmt.Println("First line is not valid; should have two integers with a space between")
			os.Exit(0)
		}

		// Let's build our world
		flatMarsTheory := Structures.CreateWorld(worldX, worldY)

		// Trim off the first line we don't need anymore
		lines = lines[1:]

		// Loop through the rest of the lines
		var newRobot = true
		for _, line := range lines {

			if(newRobot) {

				parts = strings.Split(line, " ")

				if(len(parts) != 3) {

					fmt.Println("Invalid robot initialisation string: [", line, "]")

					// Set up a default robot
					flatMarsTheory.NewRobot(0, 0, "N")

				} else {

					var newRobotX, newRobotY int
					if newRobotX, err = strconv.Atoi(parts[0]); err != nil {

						newRobotX = 0
					}

					if newRobotY, err = strconv.Atoi(parts[1]); err != nil {

						newRobotY = 0
					}

					// Set up our new world
					flatMarsTheory.NewRobot(newRobotX, newRobotY, parts[2])
				}

				newRobot = false
				continue
			}

			// If this line is a blank line, end our current robot
			
		}

		fmt.Println(lines)
	}
