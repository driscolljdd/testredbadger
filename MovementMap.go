package main

	// Creating a movement map could have been done either with a struct or with a nested map. Maps aren't thread safe unless you use sync.Lock, but
	// it's easier to use the data from the input directly against the map, and we can use the 'value, exists := map[x]' format to quickly determine if a direction exists in the map and so is valid
	// Ultimately that makes a map, for this anyway, slightly more desireable than a struct
	type MovementMap map[string]map[string]map[string]int



	// There is probably a smarter algorithmic way to work this out but as a quick exercise this is simple and clear.
	// Translate the direction a robot is facing and the movement it makes into the amount which should be added to the X/Y position
	func GetMovementMap() MovementMap {

		moveMap := MovementMap{}

		// North
		moveMap["N"]["L"]["X"] = 0
		moveMap["N"]["L"]["Y"] = 1

		moveMap["N"]["R"]["X"] = 0
		moveMap["N"]["R"]["Y"] = -1

		moveMap["N"]["F"]["X"] = 0
		moveMap["N"]["F"]["Y"] = 1


		// East
		moveMap["E"]["L"]["X"] = 0
		moveMap["E"]["L"]["Y"] = 1

		moveMap["E"]["R"]["X"] = 0
		moveMap["E"]["R"]["Y"] = -1

		moveMap["E"]["F"]["X"] = 1
		moveMap["E"]["F"]["Y"] = 0


		// South
		moveMap["S"]["L"]["X"] = 1
		moveMap["S"]["L"]["Y"] = 0

		moveMap["S"]["R"]["X"] = -1
		moveMap["S"]["R"]["Y"] = 0

		moveMap["S"]["F"]["X"] = 0
		moveMap["S"]["F"]["Y"] = -1


		// West
		moveMap["W"]["L"]["X"] = 0
		moveMap["W"]["L"]["Y"] = -1

		moveMap["W"]["R"]["X"] = 0
		moveMap["W"]["R"]["Y"] = 1

		moveMap["W"]["F"]["X"] = -1
		moveMap["W"]["F"]["Y"] = 0

		return moveMap
	}
