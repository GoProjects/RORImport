// rorimport
package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	//"fmt"
	//"github.com/jlaffaye/ftp"
	_ "github.com/lib/pq"
	"io"
	"os"
	"runtime"
	//"strings"
	"sync"
	"time"
)

const (
	bufferSize = 512000
)

type sharedMap struct {
	name	map[string]int
	rw		*sync.RWMutex
}

func (s *sharedMap)readMap(n string) (int, bool) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	i, ok := s.name[n]
	return i, ok
}

func (s *sharedMap)writeMap(n string, i int) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.name[n] = i
	return
}

var wgFile, wgFTP sync.WaitGroup

func main() {
	
	filog := InitFunc()

	vNumCPU, _ := Conf.Int("Default", "Quantity Proc")
	if vNumCPU <= 0 || vNumCPU > 32 {
		vNumCPU = runtime.NumCPU()
	}

	db, err := openDB()
	defer db.Close()
	if err != nil {
		CLog.PrintLog(true, err)
		os.Exit(1)
	}

	//mapFile, _ := Conf.GetValue("Default", "Map File")
	//mapBase, err := readJSONmap(mapFile)
	if err != nil {
		CLog.PrintLog(true, err)
		os.Exit(1)
	}

	chFileNames := make(chan string, vNumCPU)

	t0 := time.Now()
	
	dealerID := sharedMap{map[string]int{}, new(sync.RWMutex)}
	s3FileID := sharedMap{map[string]int{}, new(sync.RWMutex)}
	vehicleID := sharedMap{map[string]int{}, new(sync.RWMutex)}
	custVehicleID := sharedMap{map[string]int{}, new(sync.RWMutex)}
	customerID := sharedMap{map[string]int{}, new(sync.RWMutex)}

	for i := 1; i <= vNumCPU; i++ {
		go analizeFile(vehicleID, custVehicleID, customerID, s3FileID, dealerID, chFileNames, db)
	}
	getFTPFiles(dealerID, chFileNames, db)

	t1 := time.Now()
	PrintDeb("The call took %v to run.\n", t1.Sub(t0))
	os.Chdir("..")
	/*if err := os.RemoveAll("_temp_"); err != nil {
		CLog.PrintLog(true, "Can't remove the temporary directory '_temp_'. ", err)
	}*/
	defer filog.Close()
}

func analizeFile(vehicleID, custVehicleID, customerID, s3FileID, dealerID sharedMap, ch chan string, db *sql.DB) {

	for fn := range ch {
		records, err := parceFile(fn)
		if err != nil {
			continue
		}
		l := len(records)
		if l < 2 {
			CLog.PrintLog(true, "The file ", fn, " is empty or has wrong format.")
		} else if records[0][0] != "FileType" {
			CLog.PrintLog(true, "The file ", fn, " has wrong format.")
		} else {
			fillDataBase(vehicleID, custVehicleID, customerID, s3FileID, dealerID, records, fn, db)
		}
		wgFTP.Done()
	}
	return
}

func parceFile(fn string) (records [][]string, err error) {
	fileDesc, err := os.Open(fn)
	if err != nil {
		CLog.PrintLog(false, "Error open file:", fn, " ", err)
		wgFTP.Done()
		return nil, err
	}
	defer fileDesc.Close()
	ioreader := io.Reader(fileDesc)
	reader := bufio.NewReader(ioreader)
	lineReader := csv.NewReader(reader)
	lineReader.Comma = 0x09
	records, err = lineReader.ReadAll()
	if err != nil {
		CLog.PrintLog(false, "Error read from file:", fn, " ", err)
		wgFTP.Done()
		return nil, err
	}
	return records, nil
}

