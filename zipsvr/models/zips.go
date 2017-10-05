package models

import (
    "os"
    "fmt"
    "encoding/csv"
    "io"
)

type Zip struct {
    Code  string
    City  string
    State string
}

type ZipSlice []*Zip

type ZipIndex map[string]ZipSlice

func LoadZips(filePath string) (ZipSlice, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("error opening file %v", err)
    }

    reader := csv.NewReader(f)
    // Read and discard header row
    _, err = reader.Read()
    if err != nil {
        return nil, fmt.Errorf("error reading header row: %v", err)
    }

    zips := make(ZipSlice, 0, 43000)

    for {
        fields, err := reader.Read()

        if err == io.EOF {
            return zips, nil
        }

        if err != nil {
            return nil, fmt.Errorf("error reading record: %v", err)
        }

        z := &Zip{
            Code:  fields[0],
            City:  fields[3],
            State: fields[6],
        }

        zips = append(zips, z)
    }

    return zips, err
}
