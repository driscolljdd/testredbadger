package Structures

	// Creating a movement map could have been done either with a struct or with a nested map. Maps aren't thread safe unless you use sync.Lock, but
	// it's easier to use the data from the input directly against the map, and we can use the 'value, exists := map[x]' format to quickly determine if a direction exists in the map and so is valid
	// Ultimately that makes a map, for this anyway, slightly more desireable than a struct
	type movementMap struct {

		movement map[string]map[string]map[string]int
		turn map[string]map[string]string
	}



	// There is probably a smarter algorithmic way to work this out but as a quick exercise this is simple and clear.
	// Translate the direction a robot is facing and the movement it makes into the amount which should be added to the X/Y position
	func GetMovementMap() movementMap {

		moveMap := movementMap{}

		// Build up a steadily more complex initialised map structure
		xyMap := make(map[string]int)

		// Create the middle part
		actionMap := make(map[string]map[string]int)
		actionMap["L"] = xyMap
		actionMap["R"] = xyMap
		actionMap["F"] = xyMap

		// Now the top part
		moveMap.movement = make(map[string]map[string]map[string]int)
		moveMap.movement["N"] = actionMap
		moveMap.movement["E"] = actionMap
		moveMap.movement["S"] = actionMap
		moveMap.movement["W"] = actionMap

		// North
		moveMap.movement["N"]["L"]["X"] = 0
		moveMap.movement["N"]["L"]["Y"] = 1

		moveMap.movement["N"]["R"]["X"] = 0
		moveMap.movement["N"]["R"]["Y"] = -1

		moveMap.movement["N"]["F"]["X"] = 0
		moveMap.movement["N"]["F"]["Y"] = 1


		// East
		moveMap.movement["E"]["L"]["X"] = 0
		moveMap.movement["E"]["L"]["Y"] = 1

		moveMap.movement["E"]["R"]["X"] = 0
		moveMap.movement["E"]["R"]["Y"] = -1

		moveMap.movement["E"]["F"]["X"] = 1
		moveMap.movement["E"]["F"]["Y"] = 0


		// South
		moveMap.movement["S"]["L"]["X"] = 1
		moveMap.movement["S"]["L"]["Y"] = 0

		moveMap.movement["S"]["R"]["X"] = -1
		moveMap.movement["S"]["R"]["Y"] = 0

		moveMap.movement["S"]["F"]["X"] = 0
		moveMap.movement["S"]["F"]["Y"] = -1


		// West
		moveMap.movement["W"]["L"]["X"] = 0
		moveMap.movement["W"]["L"]["Y"] = -1

		moveMap.movement["W"]["R"]["X"] = 0
		moveMap.movement["W"]["R"]["Y"] = 1

		moveMap.movement["W"]["F"]["X"] = -1
		moveMap.movement["W"]["F"]["Y"] = 0

		// What happens if we turn?
		simpleMap := make(map[string]string)
		moveMap.turn = make(map[string]map[string]string)
		moveMap.turn["N"] = simpleMap
		moveMap.turn["E"] = simpleMap
		moveMap.turn["S"] = simpleMap
		moveMap.turn["W"] = simpleMap

		moveMap.turn["N"]["L"] = "W"
		moveMap.turn["N"]["R"] = "E"
		moveMap.turn["E"]["L"] = "N"
		moveMap.turn["E"]["R"] = "S"
		moveMap.turn["S"]["L"] = "E"
		moveMap.turn["S"]["R"] = "W"
		moveMap.turn["W"]["L"] = "S"
		moveMap.turn["W"]["R"] = "N"

		return moveMap
	}
