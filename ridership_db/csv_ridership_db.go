package ridershipDB

import (
        "encoding/csv"
        "fmt"
        "os"
        "strconv"
)

type CsvRidershipDB struct {
        idIdxMap      map[string]int
        csvFile       *os.File
        csvReader     *csv.Reader
        num_intervals int
}

func (c *CsvRidershipDB) Open(filePath string) error {
        c.num_intervals = 9

        // Create a map that maps MBTA's time period ids to indexes in the slice
        c.idIdxMap = make(map[string]int)
        for i := 1; i <= c.num_intervals; i++ {
                timePeriodID := fmt.Sprintf("time_period_%02d", i)
                c.idIdxMap[timePeriodID] = i - 1
        }

        // create csv reader
        csvFile, err := os.Open(filePath)
        if err != nil {
                return err
        }
        c.csvFile = csvFile
        c.csvReader = csv.NewReader(c.csvFile)

        return nil
}

// TODO: some code goes here
// Implement the remaining RidershipDB methods

func (c *CsvRidershipDB) GetRidership(lineId string) ([]int64, error) {
        boarding := make([]int64, 9)
        contents, err := c.csvReader.ReadAll()
        if err != nil {
                fmt.Printf("Cannot read csv file")
        }

        for _, line := range contents {
                if line[0] != lineId {
                        continue
                }
                timePeriodId := line[2]
                index := c.idIdxMap[timePeriodId]
                //fmt.Printf("num: %s\n", line[4])
                if line[4] == "total_ons" {
                        continue
                }
                value, err := strconv.ParseInt(line[4], 10, 64)

                if err != nil {
                        fmt.Printf("Error parsint integer\n")
                        fmt.Println(err)
                }
                boarding[index] += value

        }
        return boarding, err
}
func (c *CsvRidershipDB) Close() error {
        err := c.csvFile.Close()
        return err

}
