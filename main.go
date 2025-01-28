
package main

import (
    "fmt"
    "sync"
    "land-registry/models"   
    "land-registry/registry" 
)

func displayParcel(parcel models.Parcel) {
    fmt.Printf("\nParcel Details:\n"+
        "Parcel Number: %d\n"+
        "Owner: %s\n"+
        "Description: %s\n"+
        "Value: $%d\n"+
        "------------------------\n",
        parcel.ParcelNumber,
        parcel.OwnerName,
        parcel.LegalDescription,
        parcel.AssessedValue)
}

func main() {
    landRegistry := registry.NewLandRegistry()
    
    newParcel := models.Parcel{
        ParcelNumber:     1,
        OwnerName:        "Jethro B",
        LegalDescription: "Plot 123, Block 45, Giring District, Jos south",
        AssessedValue:    50000,
    }
  
    if err := landRegistry.AddParcel(newParcel); err != nil {
        fmt.Printf("Error adding parcel: %v\n", err)
        return
    }

    fmt.Println("Initial Parcel Added:")
    if parcel, err := landRegistry.GetParcel(1); err == nil {
        displayParcel(parcel)
    }

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            parcel := models.Parcel{
                ParcelNumber:     uint64(i + 100),
                OwnerName:        fmt.Sprintf("Owner %d", i),
                LegalDescription: fmt.Sprintf("Plot %d Giring Distric Jos south", i),
                AssessedValue:    uint64(200000 + i*1000),
            }
            if err := landRegistry.AddParcel(parcel); err != nil {
                fmt.Printf("Error adding parcel %d: %v\n", i+100, err)
            }
        }(i)
    }
    wg.Wait()

    fmt.Println("\nAll Parcels in Registry:")
    if parcel, err := landRegistry.GetParcel(1); err == nil {
        displayParcel(parcel)
    }

    for i := 0; i < 10; i++ {
        if parcel, err := landRegistry.GetParcel(uint64(i + 100)); err == nil {
            displayParcel(parcel)
        }
    }
}