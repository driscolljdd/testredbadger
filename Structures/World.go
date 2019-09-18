package Structures

import (
	"strconv"
	"strings"
)

	type world struct {

		sizeX, sizeY int
		scentPositions map[string]interface{}
		robots []robot
		currentRobot robot
		moves movementMap
	}

	func (w *world) NewRobot(x, y int, direction string) {

		// Save down the current robot
		w.robots = append(w.robots, w.currentRobot)

		// And fetch a new one
		w.currentRobot = robot{ location:position{ X: x, Y: y}, setup: true, direction: strings.ToUpper(strings.TrimSpace(direction[:1])) }
	}

	func (w *world) Instruction(command string) {

		// If we're off the grid, just stop right here (or if we haven't yet set up a robot)
		if(w.currentRobot.offGrid || !w.currentRobot.setup) {

			return
		}

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
				movementX, validX := w.moves.movement[w.currentRobot.direction]["F"]["X"]
				movementY, validY := w.moves.movement[w.currentRobot.direction]["F"]["Y"]

				// If the incoming instruction isn't in our movement map, stop here
				if(!validX || !validY) {

					return
				}

				// Calculate our new positions
				newPositionX := w.currentRobot.location.X + movementX
				newPositionY := w.currentRobot.location.Y + movementY

				// OK we have the new position we will move to, but does it take us outside the world?
				if(newPositionX < 0 || newPositionY < 0 || newPositionX > w.sizeX || newPositionY > w.sizeY) {

					// We're about to be off grid. One chance to save ourselves; has anyone else dropped off from here before, in which case we can just pretend this never happened...
					_, safe := w.scentPositions[strconv.Itoa(w.currentRobot.location.X) + "/" + strconv.Itoa(w.currentRobot.location.Y)]

					if(safe) {

						// This never happened
						return
					}

					// We are off grid; record the fact
					w.currentRobot.offGrid = true

					// We also gain immunity for this square in the future
					var void interface{}
					w.scentPositions[strconv.Itoa(w.currentRobot.location.X) + "/" + strconv.Itoa(w.currentRobot.location.Y)] = void

					// And that's all we need to do for this scenario
					return
				}

				// In this scenario, not going off grid, we just update the robot position
				w.currentRobot.location.X = newPositionX
				w.currentRobot.location.Y = newPositionY
		}
	}

	func CreateWorld(maxX, maxY int) world {

		// Set up the size of our world
		newWorld := world{ sizeX: maxX, sizeY: maxY }

		// Load up the current movements map
		newWorld.moves = movementMap{}

		// Initialise our slice for safe robot scent squares and used robots
		newWorld.robots = make([]robot, 0)

		// This is [string]interface because actually I only want the keys, not the values. And as I don't want the values, actually interface{} takes 0 bytes, so a million interface{} is 0 bytes.
		newWorld.scentPositions = make(map[string]interface{})

		return newWorld
	}
