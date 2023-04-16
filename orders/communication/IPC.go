package communication

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func GetProductDetails(id int) *Product {

	req_body := FetchProductsRequest{
		IDs: []int{id},
	}
	json_str_data, err := json.Marshal(req_body)
	if err != nil {
		log.Println(err)
		return nil
	}
	req, err := http.Post(os.Getenv("PRODUCTS_ADDRESS")+"/api/products/fetch", "application/json", bytes.NewBuffer(json_str_data))
	if err != nil {
		log.Println("hello", err)
		return nil
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}
	resp_struct := Products{}
	json.Unmarshal(body, &resp_struct)
	if !resp_struct.Status || len(resp_struct.Products) == 0 {
		//Means the service is down or cant find product
		return nil
	} else if resp_struct.Status {
		return &resp_struct.Products[0]
	}
	return nil
}
