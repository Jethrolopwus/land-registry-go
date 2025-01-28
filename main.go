package main

import (
    "fmt"
    "sync"
    "land-registry/models"   
    "land-registry/registry" 
)

func main() {

    landRegistry := registry.NewLandRegistry()

    newParcel := models.Parcel{
        ParcelNumber:     1,
        OwnerName:        "Jethro B",
        LegalDescription: "Plot 123, Block 45, Giring District, Jos south",
        AssessedValue:    250000,
    }

  
    if err := landRegistry.AddParcel(newParcel); err != nil {
        fmt.Printf("Error adding parcel: %v\n", err)
        return
    }

    if parcel, err := landRegistry.GetParcel(1); err == nil {
        fmt.Printf("Found parcel:\nOwner: %s\nDescription: %s\nValue: $%d\n",
            parcel.OwnerName,
            parcel.LegalDescription,
            parcel.AssessedValue)
    }


    // Test concurrent access
    var wg sync.WaitGroup
        for i := 0; i < 10; i++ {
     wg.Add(1)
        go func(i int) {
        defer wg.Done()
        parcel := models.Parcel{
            ParcelNumber:     uint64(i + 100),
            OwnerName:        fmt.Sprintf("Owner %d", i),
            LegalDescription: fmt.Sprintf("Lot %d", i),
            AssessedValue:    uint64(200000 + i*1000),
        }
        landRegistry.AddParcel(parcel)
    }(i)
    }
    wg.Wait()
}