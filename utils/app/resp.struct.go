package app

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}
type Detail struct {
	Field       string      `json:"field"`
	Value       interface{} `json:"value"`
	Location    string      `json:"location"`
	Description string      `json:"description"`
}
type ResponseGetListData struct {
	Items       []interface{} `json:"items"`
	Links       []Link        `json:"links"`
	CurrentPage int           `json:"current_page"`
	TotalItems  int           `json:"total_items"`
	TotalPages  int           `json:"total_pages"`
}

type ResponsePostData struct {
	Items []interface{} `json:"items"`
	Links []Link        `json:"links"`
}

type ResponseErrorData struct {
	Name    string   `json:"name"`
	Message string   `json:"message"`
	DebugId string   `json:"debug_id"`
	Details []Detail `json:"details"`
	Links   []Link   `json:"links"`
}

type ResponseUnauthorizedData struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
}
