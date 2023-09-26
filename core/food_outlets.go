package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserRating struct {
	AverageRating float32 `json:"average_rating"`
	Votes         int32   `json:"votes"`
}

type Outlet struct {
	Id            int        `json:"id"`
	City          string     `json:"city"`
	Name          string     `json:"name"`
	EstimatedCost int        `json:"estimated_cost"`
	UserRating    UserRating `json:"user_rating"`
}

type Response struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       []Outlet `json:"data"`
}

func (r *Response) fetchOutlets(city string, total int) error {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/food_outlets?city=%s&page=%d", city, total)
	// fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("err initiate request %+v", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("err doing request %+v", err)
		return err
	}

	if resp.StatusCode >= 300 {
		log.Fatalf("err with status code: %d, reason: %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		er := fmt.Errorf("failed with status code %d, reason %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		return er
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		log.Fatalf("err while decode process %+v", err)
		return err
	}

	return nil
}

func ViableOutlets(city string, maxCost int32) {
	var outlets []Outlet
	r := &Response{}

	page := 1
	r.fetchOutlets(city, page)

	outlets = append(outlets, r.Data...)

	for page < r.TotalPages {
		if page > r.TotalPages {
			break
		}
		page += 1
		r.fetchOutlets(city, page)
		if r.Data != nil {
			outlets = append(outlets, r.Data...)
		}
	}

	for _, outlet := range outlets {
		if outlet.EstimatedCost <= int(maxCost) {
			fmt.Println(outlet.Name, outlet.EstimatedCost)
		}
	}

}

func getOutlets(city string, page int, ch chan Response) {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/food_outlets?city=%s&page=%d", city, page)
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("err initiate request %+v", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatalf("err while doing request %+v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("err doing request %+v", err)
		return
	}

	if resp.StatusCode >= 300 {
		log.Fatalf("err with status code: %d, reason: %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		return
	}

	defer resp.Body.Close()

	data := &Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("err decode response %+v", err)
		return
	}

	ch <- *data
}

func BestOutlets(city string, maxCost int32) {
	var outlets []Outlet
	r := &Response{}

	page := 1
	r.fetchOutlets(city, page)

	outlets = append(outlets, r.Data...)

	ch := make(chan Response, r.TotalPages-1)

	go func() {
		for page < r.TotalPages {
			page += 1
			getOutlets(city, page, ch)
		}
		close(ch)
	}()

	for resp := range ch {
		if len(resp.Data) > 0 {
			outlets = append(outlets, resp.Data...)
		}
	}

	for _, outlet := range outlets {
		if outlet.EstimatedCost <= int(maxCost) {
			fmt.Println(outlet.Name, outlet.EstimatedCost)
		}
	}
}
