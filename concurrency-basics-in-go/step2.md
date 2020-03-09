# Blocking with channels
Channels are conduits to transfer data in Go. Channels can also be used as blockers to make sure concurrent code does what we want. Let's first again identify by ourselves where the data races are happening in this code. Here's the [documentation](https://golang.org/doc/articles/race_detector.html) on data races from the previous scenario if needed.

Open up the `blocking_with_channels.go` file.

>>Q1: What line is writing in the `updateUserId` function?<<
=~= 24

>>Q2: What line is reading in the `updateUserId` function?<<<
=~= 27

Go ahead and run the code with the `-race` flag to confirm your suspicions and change your answers as needed.


# Blocking with unbuffered channels
Unbuffered channels can be used as defaults to block operations until another operation has finished. 

>>Q3: What operation do we want to block so we don't have a data race?<<
(*) Read
( ) Write

# Challenge
Change the code using an unbuffered channel to block and synchronize the code so a data race no longer happens. The output should be 200 if the code is synchronized correctly. 

# Conclusion
In this exercise you succesfully implemented unbuffered channels to synchronize concurrent code. Great job!

In the next step we'll further take at an alternative way to synchronize concurrent code with WaitGroups. Once you're ready, click continue to check your answers and let's go!