package main

import "fmt"

func main(){
	//Hello World
	fmt.Println("Hello World")

	//Variabel

	   //***Manifest Typing***
	   var firstName string = "Author"
	   var lastName string
	   lastName = "Afryancee"
	   fmt.Println(firstName,lastName+" "+"By Manifest Typing")

	   //***Multi Variable***
	   var one,two,three string
	   one,two,three = "1","2","3"
	   fmt.Println(one,two,three);
	 
	//end

	//Konstanta
	   const firstNamee string = "User"
	   fmt.Println("cons : ",firstNamee)
	//end
	
	//Operator
	  var value = (((2+3)+5)) / 5
	  var equal = (value == 2)
	  fmt.Printf("nilai %d (%t) \n",value,equal)

	  //***Operator Logika
	  var left  = false
	  var right = true
	  var leftAndRight = left && right
	  fmt.Printf("And : \t(%t) \n",leftAndRight)
	  var leftOrRight = left || right
	  fmt.Printf("Or : \t(%t) \n",leftOrRight)
	  var leftReverse = !left
	  fmt.Printf("Reverse : \t(%t) \n",leftReverse)
	//
	
	//If Else
	  var point int = 8
      if point == 10 {
		fmt.Println("5")
      }else if point == 8{
        fmt.Println("4")
	  }else if point == 5{
		fmt.Println("3")
	  }else{
		fmt.Println("Are You Human")
	  }

	  //if-case
	  switch{
	  case point == 8:
		  fmt.Println("perfect")
	      case(point < 8 )  && (point > 3):
			  fmt.Println("Awesome")
		  default:{
			  fmt.Println("Not Bad")
			  fmt.Println("Huuhhhh")
		  }
		}
		
		//if-Array
		var mnb = 10
		if mnb > 5{
			switch point{
			case 10:
					fmt.Println("Perfect")
			default:
				fmt.Println("No Effect")
			}
		}else{
           if mnb == 7 {
               fmt.Print("A")
		   }else if mnb == 3{
              fmt.Print("keep trying")
		   }else{
			  fmt.Println("Fuck With You")
			  if point == 0 {
				  fmt.Print("Try harder")
			  }
		   }
		}
	//end

	//foreach 
	   //***for-keyword
	   for i := 0; i < 5; i++{
		   fmt.Println("Number",i)
	   }

	   //***for-array
	   for i := 0; i < 5; i++{
		   for j :=i; j < 5; j++{
			   fmt.Print(j," ")
		   }
		   fmt.Println()
	   }
	//end
	
}