package main

import "fmt"
import "time"

type Elevator struct {
	number           int
	floor            int
	direction        string
	status           string
	arrayDestination []int
}
func (e Elevator) move(floor int ){

	e.floor = floor;
	fmt.Println("elevator moving to : ", floor);
	time.Sleep(2 * time.Second)
}
func (e Elevator) stop(floor int) {
	fmt.Println("elevator stoping in : " , floor);
	time.Sleep(2 * time.Second)
}
func (e Elevator) openDoor(floor int){
	fmt.Println("elevator open the door " ,floor);
	time.Sleep(2 * time.Second)
}
func (e Elevator) closeDoor( destination int ){
   
	fmt.Println("elevator close the door ");
	time.Sleep(2 * time.Second)

}

type Button struct {
	number      int
	floor       int
	direction   string
	destination int
}

type Column struct {
	numbercolumn int
	firstfloor   int
	lastfloor    int
	elevator    []Elevator
}

////-------------------------------show column-------------------------
func (c Column) show() {

	fmt.Println("-------number column :", c.numbercolumn,"-------------")
	fmt.Println("firstfloor  :", c.firstfloor)
	fmt.Println("lastfloor  :", c.lastfloor)
	
	for i := 0; i < 5; i++ {
		fmt.Println("------- elevator number :  ",i,"  --------")
		fmt.Println("number  :", c.elevator[i].number)
		fmt.Println("floor  :", c.elevator[i].floor)
		
		fmt.Println("diriction  :", c.elevator[i].direction)
		fmt.Println("status :", c.elevator[i].status)

	}
	}
func (c Column) Max(t [5]int) int {
	if len(t)==0 {
	  return 0
	}
	max := t[0]
	for i:=0;i<len(t);i++ {
	  if t[i] > max {
		 max = t[i]
	  }
	}
	return max
  }
func (c Column) indexMax(t [5]int) int {
	if len(t)==0 {
	  return 0
	}
	max := t[0]
	maxindex := 0
	for i:=0;i<len(t);i++ {
	  if t[i] > max {
		 maxindex = i
	  }
	}
	return maxindex
  }
  //----------------------chalange distance------------------------
func (c Column) challengedistance(chalangeelevatorlist []Elevator, floornumber int) Elevator{
	 var poselevator1 int
	 var poselevator2 int
	 var bestpos int
	 elevatorswitch :=  Elevator{
		number :    -1,
		floor   :   0,
		direction  : "UP",
		status    :  "FREE",
	}
     
	 bestposelevator :=  Elevator{
		number :    -1,
		floor   :   0,
		direction  : "UP",
		status    :  "FREE",
	}
	
	for  i :=0; i < len(chalangeelevatorlist) - 1; i++ { 

		poselevator1 = floornumber - chalangeelevatorlist[i].floor;
		poselevator2 = floornumber - chalangeelevatorlist[i + 1].floor;

		if poselevator1 < 0 { poselevator1 = poselevator1 * (-1) }
		if poselevator2 < 0 { poselevator2 = poselevator2 * (-1) }
		 bestpos = poselevator2 - poselevator1;

		if bestpos > 0 {
		bestposelevator = chalangeelevatorlist[i];
		elevatorswitch = chalangeelevatorlist[i + 1];
		chalangeelevatorlist[i + 1] = chalangeelevatorlist[i];
		chalangeelevatorlist[i] = elevatorswitch;

		} else{
		bestposelevator = chalangeelevatorlist[i+1];
		}
	   

	}
	 
	return bestposelevator
} 

  //-----------------------chalange elevators---------------------------------------------------
