# Red Badger Technical Test

This is my submission for the Red Badger mars rover technical test.

## Starting Out

The first thing I did was come up with a detailed struct which mapped out how the X and Y co-ordinates of anything should change
if the thing were to move right, left or forwards, and dependent on which direction they were facing. My plan was to work from the 
inside out - first planning how to manage a single movement, and then using that algorithm iteratively to solve the problem.

I changed from a struct to a map type early in the process. The struct was elegant but it actually had nested structs with names such
as East and West; these looked great but from a coding point of view I would need to do something like a chain of select or IF 
statements to make use of them programatically. 

Moving to map data types comes with costs. Maps are not thread safe so careful consideration must be given to using them in a multi
threaded application. The sync package has RWMutexes available which allow for thread safe read and write locking. Maps also have to
be initialised, which can be a pain when you have nested maps. On the upside, they allowed me to feed in the direction the object is 
facing and the action to be carried out as a key to a nested map, allowing easy calculations of next positions.

I realised after re-reading the instructions that in fact the robots only turn left and right; they don't actually move left and right.
As such my map (MovementMap.go/movementMap) is more populous with data than it needs to be; only the F elements need remain but I
left them in - actually part of the spec called for supporting future movements and this map becomes the way to do this. Consider 
adding B to the map - this would allow calculations to allow the program to simulate a backwards move, or J to consider a jump.

## Model, Hub, Controller

I suppose one way to describe this structure might be model, hub, controller. Technically my World struct is also a model but it has
become, in practice, the hub which holds most of the code together. The Mars.go main application is truly a controller in the sense it
does little more than input data, carry out trivial error checking and then interact with the World struct. The world struct uses the
other models (robot, movementMap) in order to carry out it's task. You could almost argue world is the controller and Mars.go the 
view. Almost.

## Improvements

The exercise was fun. I did in fact end up spending no more than the ideal two to three hours (I think). Completing the task after
work there wasn't much more time than that available anyway. Had I had more time, I would probably have done more testing. In fact,
one good exercise would be to simply walk through the sample data in a pen and pencil exercise and use the results to create some
go tests which I could use to validate the end code.

The structure of the code with the world struct having a method to create a new robot and run commands lends itself to remodelling
as a web service - it could be possible to do simple json web service calls with some information like a robot name and a movement,
and have the web service create the respective robot if it does not exist already, and keep track of movements. There's plenty of 
new places this could go.

## Other Code

I would also like you to view my recent Process Group project. It's something I did over the course of an evening actually quite
recently. It makes it trivial to gain advanced control over goroutines in an application - making advanced use of wait groups, 
contexts and channels to achieve this. It's a nice demonstration of some of my depth of understanding of the Go language. You
can see the project [here](https://github.com/driscolljdd/processgroup).

## How To Use

To run my impelementation, grab this repo:

```shell
go get -u github.com/driscolljdd/testredbadger
```

Run the main program (Mars.go) with the go run command:

```go
go run Mars.go
```