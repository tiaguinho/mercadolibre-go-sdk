package meli

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var (
	productID string
)

//Test Get function
func TestGet(t *testing.T) {
	params := map[string]string{
		"access_token": client.MLToken.AccessToken,
	}

	data, err := client.Get("sites/MLB/categories", params)

	if err != nil {
		t.Error(err)
	}

	var categories []Category
	json.Unmarshal(data, &categories)

	if len(categories) == 0 {
		t.Error("Categories not found")
	}
}

//Test Post function
func TestPost(t *testing.T) {
	params := map[string]string{
		"access_token": client.MLToken.AccessToken,
	}

	product := Product{
		ListingTypeID:     "free",
		Title:             "Golang SDK Title",
		Description:       "Golang SDK description product test",
		CategoryID:        "MLB50655",
		BuyingMode:        "buy_it_now",
		CurrencyID:        "BRL",
		Condition:         "new",
		Price:             100.00,
		AvailableQuantity: 1,
		Pictures: []Image{
			{
				Source: "http://i.stack.imgur.com/DJBD5.png",
			},
		},
	}

	data, err := client.Post("items", product, params)
	if err != nil {
		t.Error(err)
	}

	json.Unmarshal(data, &product)

	productID = product.ID
}

//Test Put function
func TestPut(t *testing.T) {
	params := map[string]string{
		"access_token": client.MLToken.AccessToken,
	}

	var status Status

	current_status := "not_yet_active"
	for current_status != "active" {
		data, _ := client.Get(fmt.Sprintf("items/%s", productID), params)
		json.Unmarshal(data, &status)

		current_status = status.Status
		time.Sleep(time.Second * 20)

		fmt.Println("Wait for the product to be active...")
	}

	status.Status = "closed"
	_, err := client.Put(fmt.Sprintf("items/%s", productID), status, params)
	if err != nil {
		t.Error(err)
	}
}

//Test Delete function
func TestDelete(t *testing.T) {
	params := map[string]string{
		"access_token": client.MLToken.AccessToken,
	}

	_, err := client.Delete("questions/4161429753", params)
	if err != nil {
		t.Error(err)
	}
}