func (c Column) chalangeelevator( direction string, floor int ) Elevator{
             ///-----------------VARIBLE LOCAL-----------------
	            var scoreElevator  [5]int
                bestlengthelevator :=  Elevator{
					number :    -1,
	                floor   :   0,
	                direction  : "UP",
	                status    :  "FREE",
				}
              
                bestposelevator :=  Elevator{
					number :    -1,
	                floor   :   0,
	                direction  : "UP",
	                status    :  "FREE",
				}
                 elevator :=  Elevator{
					number :    -1,
	                floor   :   0,
	                direction  : "UP",
	                status    :  "FREE",
				}
                var chalangeelevatorlist1 []Elevator
				var theelemeneselevators  []Elevator
               
				 //-------------------------------challenge elemenation -------------------------
				///----------------------------- IN THIS LOOP I WANT TO COMPLETE TWO TABLE THE FIRST WITH 
				///---------------------------THE ELEVATORS THAT THEY HAVE THE SAME DERICTION AS THE REQUEST AND THE SECOND WITH THE REST
				for i := 0; i < 5; i++ {
					 elevator = c.elevator[i]
					 
                    if elevator.direction == direction  {

                        if direction == "UP" && floor >= elevator.floor || direction == "DOWN" && floor <= elevator.floor {  
							chalangeelevatorlist1=append(chalangeelevatorlist1,elevator)
						
						 }
                        
                     }else { 
								  theelemeneselevators=append(theelemeneselevators,elevator)
								
						   }
				 }
				 
				 if len(chalangeelevatorlist1) == 1 {
					 
					  return c.elevator[chalangeelevatorlist1[0].number]; 
					}
				 /////////////////////////START A CHALLENGE WITH THE ELEVATOR LIST///////////////////////////////

                for  i := 0; i< len(chalangeelevatorlist1) - 1 ;i++ {

                    // '---------------------------------------challenge status if STATU OF ELEVATOR IS free ADD 1 POINT FOR HIM ------------------------------
                    if chalangeelevatorlist1[i].status == "free" { scoreElevator[chalangeelevatorlist1[i].number]++; }


                    //-----------------------------------chalange length--------------------------------------------------------///
                    //////////////////I WANT TO STORE THE ELEVATOR THAT IS THE SMALL REQUEST LIST
                    if len(chalangeelevatorlist1[i].arrayDestination) < len(chalangeelevatorlist1[i+1].arrayDestination) {
						
                        bestlengthelevator = chalangeelevatorlist1[i];
                    }

                        
                    
                  
                }	
				
                 //--------------------------------challange distance------------------------------------
				 bestposelevator = c.challengedistance(chalangeelevatorlist1, floor)
				
				 ////////////////////////ADD 1 POINT FOR THE ELEVATOR IN THE BEST POSSITION/////////////////////////////////
				 if bestposelevator.number != -1 { 
					
					 scoreElevator[bestposelevator.number]++
                    
					}
				 ////////////////////////ADD 1 POINT FOR THE ELEVATOR HE HAS A SMALL LIST/////////////////////////////////
				 if bestlengthelevator.number != -1 { 
					
					scoreElevator[bestlengthelevator.number]++ }
				 
 
				 
 
				 ////////////////////FIND THE  BEST SCOR /////////
				   var score = c.Max(scoreElevator)
				  
				   var numberelevator = c.indexMax(scoreElevator)
				  
				 ////////// IF THE SCOR == 0 IT IS TO SAY THAT THE FIRST LIST IS EMPTY. LAUNCH A CHELANGE WITH THE SECOND LIST
				   if score == 0 {
				
					return c.challengedistance(theelemeneselevators, floor);
				      }
				 return c.elevator[numberelevator]


}
//---------------
func (c Column) requestelevator( direction string, floornumber int ,  destination int) Elevator{
	 elevator := c.chalangeelevator(direction, floornumber);
	 fmt.Println("exactly in elevator number :   ",elevator.number);
	            time.Sleep(2 * time.Second)
                elevator.move(floornumber);
                elevator.stop(floornumber);
                elevator.openDoor(floornumber);
                elevator.closeDoor(destination);
	 return elevator
}
//------------------------------------------initialization column--------------------------------------
func initColumn( number int,first int,last int) Column{
	column := Column{
		numbercolumn: number,
		firstfloor:   first,
		lastfloor:    last,
	}
	

    for i := 0; i < 5; i++ {
	        elevator := Elevator{
			number:    i,
			floor:     column.firstfloor,
			direction: "UP",
			status:    "FREE",
		}

		column.elevator=append(column.elevator,elevator)
	}
	return column

}


