package core

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/adrg/xdg"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const SUBDIR = "/habit-thingemabob"

type someData [][]any

func New(name string, description string) (string, appErr) {
	data, exist, err := checkAndLoadFile()
	if Exists(err) {
		return "", MakeAppErr(2, "Adding new entry", err)
	}
	streak, startTime, lastTime, metadata := "_", time.Now().Unix(), time.Now().Unix(), ""
	if false == exist {
		data = CoreData{Name: []string{name}, Description: []string{description}, Streak: []string{streak}, StartTime: []int64{startTime}, LastTime: []int64{lastTime}, Metadata: []string{metadata}}
	} else {
		data.Name = append(data.Name, name)
		data.Description = append(data.Description, description)
		data.Streak = append(data.Streak, streak)
		data.StartTime = append(data.StartTime, startTime)
		data.LastTime = append(data.LastTime, lastTime)
		data.Metadata = append(data.Metadata, metadata)
	}
	updateFile(data)
	return xdg.DataHome + SUBDIR, nil
}

func Check(name string, isUnCheck bool) (string, appErr) {
	data, exist, err := checkAndLoadFile()
	if Exists(err) {
		return "", err.Update("Operation failed when loading file. Further details:")
	}
	if exist == false {
		return "", appError{code: 69, error: "Habit not found, in fact no habits found. Try the 'new' sub-command"}
	}
	index := slices.Index(data.Name, name)
	if index < 0 {
		return "", appError{code: 420, error: "Habit not found, run 'new' for a new habit or 'list' to check existing habits."}
	}
	slen := len(data.Streak[index])
	data.Streak[index] = data.Streak[index][:slen-1] + "*"
	splitstr := strings.Split(data.Streak[index], "_")
	updateFile(data)
	entry := "*"
	if isUnCheck {
		entry = "_"
	}
	return fmt.Sprintf("Streak: %d", strings.Count(splitstr[len(splitstr)-1], entry)), nil
}

func generalDataUpdate(data CoreData) (CoreData, appErr) {
	for i := 0; i < len(data.Name); i++ {
		recDays := len(data.Streak[i])
		days := (time.Now().Unix() - data.StartTime[i]) * 60 * 60 * 24
		data.Streak[i] = data.Streak[i] + strings.Repeat("_", int(days)-int(recDays))

	}
	return data, nil

}

func checkAndLoadFile() (CoreData, bool, appErr) {
	directory := xdg.DataHome + SUBDIR
	if _, error := os.Stat(directory); error != nil {
		return CoreData{}, false, nil
	}
	data, err := os.ReadFile(directory)
	if Exists(err) {
		return CoreData{}, true, MakeAppErr(7, "loading file", err)
	}
	newData, aErr := processCSV(data, '\t')
	if Exists(aErr) {
		return CoreData{}, true, aErr.Update("check and load file")
	}
	newData, aErr = generalDataUpdate(newData)
	return newData, true, nil
}

func updateFile(data CoreData) error {
	str, err := processData(data, '\t')
	if Exists(err) {
		return err
	}
	return os.WriteFile(xdg.DataHome+SUBDIR, []byte(str), 0644)
}

func processData(data CoreData, split rune) (string, error) {
	sliceSlices := make([][]string, len(data.Name))
	for i := 0; i < len(data.Name); i++ {
		slice := make([]string, 6)
		slice[0] = data.Name[i]
		slice[1] = data.Description[i]
		slice[2] = data.Streak[i]
		slice[3] = fmt.Sprintf("%d", data.StartTime[i])
		slice[4] = fmt.Sprintf("%d", data.LastTime[i])
		slice[5] = data.Metadata[i]
		sliceSlices[i] = slice
	}
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	writer.Comma = split
	err := writer.WriteAll(sliceSlices)
	if Exists(err) {
		return "", err
	}
	return buf.String(), nil
}
func processCSV(csvRaw []byte, split rune) (CoreData, appErr) {
	r := bytes.NewReader(csvRaw)

	reader := csv.NewReader(r)
	reader.Comma = split
	records, err := reader.ReadAll()
	if Exists(err) {
		return CoreData{}, MakeAppErr(5, "decoding the file", err)
	}
	name, description, streak, startTime, lastTime, metadata := make([]string, 0, 10), make([]string, 0, 10), make([]string, 0, 10), make([]int64, 0, 10), make([]int64, 0, 10), make([]string, 0, 10)

	for _, l := range records {
		sT, err := strconv.ParseInt(l[3], 10, 64)
		lT, err := strconv.ParseInt(l[4], 10, 64)
		if Exists(err) {
			return CoreData{}, MakeAppErr(6, "reconstructing struct", err)
		}
		name = append(name, l[0])
		description = append(description, l[1])
		streak = append(streak, l[2])
		startTime = append(startTime, sT)
		lastTime = append(lastTime, lT)
		metadata = append(metadata, l[5])
	}
	return CoreData{name, description, streak, startTime, lastTime, metadata}, nil
}
