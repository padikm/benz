package converter

import (
	"bufio"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"svc2/svc2/data"
)

type Converter struct {
	CSVfile string
	Xmlfile string
}

func (c Converter) JSONtoCSV(e data.Emp) (string,error) {
	var res string
	uuid := uuid.New()
	f, err := os.OpenFile(c.CSVfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err.Error() + c.CSVfile)
		return "",err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	var record []string
	if len(e.Id) == 0 {
		res = uuid.String()
	} else {
		res = e.Id
	}
	record = append(record, res)
	record = append(record, e.Name)
	record = append(record, e.Address)
	record = append(record, strconv.Itoa(int(e.Age)))
	fmt.Println(record)
	err = w.Write(record)
	if err != nil {
		log.Println(err.Error())
		return "",err
	}
	w.Flush()
	return res,nil
}

func (c Converter) CSVtoJson() ([]*data.Emp, error) {
	f, err := os.Open(c.CSVfile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
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
	return emp, nil
}

func (c Converter) JSONtoXML(e data.Emp) (string,error) {
	var res string
	uuid := uuid.New()
	type Emp struct {
		Name    string `xml:"Name" json:"name"`
		Age     int    `xml:"Age" json:"age"`
		Address string `xml:"Address" json:"address"`
		Id      string `xml:"Id" json:"id"`
	}
	if len(e.Id) == 0 {
		res = uuid.String()
	} else {
		res = e.Id
	}
	e.Id = res
	b, _ := xml.Marshal(&e)
	log.Println(string(b))

	f, err := os.OpenFile(c.Xmlfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err.Error())
		return "",err
	}
	defer f.Close()

	_, err = f.WriteString(string(b)+"\n")
	if err!=nil {
		return "",err
	}
	return res,nil
}
type Emp struct {
	Name    string `xml:"Name" json:"name"`
	Age     int32    `xml:"Age" json:"age"`
	Address string `xml:"Address" json:"address"`
	Id      string `xml:"Id" json:"id"`
}
func (c Converter) XMLtoJSON() ([]Emp,error) {

	emps := []Emp{}
	xmlFile, err := os.Open(c.Xmlfile)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
		return nil,err

	}
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	//decoder := xml.NewDecoder(xmlFile)
	//decoder.Decode(emps)
	scanner := bufio.NewScanner(xmlFile)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		scanner.Text()
		emp := Emp{}
		err = xml.Unmarshal([]byte(scanner.Text()),&emp)
		emps = append(emps,emp)
		if err!=nil {
			log.Println(err)
		}
	}
	log.Println(emps)
	return emps,nil
}