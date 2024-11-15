# Group_G_week_11
## Project_Overview
The Purpose of the project is to learn diffrent mathematical functions including add, substract,divide,square along with this,
this project implements a simple calculator module in Go.
Moreover, we have initiated numbers divided by zero and tested it using error handling and used benchmarking.
### Instructions
### How to run the Program:
First we initialize the go module and use the module by importing it in our go project.
for calculator.go file, we have written code and created function divide with valid parameters. After that, we created square function and returned the square of a
### How to execute the tests and benchmarks:
To execute the test file, we added Calculator_Test.go file then write code for test.go by adding functions and execute with go test -v command to check the output.
For benchmark we added functions in the same file and checked the output.
### Expected output of the tests
Output for calculator.go file would be a/b(10/2=5) and for the square functions (a2)(2)2=4.
and output for calculator_test.go file would be passed or failed
The expected output for benchmark was 0.3098 ns/op. 
### Explaination of test and benchmark design
test design:
1. divide: we validates the correctness of division for basic inputs for ex:Divide(30,7)
after that we added if condition to check the number can't divide by zero.
2. square: we verified calculation for innput for ex:square(29).

Benchmark design:
1. divide: we tested the divide function using common numbers where the denominator is not zero this will check functions performance and conditions.
2. square: we tested the function with various types of numbers which handles both numbers like normal and extreme cases.

## Group Contributions:
Task1 completed by Meghav Shah
Task2 completed by Hetvi Patel and Shikhar Gupta
Task3 completed by Himani Patel and Jenish Gaudani
Task4 completed by Maulik Avaiya and Rajveersinh Zala
PDF work done by all group members.
