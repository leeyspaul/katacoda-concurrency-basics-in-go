# Blocking with Channels
Channels are conduits to transfer data in Go. Channels can be used as blockers to make sure concurrent code does what we want.

Open up the `blocking_with_channels.go` file and answer the questions below.

>>Q1: What line is writing in the `updateUserId` function?<<
=~= 24

>>Q2: What line is reading in the `updateUserId` function?<<<
=~= 27

Go ahead and run the code with the `-race` flag to confirm your suspicions and change your answers as needed.

>>Q3: What operation do we want to block so we don't have a data race?<<
(*) Read
( ) Write


# Blocking with unbuffered channels
This [documentation](https://tour.golang.org/concurrency/2) on channels will be useful for the below challenge.

Unbuffered channels can be used as defaults to block operations until another operation has finished. Take a look at the tour and familiarize yourself with channels at a high level.

# Challenge
Change the code using an unbuffered channel to synchronize the code. The output should be 200 if the code is synchronized correctly. 

# Conclusion
In this exercise you succesfully implemented unbuffered channels to synchronize concurrent code. Great job!

In the next step we will look at an alternative way to synchronize concurrent code with WaitGroups.