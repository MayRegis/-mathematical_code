package main

import (
  "fmt"
  "math"
)

//calcula o fatorial
func fatorial(num float64) float64{
  var i float64
  for i = 1.0; num > 1 ;num = num - 1{
    i = i * num
  }
  return i
}

func main() {
  var x, y, lambda, sum float64
  var op int
  e := 2.71828 

  //menu
  fmt.Print(" ______________________\n")
  fmt.Print("|        OPÇÕES        |\n")
  fmt.Print("|                      |\n")
  fmt.Print("| 1. P(X = x)          |\n")
  fmt.Print("| 2. P(X > x)          |\n") 
  fmt.Print("| 3. P(X >= x)         |\n")
  fmt.Print("| 4. P(X < x)          |\n") 
  fmt.Print("| 5. P(X <= x)         |\n")
  fmt.Print("| 6. P(x <= X <= y)    |\n")
  fmt.Print("| 7. P(x < X <= y)     |\n")
  fmt.Print("| 8. P(x <= X < y)     |\n")
  fmt.Print("| 9. P(x < X < y)      |\n")
  fmt.Print("|______________________|\n\n")
  fmt.Print("Escolha o tipo de probabilidade:\n")
  fmt.Scanf("%d", &op)
  
  //entrada de dados
  if op > 5 {
    fmt.Print("Escolha o valor de x, y e λ:\n")
    fmt.Scanf("%f\n",&x)
    fmt.Scanf("%f\n",&y)
    fmt.Scanf("%f",&lambda)
  }else{
    fmt.Print("Escolha o valor de x e λ:\n")
    fmt.Scanf("%f\n",&x)
    fmt.Scanf("%f",&lambda)
  }
  
  //calculo final
  switch (op){
    case 1: //P(X = x)
      result:=(math.Pow(e,-lambda)*math.Pow(lambda,x))/fatorial(x)
      fmt.Printf("%f\n", result)
    break

    case 2: //P(X > x) 
      for i := 0.0 ; i <= x ; i++ {
        result := (math.Pow(e,-lambda)*math.Pow(lambda,i))/fatorial(i)
        sum += result
      } 
      prob := 1-sum
      fmt.Printf("%f\n", prob);
    break

    case 3: //P(X >= x)
      for i := 0.0 ; i < x ; i++ {
        result := (math.Pow(e,-lambda)*math.Pow(lambda,i))/fatorial(i)
        sum += result
      } 
      prob := 1 - sum
      fmt.Printf("%f\n", prob)
    break     

    case 4: //P(X < x) 
      for i := 0.0 ; i < x ; i++ {
        result := (math.Pow(e,-lambda)*math.Pow(lambda,i))/fatorial(i)
        sum += result
      } 
      fmt.Printf("%f\n", sum)
    break  

    case 5: //P(X <= x) 
      for i := 0.0 ; i <= x ; i++ {
        result := (math.Pow(e,-lambda)*math.Pow(lambda,i))/fatorial(i)
        sum += result
      } 
      fmt.Printf("%f\n", sum)
    break    

    case 6: //P(x <= X <= y) 
      for i := x ; i <= y ; i++ {
        result := (math.Pow(e,-lambda)*math.Pow(lambda,i))/fatorial(i)
        sum += result
      } 
      fmt.Printf("%f\n", sum)
    break  

    case 7: //P(x < X <= y) 
      for i := x+1 ; i < y ; i++ {
        result := (math.Pow(e,-lambda)*math.Pow(lambda,i))/fatorial(i)
        sum += result
      } 
      fmt.Printf("%f\n", sum)
    break  

    case 8: //P(x <= X < y) 
      for i := x ; i < y ; i++ {
        result := (math.Pow(e,-lambda)*math.Pow(lambda,i))/fatorial(i)
        sum += result
      } 
      fmt.Printf("%f\n", sum)
    break

    case 9: //P(x < X < y) 
      for i := x+1 ; i <= y ; i++ {
        result := (math.Pow(e,-lambda)*math.Pow(lambda,i))/fatorial(i)
        sum += result
      } 
      fmt.Printf("%f\n", sum)
    break    
  
    default:
        fmt.Println("Error")
  }
}