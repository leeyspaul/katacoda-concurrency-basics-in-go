# Data races are dangerous
Data races are dangerous because two threads will access a shared memory location with no synchronization on the accesses. Before we get started identifying a data race, a quick quiz! 

Use this [documentation](https://golang.org/doc/articles/race_detector.html) on data races from Go to answer the below questions.

>>Q1: In order for a data race to occur, one operation must be what?<<
=~= write


# Detecting data races with Go
## Identifying read/writes
In Go, we can detect data races with a simple command which you may have already seen in the relevant documentation above. Before we detect a data race with Go, take a look at the code in `data_race_example.go`. This code definitely has a data race in it. Before we use Go's magical powers to identify where the data race causing culprits are, let's try to identify it ourselves.

>>Q2: In the function `getUserId` what line is a write operation?<<
=~= 25

>>Q3: In the function `getUserId` what line is a read operation?<<
=~= 28

>>Q4: Given what you know about where the read/write operations are happening, what do you think the output will be once you run the file `data_race_example.go`?<<
(*) Output will either be 0 or 137.
( ) Output will be 0.
( ) Output will be 137.

## Running -race to detect data races
Enough hypothesizing, let's go ahead and see the results. Go ahead and run the file with `go run data_race_example.go` in the command line.

Now that you've seen the result, let's see what Go has to tell us about the read/write operations with the built in `-race` functionality. 

Run the command `go run -race data_race_example.go`. Compare the output with your answers, if anything is different, change your answers.

# Conclusion
As you can see Go specifically tells us where our data races are. The `-race` flag is a very useful functionality to use! So next time you're writing concurrent code, test your code with the `-race` flag.

In the next step we'll take a look at how we can avoid data races by blocking with channels as mentioned in the presentation. Once you're ready, click continue, check your answers and let's go!