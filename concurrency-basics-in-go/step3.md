# Blocking with WaitGroups
As we already know, WaitGroups are a goroutine synchronization mechanism where we can force a block until goroutines finish in our code.

Open up the `blocking_with_wg.go` file.

>>Q1: In order to use WaitGroups what package do we need to import?<<
=~= sync


Go ahead and import the corresponding package so we can use WaitGroups. Use your knowledge on WaitGroups and the [docuemtation](https://golang.org/pkg/sync/) to answer the below questions.

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
Change the code using WaitGroups to block and synchronize the code so a data race no longer happens. The output should be 200 if the code is synchronized correctly. 

# Conclusion
In this exercise you succesfully implemented WaitGroups to synchronize concurrent code. Great job!

This is just the beginning of writing concurrent code that works in Go! We hope that this was useful in at least getting started and understanding how to identify data races and implement basic blocking with channels and WaitGroups!