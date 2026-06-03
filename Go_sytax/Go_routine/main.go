//simulating downlading files. using go routines,.
package main
import (
	"fmt" //priting things
	"sync" //synchroization tools. 
	"time" //sleeping and working wiht time. 

)

func download(name string, wg *sync.WaitGroup){
	defer wg.Done() // when this functions ends calls this

	fmt.Println("downloading", name)
	time.Sleep(2 * time.Second) // pretending that a download takes 2 seconds

	fmt.Println("finished", name) // simulating after download text

}

func main(){
	var wg sync.WaitGroup // creates wait group conter, 

	files := []string {
		"files1.zip",
		"files2.zip",
		"files3.zip",
	}

	for _,file := range files {
		wg.Add(1)
		go download(file, &wg)
	}
	wg.Wait()
}