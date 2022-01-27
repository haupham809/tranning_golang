package writelog

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"tranning_golang/model/messageapi"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
func Writelog(o messageapi.Objectapi) {

	// filename :=  strconv.Itoa(rand.Intn(100000000)) +strconv.FormatInt(makeTimestamp(),10)
	var fromDate time.Time
	// day := fromDate.Format()
	filename := "logs " + strconv.Itoa(fromDate.Local().Day())
	fmt.Println(filename)
	file, _ := os.OpenFile("../log/"+filename+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.Println(o)
}
