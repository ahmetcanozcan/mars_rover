package main

import (
	"app/action"
	"app/common"
	"app/rover"
	"app/terrain"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func startApp(rd io.Reader, wr io.Writer) {

	reader := bufio.NewReader(rd)

	// read the first line
	bytes, _, err := reader.ReadLine()
	checkError(err)
	text := string(bytes)

	sizeStr := strings.Split(text, " ")
	if len(sizeStr) != 2 {
		panic("Invalid input")
	}
	width, err := strconv.Atoi(sizeStr[0])
	checkError(err)
	height, err := strconv.Atoi(sizeStr[1])
	checkError(err)

	mars := terrain.New("Mars", width, height)
	result := ""
	for i := 0; i < 2; i++ {
		bytes, _, err := reader.ReadLine()
		checkError(err)
		text := string(bytes)

		roverInfo := strings.Split(text, " ")
		if len(roverInfo) != 3 {
			panic("Invalid input")
		}

		x, err := strconv.Atoi(roverInfo[0])
		checkError(err)
		y, err := strconv.Atoi(roverInfo[1])
		checkError(err)

		directionStr := roverInfo[2]
		direction, err := common.NewDirectionFromString(directionStr)
		checkError(err)

		startLoc := common.NewLocation(x, y)
		r := rover.New(startLoc, direction)
		mars.AddObject(r)

		bytes, _, err = reader.ReadLine()
		checkError(err)
		text = string(bytes)

		commands := strings.Split(text, "")
		for _, s := range commands {
			act, err := action.FromString(s)
			if err == action.ErrInvalidAction { // Ignore invalid actions
				continue
			}
			checkError(err)
			r.Do(act)
		}
		result += r.String() + "\n"
		checkError(err)
	}
	fmt.Fprint(wr, result)
}
