# MATHSKILLS

Mathskills is a program that is capable of reading contents in a file and be able to perform operations like;
* Average
* Median
* Variance
* Standard Deviation

The data represents statistical population where each line contains a specific value. Here is an exaple
	189,
	176,
	847,
	243,
	981,
	...

## Objective
The main objective for mathskills is to gain a better understanding of Statistics and Mathematics.
For now we will deal with 
* Average
* Median
* variance
* Standard Deviation



## Usage
Open your terminal and clone the project's repository to you machine.
```bash
git clone
```

Move to the directory containing the project and provide some sample files containing statistical data.
The sample file should be have a txt extrension(example, data.txt) and not of any other type.
I have provided a default file which is data.txt with a few data entries inside it.


```bash
cd Math-Skills
```

For you to run the program you need to use the following command 
```go
go run . your-sample-file.txt
```
Failure to this an error will occur addressing the issue with a guideline of how you should write the command.
 
After running the program, the results of each category will be displayed in a representation similar to this;

```
Average: 132
Median: 117.5
Variance: 784.66667
Standard Deviation: 28.011902

```
Since it's statistical data the numbers can never be negative in the sample file.


## Collaborators
Mathskills was written by Wilfred Njuguna with a lot of help from articles and website to get more information about statistics and mathematics.
Here are some links for the pages :
* https://www.calculatorsoup.com/calculators/statistics/statistics.php
* https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://pkg.go.dev/os&ved=2ahUKEwiU06Gss5iGAxV8W0EAHYqGD4IQFnoECB0QAQ&usg=AOvVaw0759GFkB1SOFAEWYvEz2rG

