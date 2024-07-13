# Starcounter 

## What is starcounter 
Starcounter is a program that calculates the number of stars in a photo 

## How to use starcounter: 
- use starcounter by compiling the code with `go build .`
- once a binary is made, call the binary with `./app filename` with filename being the name of the file you want to process 

## How does starcounter work? 
- starcounter works by first processing the given input image by raising its contrast and turning it into a binary black or white photo with all brighht (star) pixels
above a certain threshold being white and all other pixels being dark. 
- once this binary image is achieved, a connected component labeling algorithm is ran on this binary image which calculates the location and the count of every cluster of white 
pixels (stars). This information is saved to a `stars.txt` file
