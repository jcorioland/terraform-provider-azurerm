package azurerm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func fakeResourceForTest() *schema.Resource {
	return &schema.Resource{
		Create: fakeResourceForTestCreate,
		Read:   fakeResourceForTestRead,
		Update: fakeResourceForTestUpdate,
		Delete: fakeResourceForTestDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		SchemaVersion: 2,

		Schema: map[string]*schema.Schema{

			"locations_list": {
				Type:     schema.TypeList,
				MinItems: 1,
				Optional: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					StateFunc:        azureRMNormalizeLocation,
					DiffSuppressFunc: azureRMSuppressLocationDiff,
				},
			},

			"locations_set": {
				Type:     schema.TypeSet,
				MinItems: 1,
				Optional: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					StateFunc:        azureRMNormalizeLocation,
					DiffSuppressFunc: azureRMSuppressLocationDiff,
				},
				Set: schema.HashString,
			},
		},
	}
}

func fakeResourceForTestCreate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("########## CREATE LIST ##########")
	// locations_list
	locationsList := d.Get("locations_list").([]interface{})
	if locationsList != nil && len(locationsList) > 0 {
		for _, location := range locationsList {
			fmt.Printf("Location (list): %s", location.(string))
			fmt.Println()
		}
	}
	fmt.Println("########## CREATE SET ##########")
	// locations_set
	locationsSet := d.Get("locations_set").(*schema.Set)
	if locationsSet != nil && locationsSet.Len() > 0 {
		for _, location := range locationsSet.List() {
			fmt.Printf("Location (set): %s", location.(string))
			fmt.Println()
		}
	}

	d.Set("locations_set", locationsSet)
	d.Set("locations_list", locationsList)
	d.SetId("1234")
	return nil
}

func fakeResourceForTestRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func fakeResourceForTestUpdate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("########## UPDATE LIST ##########")
	// locations_list
	locationsList := d.Get("locations_list").([]interface{})
	old, new := d.GetChange("locations_list")
	oldLocationsList := old.([]interface{})
	newLocationsList := new.([]interface{})
	hasChangeList := d.HasChange("locations_list")

	if hasChangeList {
		fmt.Println("locations_list variable has change")
	}

	if locationsList != nil && len(locationsList) > 0 {
		for _, location := range locationsList {
			fmt.Printf("Location (list): %s", location.(string))
			fmt.Println()
		}
	}

	if oldLocationsList != nil && len(oldLocationsList) > 0 {
		for _, oldLocation := range oldLocationsList {
			fmt.Printf("Old location (list): %s", oldLocation.(string))
			fmt.Println()
		}
	}

	if newLocationsList != nil && len(newLocationsList) > 0 {
		for _, newLocation := range newLocationsList {
			fmt.Printf("New location (list): %s", newLocation.(string))
			fmt.Println()
		}
	}

	fmt.Println("########## UPDATE SET ##########")
	// locations_set
	locationsSet := d.Get("locations_set").(*schema.Set)
	old, new = d.GetChange("locations_set")
	oldLocationsSet := old.(*schema.Set)
	newLocationsSet := new.(*schema.Set)
	hasChange := d.HasChange("locations_set")

	if hasChange {
		fmt.Println("locations_set variable has change")
	}

	if locationsSet != nil && locationsSet.Len() > 0 {
		for _, location := range locationsSet.List() {
			fmt.Printf("Location (set): %s", location.(string))
			fmt.Println()
		}
	}

	if oldLocationsSet != nil && oldLocationsSet.Len() > 0 {
		for _, oldLocation := range oldLocationsSet.List() {
			fmt.Printf("Old location (set): %s", oldLocation.(string))
			fmt.Println()
		}
	}

	if newLocationsSet != nil && newLocationsSet.Len() > 0 {
		for _, newLocation := range newLocationsSet.List() {
			fmt.Printf("New location (set): %s", newLocation.(string))
			fmt.Println()
		}
	}

	d.Set("locations_set", newLocationsSet)
	d.Set("locations_list", newLocationsList)
	return nil
}

func fakeResourceForTestDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
