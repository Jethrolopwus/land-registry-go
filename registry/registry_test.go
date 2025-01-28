package registry

import (
    "testing"
    "land-registry/models"  
)

func TestLandRegistry(t *testing.T) {
    lr := NewLandRegistry()

    t.Run("Add New Parcel", func(t *testing.T) {
        parcel := models.Parcel{
            ParcelNumber:     1,
            OwnerName:        "John Doe",
            LegalDescription: "Lot 123",
            AssessedValue:    250000,
        }

        err := lr.AddParcel(parcel)
        if err != nil {
            t.Errorf("Failed to add parcel: %v", err)
        }
    })

   
    t.Run("Get Existing Parcel", func(t *testing.T) {
        retrieved, err := lr.GetParcel(1)
        if err != nil {
            t.Errorf("Failed to get parcel: %v", err)
        }
        if retrieved.OwnerName != "John Doe" {
            t.Errorf("Expected owner 'John Doe', got '%s'", retrieved.OwnerName)
        }
    })

    t.Run("Add Duplicate Parcel", func(t *testing.T) {
        parcel := models.Parcel{
            ParcelNumber:     1,
            OwnerName:        "Jane Doe",
            LegalDescription: "Lot 456",
            AssessedValue:    300000,
        }

        err := lr.AddParcel(parcel)
        if err == nil {
            t.Error("Expected error when adding duplicate parcel, got nil")
        }
    })

    t.Run("Update Parcel", func(t *testing.T) {
        updatedParcel := models.Parcel{
            ParcelNumber:     1,
            OwnerName:        "Jane Doe",
            LegalDescription: "Lot 123 Updated",
            AssessedValue:    300000,
        }

        err := lr.UpdateParcel(updatedParcel)
        if err != nil {
            t.Errorf("Failed to update parcel: %v", err)
        }

        // Verify update
        retrieved, err := lr.GetParcel(1)
        if err != nil {
            t.Errorf("Failed to get updated parcel: %v", err)
        }
        if retrieved.OwnerName != "Jane Doe" {
            t.Errorf("Expected updated owner 'Jane Doe', got '%s'", retrieved.OwnerName)
        }
    })

    // Test case 5: Deleting a parcel
    t.Run("Delete Parcel", func(t *testing.T) {
        err := lr.DeleteParcel(1)
        if err != nil {
            t.Errorf("Failed to delete parcel: %v", err)
        }

        _, err = lr.GetParcel(1)
        if err == nil {
            t.Error("Expected error when getting deleted parcel, got nil")
        }
    })
}