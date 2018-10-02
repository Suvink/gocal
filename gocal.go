
package main

//importing fmt package for basic functions
import "fmt"

//declaration of variables
var num1 float32
var num2 float32
var num3 float32
var operator string
var view bool
var mode string


//declaration of functions
func add(a float32, b float32)float32{
  return a+b
}

func substract(a float32, b float32)float32{
  return a-b
}

func multiply(a float32, b float32)float32{
  return a*b
}

func divide(a float32, b float32)float32{
  return a/b
}

//main function
func main(){



fmt.Print("Select a mode. Simple | Advanced ")
fmt.Scanln(&mode)
if(mode == "advanced"){

      fmt.Print("Enter a number to find the square:  ")
      fmt.Scanln(&num3)
      fmt.Print("Your answer is: ")
      fmt.Println(multiply(num3,num3))
      fmt.Println("\n")
  


}else if(mode == "simple"){

      fmt.Print("Enter No 1: ")
      fmt.Scanln(&num1)

      fmt.Print("Enter No 2: ")
      fmt.Scanln(&num2)

      fmt.Print("Enter operator: ")
      fmt.Scanln(&operator)

      if (operator == "+"){
        fmt.Print("Your answer is: ")
        fmt.Println(add(num1,num2))
        fmt.Println("\n")

      }else if (operator == "-"){
        fmt.Print("Your answer is: ")
        fmt.Println(substract(num1,num2))
        fmt.Println("\n")

      }else if (operator == "*"){
        fmt.Print("Your answer is: ")
        fmt.Println(multiply(num1,num2))
        fmt.Println("\n")

      }else if (operator == "/"){
        fmt.Print("Your answer is: ")
        fmt.Println(divide(num1,num2))
        fmt.Println("\n")
      }

    }




}
