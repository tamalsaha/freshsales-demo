package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

type EntityType string

const (
	EntityContact      EntityType = "Contact"
	EntitySalesAccount EntityType = "SalesAccount"
	EntityDeal         EntityType = "Deal"
)

func (c *Client) CreateContact(lead *Contact) (*Contact, error) {
	resp, err := c.client.R().
		SetBody(APIObject{Contact: lead}).
		SetResult(&APIObject{}).
		Post("/api/leads")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request failed with status code = %d", resp.StatusCode())
	}
	return resp.Result().(*APIObject).Contact, nil
}

// Get Contact by id

// ref: https://developer.freshsales.io/api/#view_a_lead
// https://appscode.freshsales.io/leads/5022967942
//  /api/leads/[id]
/*
	curl -H "Authorization: Token token=sfg999666t673t7t82" -H "Content-Type: application/json" -X GET "https://domain.freshsales.io/api/leads/1"
*/
func (c *Client) GetContact(id int) (*Contact, error) {
	resp, err := c.client.R().
		SetResult(APIObject{}).
		Get(fmt.Sprintf("/api/leads/%d", id))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request failed with status code = %d", resp.StatusCode())
	}
	return resp.Result().(*APIObject).Contact, nil
}

func (c *Client) UpdateContact(lead *Contact) (*Contact, error) {
	resp, err := c.client.R().
		SetBody(APIObject{Contact: lead}).
		SetResult(&APIObject{}).
		Put(fmt.Sprintf("/api/leads/%d", lead.ID))
	if err != nil {
		panic(err)
	}
	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request failed with status code = %d", resp.StatusCode())
	}
	return resp.Result().(*APIObject).Contact, nil
}

func (c *Client) GetContactFilters() ([]ContactView, error) {
	resp, err := c.client.R().
		SetResult(ContactFilters{}).
		Get("/api/leads/filters")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request failed with status code = %d", resp.StatusCode())
	}
	return resp.Result().(*ContactFilters).Filters, nil
}

func (c *Client) ListAllContacts() ([]Contact, error) {
	views, err := c.GetContactFilters()
	if err != nil {
		return nil, err
	}
	viewId := -1
	for _, view := range views {
		if view.Name == "All Contacts" {
			viewId = view.ID
			break
		}
	}
	if viewId == -1 {
		return nil, fmt.Errorf("failed to detect view_id for \"All Contacts\"")
	}

	page := 1
	var leads []Contact
	for {
		resp, err := c.getContactPage(viewId, page)
		if err != nil {
			return nil, err
		}
		leads = append(leads, resp.Contacts...)
		if page == resp.Meta.TotalPages {
			break
		}
		page++
	}
	return leads, nil
}

func (c *Client) getContactPage(viewId, page int) (*ListResponse, error) {
	resp, err := c.client.R().
		SetResult(ListResponse{}).
		SetQueryParam("page", strconv.Itoa(page)).
		Get(fmt.Sprintf("/api/leads/view/%d", viewId))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request failed with status code = %d", resp.StatusCode())
	}
	return resp.Result().(*ListResponse), nil
}

func New(host, token string) *Client {
	return &Client{
		client: resty.New().
			EnableTrace().
			SetHostURL(host).
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", fmt.Sprintf("Token token=%s", token)),
	}
}

func main() {
	// https://codefx9.myfreshworks.com/crm/sales/personal-settings/api-settings
	domain := os.Getenv("FRESHSALES_DOMAIN")
	token := os.Getenv("FRESHSALES_TOKEN")
	fmt.Println(token)
	c := New(fmt.Sprintf("https://%s.myfreshworks.com/crm/sales", domain), token)
	fmt.Println(c)
}
