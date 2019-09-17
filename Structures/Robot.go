package Structures

import "strings"

	type robot struct {

		location position
		isLost bool
		direction string
		moves int
	}

	func (r *robot) Exec(command string) {

		// Get our movement map
		r.MoveMap = GetMovementMap()

		// Check this command is basically valid
		if(len(command) > 1) {

			return
		}

		// Uppercase the command
		command = strings.ToUpper(command)

		switch(command) {

			case "L", "R":

				// Spin around
				newDirection, valid := r.MoveMap.turn[r.Direction][command]

				if(valid) {

					r.Direction = newDirection
				}

			case "F":

				// We have movement. Calculate where we're going to be next
				newPositionX, validX := r.MoveMap.movement[r.Direction]["F"]["X"]
				newPositionY, validY := r.MoveMap.movement[r.Direction]["F"]["Y"]

				if(!validX || !validY) {

					return
				}

				// OK we have the new position we will move to, but does it take us outside the world?

		}
	}
