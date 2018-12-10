package main

func main() {  
	x := 0  
	increment := func() int {        
		x++        
		return x    
		}    
		println(increment())    
		println(increment())   
}