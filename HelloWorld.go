
'''
just simple code for print "Hello word in go"
'''
package main 
import "fmt"


//declare  variabel hello 
const hello string = "Hello"


//defisinsian functiion tambahan World
func tambahWorld(word string)string{
	return word + " World"
}

//function main
func main(){
	var helloWorld = tambahWorld(hello)
	fmt.Println(helloWorld)
}