package interceptor

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}
type DetailValidation struct {
	Field       string `json:"field"`
	Value       string `json:"value"`
	Location    string `json:"location"`
	Description string `json:"description"`
}
type ResponseGetListData struct {
	Items      interface{} `json:"items"`
	Links      []Link      `json:"links"`
	TotalItems int32       `json:"total_items"`
	TotalPages int32       `json:"total_pages"`
}

type ResponsePostData struct {
	Items interface{} `json:"items"`
	Links []Link      `json:"links"`
}

type ResponseErrorData struct {
	Name    string           `json:"name"`
	Message string           `json:"message"`
	DebugId string           `json:"debug_id"`
	Details DetailValidation `json:"details"`
	Links   []Link           `json:"links"`
}

type ResponseUnauthorizedData struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
