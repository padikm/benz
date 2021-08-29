package converter

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"svc2/svc2/data"
)

func JSONtoCSV(e data.Emp) string {
	 var res string
	uuid := uuid.New()
	f, err := os.OpenFile("outputfile.csv",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	w := csv.NewWriter(f)
	var record []string
	if len(e.Id)==0 {
		res = uuid.String()
	}else {
		res = e.Id
	}
	record = append(record, res)
	record = append(record, e.Name)
	record = append(record, e.Address)
	record = append(record, strconv.Itoa(int(e.Age)))
	fmt.Println(record)
	err = w.Write(record)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Flush()
	return res
}


func CSVtoJson() []*data.Emp {
	f, err := os.Open("outputfile.csv")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)


	emp := []*data.Emp{}
	for {
		records, err := r.Read()
		log.Println(records)
		if err != nil {
			log.Println(err.Error())
			break
		}
	for i := 0; i < len(records); i += 4 {
		age, _ := strconv.Atoi(records[i+3])
		e := data.Emp{
			Id:      records[i],
			Name:    records[i+1],
			Age:     int32(age),
			Address: records[i+2],
		}
		emp = append(emp, &e)
	}
}
	return emp
}


func XMLConverter(s string) {

}
