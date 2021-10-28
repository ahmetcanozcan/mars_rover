# Mars Rover

![example workflow](https://github.com/ahmetcanozcan/mars_rover/actions/workflows/ci.yaml/badge.svg)

## Usage
```bash
# clone repository
git clone https://github.com/ahmetcanozcan/mars_rover.git
# go to project directory
cd mars_rover
go run .
```

## What is Mars Rover?
a cli application that takes a grid size, start position of the rovers, and a list of instructions and returns the final position of the rovers.

## Input format

The first line of input is the upper-right coordinates of the plateau, the lower-left coordinates are
assumed to be 0,0.

The rest of the input is information pertaining to the rovers that have been deployed. Each rover
has two lines of input. The first line gives the rover's position, and the second line is a series of
instructions telling the rover how to explore the plateau.


## Error Handling

### Starting Position
if starting position is outbound of the plateau, the rover will be placed at the default position (0,0)

### Moving Unavailable Location
if the rover tries to move to an unavailable location, it means that the rover is trying to move out of the plateau or it is trying to move to a location that is already occupied by another rover. In this case, the rover will not move and will continue with the next action.

### Invalid Action
the actions can be `M`, `L` or `R`. if the actions are not `M`, `L` or `R`, the rover will continue with the next action.



