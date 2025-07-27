package core

import (
	"github.com/adrg/xdg"
	"time"
	"fmt"
	"io"
	"os"
	"errors"
	"encoding/csv"
	"strings"
	"bytes"
)
const SUBDIR= "/habit-thingemabob"
func checkAndLoadFile() (CoreData, bool, error){
	if _ , error := os.Stat(fmt.Sprintf(xdg.DataHome, SUBDIR)); error != nil {
		return CoreData{}, false,nil
}
	data, err := os.ReadFile(fmt.Sprintf(xdg.DataHome, SUBDIR))
	if Is(err) {
		return CoreData{}, true, fmt.Errorf("failed reading file: %e",err)
	}
	return CoreData{}, true,nil
}


func New(name string) (CoreData, appError){
	data, exist := checkAndLoadFile()
	if false == exist {
		data = CoreData{Name: []string{name}, Streak: "_",StartTime: time.Now().Unix(),LastTime: time.Now().Unix(),Metadata: ""}
	} else {
		panic("didn't bother with this yet")
	}

	
}


func updateFile(data CoreData) error{
	for i := 0; i < len(data.Name); i++ {
		make([]string,)
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	writer.Comma = '\t'
	writer.WriteAll(
	return os.WriteFile(fmt.Sprintf(xdg.DataHome, SUBDIR),0644)
}

func processData(data CoreData, split rune) string {

}

