package main

import (
    "encoding/json"
    "fmt" //added _ in front to get rid of warning //undid that 
    "github.com/tealeg/xlsx"
    "io/ioutil"
)

func main() {
    excelFileName := "west_mon_thurs.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    var stops []string //added cell in between []
    for _, sheet := range xlFile.Sheets {   
        //fmt.Printf(sheet)

        
        for _, row := range sheet.Rows { //og
            for _, cell := range row.Cells{
                stops = append(stops, cell.String() )//og
            }
        }
        
        //have given up on this one
        /*
        for _, row := range sheet.Rows{
            for _, col := range sheet.Cols {
                for _, cell := range col.Cells{
                    stops = append(stops, cell.String() )//og
            }
        }
    }
    */
        
    }

    // export as json
    stopsJson, err := json.Marshal(stops)
    if err != nil {
        fmt.Printf("error:", err) //changed from println to Printf
    }
    err = ioutil.WriteFile("schedule.json", stopsJson, 0644)
}