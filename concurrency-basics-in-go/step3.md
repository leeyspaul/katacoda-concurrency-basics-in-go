# Blocking with WaitGroups
WaitGroups are a goroutine synchronization mechanism that can force a block until goroutines finish.

Open up the `blocking_with_wg.go` file and use this [documentation](https://golang.org/pkg/sync/) to answer the questions below.

>>Q1: In order to use WaitGroups what package do we need to import?<<
=~= sync


>>Q2: What WaitGroup function must we run to tell Go the number of goroutines we want to run?<<
( ) Write
( ) Block
( ) Minus
(*) Add

>>Q3: What WaitGroup function will block the main thread until a goroutine finishes?<<
( ) Write
( ) Block
(*) Wait
( ) Add

>>Q4: What WaitGroup function will tell Go we've finished with a goroutine?<<
( ) Write
( ) Block
(*) Done
( ) Add


# Challenge
Change the code using WaitGroups to synchronize the code. The output should be 200 if the code is synchronized correctly. 

# Conclusion
In this exercise you succesfully implemented WaitGroups to synchronize concurrent code. Great job!

In this Katacoda scenario you learned how to test for data races, block with Channels, and block with WaitGroups!