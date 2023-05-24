package main
import "fmt"

func main() {
  var (
    numero int
    calc uint64 = 1
 )

  fmt.Print("Você deseja calcular o fatorial de qual numero? ")
  fmt.Scan(&numero)

  for i := numero; i > 0; i-- {
    calc *= uint64(i)
  } 

  fmt.Println("Pronto:", calc)
}
