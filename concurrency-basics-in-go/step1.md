# Data races are dangerous
As mentioned briefly in the presentation, data races are dangerous because two threads will access a shared memory location with no synchronization on the accesses.

<pre>
>>Q1: In order for a data race to occur, one operation must be what? <<

=~= write

*Go ahead and click "check answer" to see the results!* 

# Detecting data races with Go
In Go, we can detect data races with a simple command which we'll reveal in a bit. Before we do that, take a look at the code in `data_race_example.go`. This code definitely has a data race in it. Before we use Go's magical powers to identify where the data race causing culprits are, let's try to identify it ourselves.

>>Q2: In the function `getUserId` what line is a write operation? <<

=== 25

>>Q3: In the function `getUserId` what line is a read operation? <<

=== 28

*Go ahead and click "check answer" to see the results!* 

</pre>