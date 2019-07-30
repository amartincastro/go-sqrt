package main
 
import "fmt"
 
func main() {
    for i := 1; i <= 420; i++ {
        switch {
        case i%15==0:
            fmt.Println("DickButt")
        case i%3==0:
            fmt.Println("Dick")
        case i%5==0:
            fmt.Println("Butt")
        default: 
            fmt.Println(i)
        }
    }
}