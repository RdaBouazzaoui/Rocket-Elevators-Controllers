using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;

namespace myFirstProject
{


    public class Program
    {
        public class Callbutton
        {
            
            public int floor;
            public string  direction;
            public int destination;
            public Callbutton() { }
            public Callbutton(int floor , string direction, int destination) {

                this.floor = floor;
                this.direction = direction;
                this.destination = destination;
            }



        }
        
        public class Elevator
        {
            public int number ;
            public int floor;
            public string direction;
            public Callbutton[] destinatonlist=new Callbutton[2];
            public string status;

            public Elevator()

            {
                


            }
            public Elevator(int number , int floor,string direction,string status)

            {
                this.floor = floor;
                this.direction = direction;
                this.status = status;
                this.number = number;


            }
            public void move(int floor)
            {

                this.floor = floor;
                Console.WriteLine("elevator moving to : " + floor);
                Thread.Sleep(4000);
            }
            public void stop(int floor) 
            {
                Console.WriteLine("elevator stoping in : " + floor);
                Thread.Sleep(4000);
            }
            public void openDoor(int floor)
            {
                Console.WriteLine("elevator open the door " + floor);
                Thread.Sleep(4000);
            }
            public void closeDoor(int destination)
            {
               
                Console.WriteLine("elevator close the door ");
                Thread.Sleep(4000);

            }

        }
        public class Colonne
        {
            

            public int numbercolonne;
            public int firstfloor;
            public int lastfloor;
            public Elevator[] elevators= new Elevator[5];
            public Colonne(int numbercolonne, int firstfloor, int lastfloor)
            {
                this.numbercolonne = numbercolonne;
                this.firstfloor = firstfloor;
                this.lastfloor = lastfloor;

                for (int i = 0; i < 5; i++)
                {
                    Elevator elevator = new Elevator(i,firstfloor, "UP", "FREE");

                    this.elevators[i] = elevator;

                }
            }
            private Elevator challengedistance(Elevator[] elevator,int floornumber)
            {
                int poselevator1;
                int poselevator2;
                Elevator elevatorchenge = new Elevator();
                int bestpos;
                Elevator bestposelevator = new Elevator();
                bestposelevator.number = -1;
                for (int i=0; i < elevator.Length - 1; i++) { 

                    poselevator1 = floornumber - elevator[i].floor;
                    poselevator2 = floornumber - elevator[i + 1].floor;

                    if (poselevator1 < 0) { poselevator1 = poselevator1 * (-1); }
                    if (poselevator2 < 0) { poselevator2 = poselevator2 * (-1); }
                     bestpos = poselevator2 - poselevator1;

                    if (bestpos > 0)
                    {
                    bestposelevator = elevator[i];
                    elevatorchenge = elevator[i + 1];
                    elevator[i + 1] = elevator[i];
                    elevator[i] = elevatorchenge;

                    }
                    else
                    {
                    bestposelevator = elevator[i+1];
                    }
                   

                }
                 
                return bestposelevator;
            }
            private  Elevator chalangeelevator(string direction,int floornumber )
            {
                int[] scoreElevator = new int[5];
                Elevator bestlengthelevator = new Elevator();
                bestlengthelevator.number = -1;
                Elevator bestposelevator = new Elevator();
                Elevator elevator = new Elevator();
                List<Elevator> chalangeelevatorlist1 = new List<Elevator>();
                List<Elevator> theelemeneselevators = new List<Elevator>();


                //-------------------------------challenge elemenation -------------------------
                ///----------------------------- IN THIS LOOP I WANT TO COMPLETE TWO TABLE THE FIRST WITH THE ELEVATORS THAT THEY HAVE THE SAME DERICTION AS THE REQUEST AND THE SECOND WITH THE REST---------------------------------
                for (int i = 0; i < 5; i++) 
                {
                  
                    if (this.elevators[i].direction == direction  ) 
                    {

                        if (direction == "UP" && floornumber >= this.elevators[i].floor || direction == "DOWN" && floornumber <= this.elevators[i].floor) 
                        {  chalangeelevatorlist1.Add(this.elevators[i]); }
                        
                     }else { theelemeneselevators.Add(this.elevators[i]); }
                 }
               
                
                
                     Elevator[] chalangeelevatorlist = chalangeelevatorlist1.ToArray();

                
                
                if (chalangeelevatorlist.Length == 1) { return this.elevators[chalangeelevatorlist[0].number]; }

                /////////////////////////START A CHALLENGE WITH THE ELEVATOR LIST///////////////////////////////

                for (int i=0; i< chalangeelevatorlist.Length - 1 ;i++) 
                         {

                    // '---------------------------------------challenge status if STATU OF ELEVATOR IS free ADD 1 POINT FOR HIM ------------------------------
                    if (chalangeelevatorlist[i].status == "free") { scoreElevator[chalangeelevatorlist[i].number]++; }


                    //-----------------------------------chalange length--------------------------------------------------------///
                    //////////////////I WANT TO STORE THE ELEVATOR THAT IS THE SMALL REQUEST LIST
                    if (chalangeelevatorlist[i].destinatonlist.Length < chalangeelevatorlist[i+1].destinatonlist.Length)
                    {
                        bestlengthelevator = chalangeelevatorlist[i];
                    }

                        
                    
                  
                }

               //--------------------------------challange distance------------------------------------
                bestposelevator = challengedistance(chalangeelevatorlist, floornumber);
                ////////////////////////ADD 1 POINT FOR THE ELEVATOR IN THE BEST POSSITION/////////////////////////////////
                if (bestposelevator.number != -1) { scoreElevator[bestposelevator.number]++;}
                ////////////////////////ADD 1 POINT FOR THE ELEVATOR HE HAS A SMALL LIST/////////////////////////////////
                if (bestlengthelevator.number != -1) { scoreElevator[bestlengthelevator.number]++; }
                

                

                ////////////////////FIND THE  BEST SCOR /////////
                int score = scoreElevator.Max();
                int numberelevator = scoreElevator.ToList().IndexOf(score);
                ////////// IF THE SCOR == 0 IT IS TO SAY THAT THE FIRST LIST IS EMPTY. LAUNCH A CHELANGE WITH THE SECOND LIST
                if (score == 0)
                {
                    Elevator[] theelemeneselevators1 = theelemeneselevators.ToArray();
                    
                    return challengedistance(theelemeneselevators1, floornumber);
                }
                return this.elevators[numberelevator];
            


            }
            private Elevator searchelevator(string direction, int floornumber)
            {
                return chalangeelevator(direction, floornumber);
                            
                
                


            }
            public Elevator requestelevator(string direction, int floornumber, int destination)
            {
                Elevator elevator = searchelevator(direction, floornumber);
                Console.WriteLine("exactly in elevator number :   " + elevator.number);
                Thread.Sleep(4000);
                elevator.move(floornumber);
                elevator.stop(floornumber);
                elevator.openDoor(floornumber);
                elevator.closeDoor(destination);

                return elevator;

            }

        }


            
        
      

