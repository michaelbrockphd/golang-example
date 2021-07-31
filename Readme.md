# Golang Example

A mono-repo acting as my port-folio for Golang when conversing with friends, co-workers and people interviewing me for jobs.

## leet_code_attempts

As the name implies, my leet code attempts (successful or otherwise) are kept here for easier sharing.  Even when I go to my own profile it takes some digging to get my past submissions out!

## brain_teasers

Sometimes, I recall a coding exercise from my university days.  As I revisit and solve them, they will be placed here.

## conversion

From time to time, I stumble on converting from Base-10 to Base-n.  This child folder is basically me learning the method and writing Golang classes to perform that conversion to text.

## matrix

Whenever I try a new programming lanaguage, I often practice with matrix multiplication.  It's a simple exercise and one that can be easily made concurrent if I find the language in question has means of concurrent processing.

I actually had trouble with Golang and was not getting any performance improvement.  So my original exercise was put aside and I tried another challenge - seeing if how a matrix is kept in memory would impact sequential performance.

As the Golang benchmarks in the child folder show, yes!  Ensuring matrix A in the equation is column aligned actually makes a dramatic impact on the execution time.

As a little extra, I toyed with the idea of a computed identity matrix.  Effectively a matrix that has little to no memory footprint and uses simple logic to decide if zero or the identity value is returned.

## collections

While attempting Leet Code exercises, I found out the hard way that Golang does not have a stack nor queue class.

Any programmer worth their salt could easy simulate one via an array and I'm sure there is a module out there that would offer a feature rich implementation of both, I still decided to try my hand at making basic implementations of both.
