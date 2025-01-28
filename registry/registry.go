package registry

import (
    "errors"
    "sync"
    "land-registry/models"
)

type LandRegistry struct {
    parcels map[uint64]models.Parcel
    mutex   sync.RWMutex
}

func NewLandRegistry() *LandRegistry {
    return &LandRegistry{
        parcels: make(map[uint64]models.Parcel),
    }
}

func (lr *LandRegistry) AddParcel(parcel models.Parcel) error {
    lr.mutex.Lock()
    defer lr.mutex.Unlock()

    if _, exists := lr.parcels[parcel.ParcelNumber]; exists {
        return errors.New("parcel already exists")
    }

    lr.parcels[parcel.ParcelNumber] = parcel
    return nil
}

func (lr *LandRegistry) GetParcel(parcelNumber uint64) (models.Parcel, error) {
    lr.mutex.RLock()
    defer lr.mutex.RUnlock()

    parcel, exists := lr.parcels[parcelNumber]
    if !exists {
        return models.Parcel{}, errors.New("parcel not found")
    }

    return parcel, nil
}

func (lr *LandRegistry) UpdateParcel(parcel models.Parcel) error {
    lr.mutex.Lock()
    defer lr.mutex.Unlock()

    if _, exists := lr.parcels[parcel.ParcelNumber]; !exists {
        return errors.New("parcel does not exist")
    }

    lr.parcels[parcel.ParcelNumber] = parcel
    return nil
}

func (lr *LandRegistry) DeleteParcel(parcelNumber uint64) error {
    lr.mutex.Lock()
    defer lr.mutex.Unlock()

    if _, exists := lr.parcels[parcelNumber]; !exists {
        return errors.New("parcel does not exist")
    }

    delete(lr.parcels, parcelNumber)
    return nil
}