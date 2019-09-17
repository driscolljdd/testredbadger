package Structures

import "strings"

	type world struct {

		sizeX, sizeY int
		scentPositions []position
		robots []robot
		currentRobot robot
		moves movementMap
	}

	func (w *world) NewRobot(x, y int) {

		// Do we have an existing robot?
		if(w.currentRobot.moves > 0) {

			// Save down the current robot
			w.robots = append(w.robots, w.currentRobot)

			// And fetch a new one
			w.currentRobot = robot{ location:position{ X: x, Y: y} }
		}
	}

	func (w *world) Instruction(command string) {

		// Uppercase the command
		command = strings.ToUpper(command)

		switch(command) {

			case "L", "R":

				// Spin around
				newDirection, valid := w.moves.turn[w.currentRobot.direction][command]

				if(valid) {

					w.currentRobot.direction = newDirection
				}

			case "F":

				// We have movement. Calculate where we're going to be next
				newPositionX, validX := w.moves.movement[w.currentRobot.direction]["F"]["X"]
				newPositionY, validY := w.moves.movement[w.currentRobot.direction]["F"]["Y"]

				if(!validX || !validY) {

					return
				}

				// OK we have the new position we will move to, but does it take us outside the world?
				

		}
	}

	func CreateWorld(maxX, maxY int) world {

		// Set up the size of our world
		newWorld := world{ sizeX: maxX, sizeY: maxY }

		// Set up a first robot
		newWorld.currentRobot = robot{}

		// Load up the current movements map
		newWorld.moves = movementMap{}

		// Initialise our slice for safe robot scent squares and used robots
		newWorld.robots = make([]robot, 0)
		newWorld.scentPositions = make([]position, 0)

		return newWorld
	}
