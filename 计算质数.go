package main
//2025.3.14
import ("fmt"
		"time"
		"sync"
		)
	
var wg sync.WaitGroup
func findPrime(start,end int){
	for num :=start;num<end;num++{
		flag :=true
		for i:=2;i<num;i++{
			if i>(num/2){
					break
				} 
				
			if num%i == 0{ 
				flag =false
				break
				}
			}
		 if flag {
			// fmt.Println(num)
			} 
		}
	}
		
func divideData(x,y int){
	
	findPrime(x,x+20000) //0-2,2 -4,4-6,
	findPrime(120000-20000-y,120000-y)//10-12,8-10,6-8
	fmt.Printf("divideData(%v,%v) done : %v \n",x,y,time.Now().Unix())
	
	wg.Done()
	
	}
		
func main(){
	start :=time.Now().Unix()
	fmt.Println(1)
	fmt.Println(2)
	wg.Add(3)
	go divideData(0,0)
	go divideData(20000,20000)
	go divideData(40000,40000)
	wg.Wait() 
	/* findPrime(3,120000)*/
	end :=time.Now().Unix()
	fmt.Println("time:",end - start,"s") 
	
	}
