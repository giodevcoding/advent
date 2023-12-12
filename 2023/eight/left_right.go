package eight

import "fmt"

type RoutesMap = map[string][2]string

func RunPartOne(input []string) {
	fmt.Println(LeftRightMinSteps(input))
}

func RunPartTwo(input []string) {
    fmt.Println(LeftRightMinStepsGhost(input))
}

func LeftRightMinSteps(input []string) int {
	sequence := GetLeftRightSequence(input)
	routes := GetRoutes(input)

	steps, sequenceIndex := 1, 0
	currentRoute := "AAA"

	for routes[currentRoute][sequence[sequenceIndex]] != "ZZZ" {
		currentRoute = routes[currentRoute][sequence[sequenceIndex]]
		steps++
		if sequenceIndex >= len(sequence)-1 {
			sequenceIndex = 0
		} else {
			sequenceIndex++
		}
	}

	return steps
}

func LeftRightMinStepsGhost(input []string) int {
	sequence := GetLeftRightSequence(input)
	routes := GetRoutes(input)

    startingRoutes := GetStartingRoutes(routes)
    stepsTaken := make([]int, len(startingRoutes))

    for i, startingRoute := range startingRoutes {
        steps, sequenceIndex := 1, 0
        currentRoute := startingRoute

        for !RouteEndsWithZ(routes[currentRoute][sequence[sequenceIndex]]) {
            currentRoute = routes[currentRoute][sequence[sequenceIndex]]
            steps++
            if sequenceIndex >= len(sequence)-1 {
                sequenceIndex = 0
            } else {
                sequenceIndex++
            }
        }

        stepsTaken[i] = steps

	}

    result := GetSliceLCM(stepsTaken)

	return result
}

func GetSliceLCM(numbers []int) int {
    result := numbers[0]

    for i := 1; i < len(numbers); i++ {
        result = GetLCM(numbers[i], result)
    }

    return result
}

func GetLCM(a, b int) int {
    return a * b / GetGCD(a, b)
}

func GetGCD(a, b int) int {
 if (a == 0) {
     return b
 }

 return GetGCD(b % a, a)
}

func RouteEndsWithZ(route string) bool {
    return route[len(route)-1] == 'Z'
}

func GetStartingRoutes(routes RoutesMap) (startingRoutes []string) {
    for routeName := range routes {
        if routeName[len(routeName)-1] == 'A' {
            startingRoutes = append(startingRoutes, routeName)
        }
    }

    return startingRoutes
}

func GetRoutes(input []string) RoutesMap {
	routes := make(RoutesMap)

	routeStrings := input[2:]
	for _, routeStr := range routeStrings {
		routeName := routeStr[0:3]
		routeLeft := routeStr[7:10]
		routeRight := routeStr[12:15]

		routes[routeName] = [2]string{routeLeft, routeRight}
	}
	return routes
}

func GetLeftRightSequence(input []string) (sequence []int) {
	sequenceString := input[0]
	for _, char := range sequenceString {
		if char == 'R' {
			sequence = append(sequence, 1)
		} else {
			sequence = append(sequence, 0)
		}
	}
	return sequence
}
