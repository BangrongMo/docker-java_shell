package main
import ("fmt"
		"time"
		"sync"
)
var wg sync.WaitGroup

func mybuffer(number int,buffer chan int){
		for i:=3;i<=number;i++{
			buffer <- i
			}
			close(buffer)
			wg.Done()
	}

func consume(buffer,product chan int, exitChan chan bool){
		for  num:=range buffer{
			var flag =true
			for i:=2;i<=num/2;i++{
				
				if num%i==0{
					flag =false
					break
					}
		
				
				}
				if flag{
					product <- num
				}
				
			}
			exitChan <- true //
			wg.Done()
	}
func disProduct(product chan int){
	for v:=range product{
		 fmt.Println(v)
		}
		wg.Done()
}

func main(){
	var threadCount = 4
	startTime :=time.Now().Unix()
	var buffer =make(chan int,4096)
	var product =make(chan int,4096)
	var exitChan =make(chan bool ,threadCount)
	wg.Add(3)
	go mybuffer(120000,buffer)
	go disProduct(product)
	for i:=1;i<=threadCount;i++{
		wg.Add(1)
		go consume(buffer ,product,exitChan)
		}
	go func (){
		for i:=1;i<=threadCount;i++{
			fmt.Println(<-exitChan)
			}
			close(product)
			wg.Done()
			
		}()
	wg.Wait()
	
	endTime :=time.Now().Unix()
	fmt.Println(endTime-startTime)
	}
