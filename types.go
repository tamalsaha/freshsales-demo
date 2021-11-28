package main

import "time"

type Contact struct {
	ID          int         `json:"id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	DisplayName string      `json:"display_name"`
	Avatar      interface{} `json:"avatar"`
	JobTitle    interface{} `json:"job_title"`
	City        interface{} `json:"city"`
	State       interface{} `json:"state"`
	Zipcode     interface{} `json:"zipcode"`
	Country     interface{} `json:"country"`
	Email       string      `json:"email"`
	Emails      []struct {
		ID        int         `json:"id"`
		Value     string      `json:"value"`
		IsPrimary bool        `json:"is_primary"`
		Label     interface{} `json:"label"`
		Destroy   bool        `json:"_destroy"`
	} `json:"emails"`
	TimeZone        interface{} `json:"time_zone"`
	WorkNumber      interface{} `json:"work_number"`
	MobileNumber    interface{} `json:"mobile_number"`
	Address         interface{} `json:"address"`
	LastSeen        interface{} `json:"last_seen"`
	LeadScore       int         `json:"lead_score"`
	LastContacted   interface{} `json:"last_contacted"`
	OpenDealsAmount interface{} `json:"open_deals_amount"`
	WonDealsAmount  interface{} `json:"won_deals_amount"`
	Links           struct {
		Conversations        string `json:"conversations"`
		TimelineFeeds        string `json:"timeline_feeds"`
		DocumentAssociations string `json:"document_associations"`
		Notes                string `json:"notes"`
		Tasks                string `json:"tasks"`
		Appointments         string `json:"appointments"`
		Reminders            string `json:"reminders"`
		Duplicates           string `json:"duplicates"`
		Connections          string `json:"connections"`
	} `json:"links"`
	LastContactedSalesActivityMode interface{} `json:"last_contacted_sales_activity_mode"`
	CustomField                    struct {
	} `json:"custom_field"`
	CreatedAt                     time.Time     `json:"created_at"`
	UpdatedAt                     time.Time     `json:"updated_at"`
	Keyword                       interface{}   `json:"keyword"`
	Medium                        interface{}   `json:"medium"`
	LastContactedMode             interface{}   `json:"last_contacted_mode"`
	RecentNote                    interface{}   `json:"recent_note"`
	WonDealsCount                 int           `json:"won_deals_count"`
	LastContactedViaSalesActivity interface{}   `json:"last_contacted_via_sales_activity"`
	CompletedSalesSequences       interface{}   `json:"completed_sales_sequences"`
	ActiveSalesSequences          interface{}   `json:"active_sales_sequences"`
	WebFormIds                    interface{}   `json:"web_form_ids"`
	OpenDealsCount                int           `json:"open_deals_count"`
	LastAssignedAt                time.Time     `json:"last_assigned_at"`
	Tags                          []interface{} `json:"tags"`
	Facebook                      interface{}   `json:"facebook"`
	Twitter                       interface{}   `json:"twitter"`
	Linkedin                      interface{}   `json:"linkedin"`
	IsDeleted                     bool          `json:"is_deleted"`
	TeamUserIds                   interface{}   `json:"team_user_ids"`
	ExternalID                    interface{}   `json:"external_id"`
	WorkEmail                     interface{}   `json:"work_email"`
	SubscriptionStatus            int           `json:"subscription_status"`
	SubscriptionTypes             string        `json:"subscription_types"`
	CustomerFit                   int           `json:"customer_fit"`
	WhatsappSubscriptionStatus    int           `json:"whatsapp_subscription_status"`
	PhoneNumbers                  []interface{} `json:"phone_numbers"`
}

type ListResponse struct {
	Contacts []Contact `json:"contacts"`
	Meta     ListMeta  `json:"meta"`
}

type ListMeta struct {
	TotalPages int `json:"total_pages"`
	Total      int `json:"total"`
	AvatarData struct {
		Contact struct {
			Num17011374038 interface{} `json:"17011374038"`
			Num17010685339 string      `json:"17010685339"`
		} `json:"Contact"`
	} `json:"avatar_data"`
}

type APIObject struct {
	Contact *Contact `json:"contact,omitempty"`
}

type ContactView struct {
	ID                     int           `json:"id"`
	Name                   string        `json:"name"`
	ModelClassName         string        `json:"model_class_name"`
	UserID                 int           `json:"user_id"`
	IsDefault              bool          `json:"is_default"`
	UpdatedAt              time.Time     `json:"updated_at"`
	UserName               interface{}   `json:"user_name"`
	CurrentUserPermissions []interface{} `json:"current_user_permissions"`
}

type ContactFilters struct {
	Filters []ContactView `json:"filters"`
}
