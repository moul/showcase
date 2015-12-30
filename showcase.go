package moulshowcase

var ActionsMap map[string]Action

type Action func([]string) (*ActionResponse, error)

func RegisterAction(name string, action Action) {
	if ActionsMap == nil {
		ActionsMap = make(map[string]Action)
	}
	ActionsMap[name] = action
}

func Actions() map[string]Action {
	return ActionsMap
}

type ActionResponse struct {
	Body        interface{}
	IsJson      bool
	ContentType string
}

func PlainResponse(body interface{}) *ActionResponse {
	return &ActionResponse{
		Body:        body,
		ContentType: "text/plain",
		IsJson:      false,
	}
}