func main() {
	column := Column{
		numbercolumn: 0,
		firstfloor:   2,
		lastfloor:    0,
	}
	

    for i := 0; i < 5; i++ {
	        elevator := Elevator{
			number:    i,
			floor:     column.firstfloor,
			direction: "UP",
			status:    "FREE",
		}

		column.elevator=append(column.elevator,elevator)
	}
	column1 := Column{
		numbercolumn: 0,
		firstfloor:   2,
		lastfloor:    0,
	}
	column2 := Column{
		numbercolumn: 0,
		firstfloor:   2,
		lastfloor:    0,
	}
	column3 := Column{
		numbercolumn: 0,
		firstfloor:   2,
		lastfloor:    0,
	}
	column4 := Column{
		numbercolumn: 0,
		firstfloor:   2,
		lastfloor:    0,
	}

    column1 = initColumn(4,-6,-1)
	column2 = initColumn(1,2,20)
	column3 = initColumn(2,21,40)
	column4 = initColumn(3,41,60)
	////////////////////////////////////INIT COLONNE NUMBER 2 FOR TESTED Scenario ///////////////////////////////////////
            ///
            //////////////////////////////////////Scenario 1:///////////////////////////////
            column2.elevator[4].floor = 6
            column2.elevator[3].floor = 15
            column2.elevator[2].floor = 13
            column2.elevator[1].floor = 3
            column2.elevator[0].floor = 20
           

            column2.elevator[4].direction = "DOWN"
            column2.elevator[3].direction = "DOWN"
            column2.elevator[2].direction = "DOWN"
            column2.elevator[1].direction = "UP"
			column2.elevator[0].direction = "DOWN"
			//////////////////////////////////////Scenario 2:///////////////////////////////
            column3.elevator[4].floor = 39
            column3.elevator[3].floor = 40
            column3.elevator[2].floor = 33
            column3.elevator[1].floor = 23
            column3.elevator[0].floor = 1
           

            column3.elevator[4].direction = "DOWN"
            column3.elevator[3].direction = "DOWN"
            column3.elevator[2].direction = "DOWN"
            column3.elevator[1].direction = "UP"
			column3.elevator[0].direction = "UP"
			//////////////////////////////////////Scenario 3:///////////////////////////////
            column4.elevator[4].floor = 60
            column4.elevator[3].floor = 1
            column4.elevator[2].floor = 46
            column4.elevator[1].floor = 50
            column4.elevator[0].floor = 58
           

            column4.elevator[4].direction = "DOWN"
            column4.elevator[3].direction = "UP"
            column4.elevator[2].direction = "UP"
            column4.elevator[1].direction = "UP"
			column4.elevator[0].direction = "DOWN"
			//////////////////////////////////////Scenario 3:///////////////////////////////
            column1.elevator[4].floor = -4
            column1.elevator[3].floor = 1
            column1.elevator[2].floor = -3
            column1.elevator[1].floor = 1
            column1.elevator[0].floor = -4
           

            column1.elevator[4].direction = "DOWN"
            column1.elevator[3].direction = "UP"
            column1.elevator[2].direction = "DOWN"
            column1.elevator[1].direction = ""
            column1.elevator[0].direction = ""
   
	column2.show()
	column3.show()
	column4.show()
	var battrie = "on"
	var floorcall int
	var destinationcall int
	for battrie == "on" {
		fmt.Println("  //////////////////////////////////////////////")
		fmt.Println(" ///////////enter the current floor ???////////")
		fmt.Println("//////////////////////////////////////////////")
		fmt.Scan(&floorcall)
		if floorcall < 0 { floorcall = floorcall + 67 }
		fmt.Println("  ///////////////////////////////////////////")
		fmt.Println(" ////enter floor where you want to go ?? ///")
		fmt.Println("///////////////////////////////////////////")
		fmt.Scan(&destinationcall)
		if (destinationcall < 0) { destinationcall = destinationcall + 67;}
		callbutton := Button {
			number : 0,
			floor: floorcall,
			direction :"DOWN",
			destination : destinationcall,
			 } 
		if floorcall < destinationcall { callbutton.direction = "UP"; }
		if callbutton.destination >= 61 && callbutton.destination <= 66 && callbutton.floor >= 61 && callbutton.floor <= 66 || callbutton.floor >= 61 && callbutton.floor <= 66 && callbutton.destination==1 || callbutton.floor >= 61 && callbutton.floor <= 66 && callbutton.floor == 1 {elevator1 := column1.requestelevator( callbutton.direction , callbutton.floor ,  callbutton.destination) 
		fmt.Println("your elevator is going to be at the column 1  exactly in elevator number :  ",elevator1.number)  
	       }
		if callbutton.destination >= 2 && callbutton.destination <= 20 && callbutton.floor >= 2 && callbutton.floor <= 20 || callbutton.floor >= 2 && callbutton.floor <= 20 && callbutton.destination == 1 || callbutton.destination >= 2 && callbutton.destination <= 20 && callbutton.floor == 1 { elevator1 := column2.requestelevator( callbutton.direction , callbutton.floor ,  callbutton.destination ) 
			fmt.Println("your elevator is going to be at the column 1  exactly in elevator number :  ",elevator1.number)  }
		if callbutton.destination >= 21 && callbutton.destination <= 40 && callbutton.floor >= 21 && callbutton.floor <= 40 || callbutton.floor >= 21 && callbutton.floor <= 40 && callbutton.destination == 1 || callbutton.destination >= 21 && callbutton.destination <= 40 && callbutton.floor == 1 {  elevator1 := column3.requestelevator( callbutton.direction , callbutton.floor ,  callbutton.destination ) 
			fmt.Println("your elevator is going to be at the column 1  exactly in elevator number :  ",elevator1.number)   }
		if callbutton.destination >= 41 && callbutton.destination <= 60 && callbutton.floor >= 41 && callbutton.floor <= 60 || callbutton.floor >= 41 && callbutton.floor <= 60 && callbutton.destination == 1 || callbutton.destination >= 41 && callbutton.destination <= 60 && callbutton.floor == 1 { elevator1 := column4.requestelevator( callbutton.direction , callbutton.floor ,  callbutton.destination ) 
			fmt.Println("your elevator is going to be at the column 1  exactly in elevator number :  ",elevator1.number)   }
	  
	}
	
}
