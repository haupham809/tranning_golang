package writelog
import (
	"log"
	"os"
	"tranning_golang/model/messageapi"
	"math/rand"
	"fmt"
	"time"
	"strconv"
)
func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}
func Writelog(o messageapi.Objectapi){
	
	filename :=  strconv.Itoa(rand.Intn(100000000)) +strconv.FormatInt(makeTimestamp(),10)
	fmt.Println(filename)
	file, _  := os.OpenFile("../log/"+filename+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.Println( o)
}