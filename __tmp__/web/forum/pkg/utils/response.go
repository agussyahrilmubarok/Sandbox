package utils

type ResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFailure struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func StructNonNilSlice(slice []string) []string {
	if slice == nil {
		return []string{}
	}
	return slice
}