        static void Main(string[] args)
        {
            string battrie = "on";
            
            Colonne colonne1 = new Colonne(1,-6,-1);
            Colonne colonne2 = new Colonne(2,2,20);
            Colonne colonne3 = new Colonne(3, 21, 40);
            Colonne colonne4 = new Colonne(4, 41,60);
            
            /**/
            ////////////////////////////////////INIT COLONNE NUMBER 2 FOR TESTED Scenario ///////////////////////////////////////
            ///
            //////////////////////////////////////Scenario 1:///////////////////////////////
            colonne2.elevators[4].floor = 6;
            colonne2.elevators[3].floor = 15;
            colonne2.elevators[2].floor = 13;
            colonne2.elevators[1].floor = 3;
            colonne2.elevators[0].floor = 20;
           

            colonne2.elevators[4].direction = "DOWN";
            colonne2.elevators[3].direction = "DOWN";
            colonne2.elevators[2].direction = "DOWN";
            colonne2.elevators[1].direction = "UP";
            colonne2.elevators[0].direction = "DOWN";
            ////////////////////////////////////INIT COLONNE NUMBER 2 FOR TESTED Scenario ///////////////////////////////////////
            //////////////////////////////////////Scenario 2:///////////////////////////////
            ///

            colonne3.elevators[4].floor = 39;
            colonne3.elevators[3].floor = 40;
            colonne3.elevators[2].floor = 33;
            colonne3.elevators[1].floor = 23;
            colonne3.elevators[0].floor = 1;
           
            colonne3.elevators[4].direction = "DOWN";
            colonne3.elevators[3].direction = "DOWN";
            colonne3.elevators[2].direction = "DOWN";
            colonne3.elevators[1].direction = "UP";
            colonne3.elevators[0].direction = "UP";
            ////////////////////////////////////INIT COLONNE NUMBER 2 FOR TESTED Scenario ///////////////////////////////////////
            //////////////////////////////////////Scenario 3:///////////////////////////////
            colonne4.elevators[4].floor = 60;
            colonne4.elevators[3].floor = 1;
            colonne4.elevators[2].floor = 46;
            colonne4.elevators[1].floor = 50;
            colonne4.elevators[0].floor = 58;

            colonne4.elevators[4].direction = "DOWN";
            colonne4.elevators[3].direction = "UP";
            colonne4.elevators[2].direction = "UP";
            colonne4.elevators[1].direction = "UP";
            colonne4.elevators[0].direction = "DOWN";
            ////////////////////////////////////INIT COLONNE NUMBER 2 FOR TESTED Scenario ///////////////////////////////////////
            //////////////////////////////////////Scenario 4:///////////////////////////////
            colonne1.elevators[4].floor = -4;
            colonne1.elevators[3].floor = 1;
            colonne1.elevators[2].floor = -3;
            colonne1.elevators[1].floor = 1;
            colonne1.elevators[0].floor = -4;

            colonne1.elevators[4].direction = "DOWN";
            colonne1.elevators[3].direction = "UP";
            colonne1.elevators[2].direction = "DOWN";
            colonne1.elevators[1].direction = "";
            colonne1.elevators[0].direction = "";



            Colonne[] colonnes = { colonne1, colonne2, colonne3, colonne4 };
//////////////////////////////////////////SHOW ALL INFORMATION ABOUT COLUMN///////////////////////////////////////////////////////////
            for (int i = 0; i < 4; i++)
            {
                
                Console.WriteLine("colonne :  "+colonnes[i].numbercolonne );
                Console.WriteLine("colonne first floor :  " + colonnes[i].firstfloor);
                Console.WriteLine("colonne last floor :  " + colonnes[i].lastfloor);
                for (int j = 0; j < 5; j++)
                {
                    Console.WriteLine("elevator numbre :  " + j);

                    Console.WriteLine("elevator floor :  " + colonnes[i].elevators[j].floor);
                    Console.WriteLine("elevator distination :  " + colonnes[i].elevators[j].direction);
                    Console.WriteLine("elevator status  :  " + colonnes[i].elevators[j].status);




                }

            }


           
            int floor;
            int destination;
            
            Elevator elevator=new Elevator();

            while (battrie == "on")
            {
                Console.WriteLine("  //////////////////////////////////////////////");
                Console.WriteLine(" ///////////enter the current floor ???////////");
                Console.WriteLine("//////////////////////////////////////////////");
                floor = int.Parse(Console.ReadLine());
              
                Console.WriteLine("  ///////////////////////////////////////////");
                Console.WriteLine(" ////enter floor where you want to go ?? ///");
                Console.WriteLine("///////////////////////////////////////////");
                destination = int.Parse(Console.ReadLine());
                
                Callbutton callbutton = new Callbutton(floor, "DOWN", destination); 
                if (floor < destination) { callbutton.direction = "UP"; }
                if (callbutton.destination >= -6 && callbutton.destination <= -1 && callbutton.floor >= -6 && callbutton.floor <= -1 || callbutton.floor >= -6 && callbutton.floor <= -1 && callbutton.destination==1 || callbutton.floor >= -6 && callbutton.floor <= -1 && callbutton.floor == 1) { Console.WriteLine("your elevator is going to be at the column 1 ");  elevator = colonne1.requestelevator(callbutton.direction, callbutton.destination, callbutton.floor);  }
                if (callbutton.destination >= 2 && callbutton.destination <= 20 && callbutton.floor >= 2 && callbutton.floor <= 20 || callbutton.floor >= 2 && callbutton.floor <= 20 && callbutton.destination == 1 || callbutton.destination >= 2 && callbutton.destination <= 20 && callbutton.floor == 1) { Console.WriteLine("your elevator is going to be at the column 2 "); elevator = colonne2.requestelevator(callbutton.direction, callbutton.floor, callbutton.destination); }
                if (callbutton.destination >= 21 && callbutton.destination <= 40 && callbutton.floor >= 21 && callbutton.floor <= 40 || callbutton.floor >= 21 && callbutton.floor <= 40 && callbutton.destination == 1 || callbutton.destination >= 21 && callbutton.destination <= 40 && callbutton.floor == 1) { Console.WriteLine("your elevator is going to be at the column 3 "); elevator = colonne3.requestelevator(callbutton.direction, callbutton.floor,callbutton.destination);  }
                if (callbutton.destination >= 41 && callbutton.destination <= 60 && callbutton.floor >= 41 && callbutton.floor <= 60 || callbutton.floor >= 41 && callbutton.floor <= 60 && callbutton.destination == 1 || callbutton.destination >= 41 && callbutton.destination <= 60 && callbutton.floor == 1) { Console.WriteLine("your elevator is going to be at the column 4 "); elevator = colonne4.requestelevator(callbutton.direction, callbutton.floor, callbutton.destination); }
              
              


            




            }


        }
    }
}
