# Data races
Data races are dangerous because two threads access a shared memory location without synchronization.

Use this [documentation](https://golang.org/doc/articles/race_detector.html) on data races from Go to answer the below questions.

>>Q1: In order for a data race to occur, one operation must be what?<<
=~= write


# Detecting data races in Go
## Identifying reads/writes
Take a look at the code in `data_race_example.go`. Before we use Go's built in functionalities to identify the data race, attempt to find it yourself.

>>Q2: In the function `getUserId` what line is a write operation?<<
=~= 25

>>Q3: In the function `getUserId` what line is a read operation?<<
=~= 28

>>Q4: What do you think the output will be once you run the file `data_race_example.go`?<<
(*) Output will either be 0 or 137.
( ) Output will be 0.
( ) Output will be 137.

Run the file with `go run data_race_example.go` in the command line and confirm your answer.

## Running -race to detect data races

Go's built in `-race` will tell you where the data races are with further information.

Run the command `go run -race data_race_example.go`. Compare the output with your answers, if anything is different, change your answers.

# Conclusion
The `-race` flag is a useful functionality to use to detect and gather information on potential races in your code.

In the next step we will look at how we can avoid data races by blocking with channels.