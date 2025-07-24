//---------simulate the code----------


package main

import "fmt"

func hof(a int, b int, op func(p int, q int)) {     // hof = higher order function
	op(a, b)
}

func add(x int, y int) {
	var z = x + y
	fmt.Println(z)
}

func main(){
	hof(2, 5, add)
}

//----------------------------------------------------

/*

  1. Arguments vs Parameters
      -> age value pass kori (arguments)
      -> pore value receive kori (parameters)

  2. First oder function
     i) name ache -> tai named or Standard Function
	 ii) name nai -> Anonymus function (name na thakai call kora jaina, tai ei function error dekabe, so reke deya jabena, but javascript e reke deya jai)
	 iii) Anonymus function k immidiately call (name chara call) krle -> IIFE
	 iv) Assign function in variable (Anonymus function k var er vetor rakha hoi, erpr oi var er nam dhore call kora jai)

  3. Higher order function (j function parameter hisebe function receive kore or return kore or 2 tai kore seta HOF)
  4. callback function (HOF j function k pass kora hoi setai callback function)
  5. first class citizen (variable e j data assign kora hoi setai first class citizen)
       -> HOF = first class funtion (karon HOF er vetor function k pass kora jai (onekta assign korar mto))


  math => logic (discrete mathematic)
     i. First Order Logic
	 ii. Higher Order Logic

	 ###First Order logic###
	 ei 3 ta boisisto nie kaj kore
	 1. Objects (people, animal, car etc.)
	 2. Properties (color, student)
	 3. Relation
     -->
		'Safwan' is a 'student'
		'Safwan' is 'taller than Ashraf'
		'All customers' must pay pizza bills. (Rule)
		'All students' must wear their uniform. (Rule)

	 ###Higher Order logic###
	 Rules/logic nie kaj kore, sate first order logic er boisisto gulo o ache
	 -->
	 	Any rule that applies to all customers must also apply to Safwan (Rule)


  functional paradigm er language -> haskel, raket
  functional paradigm e ache 'first order function' & 'last order funtion'

  functional paradigm influenced from math's logic
  -> go influenced from functionl paradigm
*/