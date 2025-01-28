## A simple Go implementation of how indexing works
### This project is an example of how Indexing basically work by making blockchain data searchable and accessible in Go lang

## USer's Guide/Understanding the program
* In this program, the indexing mechanism happens in the landRegistry struct. 
* The map data structure acts as the indexer. It uses(uint64)ParcelNumber as the key to instandtly find parcels.
* you can think of it like a filling cabinet where each drawer has it unigue number.
## How indexing works in this main program
### Single Record Operation
* This occur in the addParcel function in main.go file. when you add this parcel, it's stored in the map with key = 1.
* So looking up this parcel later is instant because the program will just check map[1] which is most faster than searching through all parcels one by one.
  `type LandRegistry struct {
    parcels map[uint64]models.Parcel
    mutex   sync.RWMutex
}`

  `newParcel := models.Parcel{
        ParcelNumber:     1,
        OwnerName:        "Jethro B",
        LegalDescription: "Plot 123, Block 45, Giring District, Jos south",
        AssessedValue:    250000,
    } 
landRegistry.AddParcel(parcel)`
### Concurrent Access(multiple records)
* This shows how indexer handles multiple operations with each parcel is added to the map concurrently.
* Each parcel is still instantly accessible by it's number(parcelNumber) and the mutex ensure safe access to the map.

` for i := 0; i < 10; i++ {
        go func(i int) {
        parcel := models.Parcel{
            ParcelNumber:     uint64(i + 100),
            OwnerName:        fmt.Sprintf("Owner %d", i),
            LegalDescription: fmt.Sprintf("Lot %d", i),
            AssessedValue:    uint64(200000 + i*1000),
        }
        landRegistry.AddParcel(parcel)
    }(i)
    }`
    
* This makes querrying the data very efficient by:
1. Instant Access
2. Memory organization
3. Concurrency Support
4. Error Handling
## Real-World Analogy
* Think of it like a library.
* Traditional search looking through every book on every shelf (without Indexing). and
* This program/system (with Indexing) is like using the library's computer system to instantly locate a book by it's code.

## usage of this program
* To clone the repo,
* Ensure that you have Go install on your device
* `git clone <url>`
* cd to the clone Directory
* open it on the IDE of your choice
* run this program on your terminal: `go run main.go`
* run the test on your terminal: `go test ./registry/ -v`



