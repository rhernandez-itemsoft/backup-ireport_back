package iresponse

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

// //ResponseModel Formato de respuesta
// // Esta estructura controla la informacion que se envía en la respuesta JSON
// type ResponseModel struct {
// 	ErrorCode int
// 	Messages  []string    `json:"messages"`
// 	Data      interface{} `json:"data"`
// }

// // //IResponse retorna un objeto tipo Response
// func IResponse() ResponseModel {
// 	return ResponseModel{
// 		Data:      nil,
// 		ErrorCode: 0,
// 		Messages:  nil,
// 	}
// }

//Definition esto se inyecta
type Definition struct {
	Ctx iris.Context //el contexto
}

//New Crea una nueva instancia de HTTPResponse
func New(ctx iris.Context) *Definition {
	return &Definition{
		Ctx: ctx,
	}
}

// //JSON retorna una respuesta en formato JSON
// func (def *Definition) JSON(_response ResponseModel) {
// 	if def.Ctx == nil {
// 		strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT.")
// 		fmt.Println(strErr)
// 		return
// 	}

// 	def.Ctx.StatusCode(200)
// 	def.Ctx.JSON(_response)
// }

// //JSONResponse retorna una respuesta JSON
// func (def *Definition) JSONResponse(statusCode int, data interface{}, iMessages ...string) {
// 	var msgs []string

// 	if def.Ctx == nil {
// 		strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT.")
// 		msgs = append(msgs, strErr)
// 	} else {
// 		for _, message := range iMessages {
// 			if (message!=""){
// 				msgs = append(msgs, message)
// 			}
// 		}
// 	}

// 	def.Ctx.StatusCode(statusCode)
// 	def.Ctx.JSON(map[string]interface{}{
// 		"Messages": msgs,
// 		"Data":     data,
// 	})
// }

//JSON retorna una respuesta JSON
func (def *Definition) JSON(statusCode int, data interface{}, iMessages ...string) interface{} {
	var msgs []string

	if def.Ctx == nil {
		strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT.")
		msgs = append(msgs, strErr)
	} else {
		for _, message := range iMessages {
			if message != "" {
				msgs = append(msgs, message)
			}
		}
	}

	def.Ctx.StatusCode(statusCode)
	def.Ctx.JSON(map[string]interface{}{
		"Messages": msgs,
		"Data":     data,
	})
	return nil
}
