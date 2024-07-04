package data

import "fmt"

func (a *model) True() error {
	response := Response{
		Ok: true,
	}
	return a.Status(200).JSON(response)
}

func (a *model) False() error {
	response := Response{
		Ok: false,
	}
	return a.Status(200).JSON(response)
}

func (a *model) Error(err interface{}) error {
	response := Response{
		Ok: false,
	}
	if err != nil {
		response.Error = fmt.Sprintf("%v", err)
	}
	return a.Status(200).JSON(response)
}
func (a *model) Data(data interface{}) error {
	response := Response{
		Ok: true,
	}
	if data != nil {
		response.Data = data
	}
	return a.Status(200).JSON(response)
}

func (a *model) Message(message string) error {
	response := Response{
		Ok:      true,
		Message: message,
	}
	return a.Status(200).JSON(response)
}
func (a *model) Success() error {
	response := Response{
		Ok:      true,
		Message: "success",
	}
	return a.Status(200).JSON(response)
}
