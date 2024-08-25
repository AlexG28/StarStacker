# StarStacker

## What is StarStacker
StarStacker is a program that stacks a number (at least 2) night sky images on top of one another in order to produce a cleaner image as its output
This is the same algorithm used by telescopes and astrophotography modes in smartphones

## How to use StarStacker: 
- use StarStacker by compiling the code with `go build .` in the `/src` directory. This will produce a `.exe` file. 
- once a binary is made, call the binary with `./src <filepath>` with the filepath being the filepath to the directory where your night sky images are located
- StarStacker currently only supports PNG images to pngs lack of loss in image quality

## How does StarStacker work? 
- Starstacker contains multiple steps listed below: 
1. Each given png file is preprocessed by passing it pixel by pixel and setting all 'star' pixels to 1 and all other pixels to 1. 
2. The binary sky images are passed into a CCL algorithm which maps out every individual star and saves the locations of each star to an array
3. A delauney triangulation algorithm is ran on the array of points from step 2 in order to create a triangulation of all the stars in the image 
4. A reference image is chosen and the translation is calculated based on the difference between the reference image's triangulation and all 
other images' triangulations. 
5. Once all translations are calculated, each image is stacked on top of the reference image which each pixels' RGB values being averaged out. 
6. The stacked image is saved. 

## Generating sample images 
- `/validate/shiftImage.py`has been specifically written to generate noisy copies of an input images with the possibility of offsetting them. 
- Use this script by downloading a random png of the night sky off the internet. Paste in the path of the image into the script code and 
set all the parameters you want such as noise levels, shift and number of copies as well as the output folder. 
- Run the script. It will paste the output into the output folder after which you can run StarStacker on that directory. 

## Notable limitations 
- The delauney triangulation algorithm implementation in StarStacker is suboptimal with a complexity of O(n^2). It is possible to reduce this value to 
at least O(nlogn) however this would significantly increase the complexity of the project without a meaningful increase in its performance besides 
processing speed. For furhter information see this [post](https://stackoverflow.com/questions/40934453/implementing-bowyer-watson-algorithm-for-delaunay-triangulation/59582271#59582271)
- The translation calculation only takes into account linear translations with no consideration for zoom or rotation
- only PNG files are supported