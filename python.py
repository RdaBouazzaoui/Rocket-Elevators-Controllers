


import time


class elevator():
           
       
       def __init__(self, number,floor,direction,status):      
            self.number = number
            self.floor=floor
            self.direction = direction
            self.status = status
            self.arrayDestination = []

         #//////////////////////method addbyorder elevator destination  ////////////////////
       def addByOrder(self,callbutton):
                                
             
                      self.arrayDestination.append(callbutton)
                      
                      sorted(self.arrayDestination, key=lambda callbutton: callbutton.floor)    
                      if self.direction == "DOWN" :
                                 self.arrayDestination.reverse()

                               
                           
                     



         
       def refreshElevator(self):
                del self.arrayDestination[0]

       def move(self) :
          self.floor = self.arrayDestination[0].floor
          print("elevator move to "+str(self.arrayDestination[0].floor) )
          time.sleep(2)

       def openDoor(self) : 
         
         print("elevator open the door  ")
         time.sleep(2)
          
       
       def stop(self):
           print("elevator stoped in "+str(self.arrayDestination[0].floor))
           time.sleep(2)
          
       def closeDoor(self):
         if self.arrayDestination[0].destination >= 0 :
                self.arrayDestination[0].floor=self.arrayDestination[0].destination
                self.arrayDestination.append(self.arrayDestination[0])
         print("elevator close the door ")
         time.sleep(2)
         self.refreshElevator()
          
          
          




class button():
      
        def __init__(self,floor,direction,destination): 
           self.floor = floor 
           self.direction = direction
           self.destination = destination
         
         





def RequestFloor(elevator):
       elevator.move()
       elevator.stop()
       elevator.openDoor()
       elevator.closeDoor()

def  bestPosition(callButton) :
    poselevator1 = callButton-elevator1.floor

    poselevator2 = callButton-elevator2.floor

    if poselevator1 < 0 : 
         poselevator1 = poselevator1* (-1)

    if poselevator2 < 0 : 
            poselevator2 = poselevator2* (-1)

    pos = poselevator2 - poselevator1

    if pos > 0 :
          return 2
    else:
           return 1

def chickCallButton():
    
           callbutton=callbuttonList[0]
      
           return callbutton


def RequestElevator(direction , floor):

     #////////////callbutton.text this is the request direction/////////////
           scoreElevator1=0
           scoreElevator2=0      
           if direction=="UP" :
                 
                if elevator1.direction == direction and elevator2.direction == direction and floor >= elevator1.floor and floor >= elevator2.floor :
                          
                    
                 #//////////////////////////////chalange with two elevator///////////////
                     best = bestPosition(floor) 
                     if  best == 1 : 

                             scoreElevator1 = scoreElevator1 +1
                            

                     else : 
                            
                             scoreElevator2 = scoreElevator2 +1
                          
                          
                          #//the lenght OF elevator1.ArrayDestination < the lenght of elevator2.ArrayDestination
                     if  len(elevator1.arrayDestination) < len(elevator2.arrayDestination) :
                                
                           scoreElevator1 = scoreElevator1 +1
                           

                     else : 
                            
                           scoreElevator2 = scoreElevator2 +1
                          
                             

                           
                             

                     if scoreElevator1 > scoreElevator2 :
                          

                            return elevator1
                          
                             
                     else :
                           
                             return elevator2
                          
                        

                         

                    
                      
                   

                if elevator1.direction == direction and  floor >= elevator1.floor :

                              return elevator1
                        

                if elevator2.direction == direction and floor >= elevator2.floor :

                              return elevator2
                if elevator1.status== "free" :
                      elevator1.status = "busy"
                      elevator1.direction = direction
                      return elevator1
                else :
                     if elevator1.status== "free" :
                         elevator2.status = "busy"
                         elevator2.direction = direction
                         return elevator2

            
                          
                       

                      
           if direction=="DOWN" :
                 
                if elevator1.direction ==direction and elevator2.direction == direction and floor <= elevator1.floor and floor <= elevator2.floor :
                          
                     scoreElevator1
                     scoreElevator2
                 #//////////////////////////////chalange with two elevator///////////////

                     if  bestPosition(floor) == 1 : 

                             scoreElevator1 = scoreElevator1 +1
                            

                     else : 
                            
                             scoreElevator2 = scoreElevator2 +1
                          
                          
                          #//the lenght OF elevator1.ArrayDestination < the lenght of elevator2.ArrayDestination
                     if  len(elevator1.arrayDestination)  <  len(elevator2.arrayDestination)  :
                                
                           scoreElevator1 = scoreElevator1 +1
                           

                     else : 
                            
                           scoreElevator2 = scoreElevator2 +1
                          
                             

                           
                             

                     if scoreElevator1 > scoreElevator2 :
                          

                            return elevator1
                          
                             
                     else :
                           
                             return elevator2
                          
                        

                         

                    
                      
                   

                if elevator1.direction == direction and  floor <= elevator1.floor :

                              return elevator1
                        

                if elevator2.direction == direction and floor <= elevator2.floor :

                              return elevator2

                if elevator1.status== "free" :
                      
                      elevator1.direction = direction
                      return elevator1
                else :
                     if elevator1.status== "free" :
                         
                         elevator2.direction = direction
                         return elevator2


          
        


elevator1 = elevator(1,10,"DOWN","free")
elevator2 = elevator(2,3,"UP","free")
callbutton1 = button(1,"UP",6)
callbutton2 = button(3,"up",5)
callbutton3 = button(9,"DOWN",2)
callbutton4 = button(6,"DOWN",1)
callbutton5 = button(8,"UP",10)
callbutton6 = button(9,"UP",10)
callbuttonList = [callbutton1,callbutton2,callbutton3]
print("call button list :")
for x in callbuttonList :
  print("callbutton direction  "+x.direction+ "    callbutton position    "+str(x.floor) )


i=0
j=0
while j< len(callbuttonList) :
  
   callb=chickCallButton()
   print("///////////////////call button "+str(i)+" /////////////")
   print("callbutton direction :  "+callb.direction+ "  callbutton position :  "+str(callb.floor))
   print("////////////////////////current elevator position ///////////////")
   print("elevator1 position  :  "+ str(elevator1.floor) +"   elevator1 direction : " +elevator1.direction)
   print("elevator2 position  :  "+str(elevator2.floor) +"  elevator2 direction :  " +elevator2.direction)
   elevator = RequestElevator(callb.direction,callb.floor)
   if elevator != False :
       del callbuttonList[0]
       print(elevator)
       

       elevator.addByOrder(callb)
      

       print("elevator finded to do that request is elevator:" + str(elevator.number))
       RequestFloor(elevator)
   else : 
      print("the request will be processed later")
      callbu=callbuttonList[0]
      del callbuttonList[0]
      callbuttonList.append(callbu)

   
   i=i+1

print("call button list :")
for x in callbuttonList :
    print("callbutton direction  "+x.direction+ "    callbutton position    "+str(x.floor) ) 



