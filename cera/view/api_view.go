package view

import (
	"encoding/json"
	"log"
)

type ApiView struct {
	View
}

func (r *ApiView) Render() {
	r.jsonRender()
}

// render templates
func (r *ApiView) jsonRender() {
	defer r.Ctx.Done()
	rs, err := json.Marshal(r.Data)
	log.Println(string(rs))
	if err != nil {
		_, err := r.Ctx.Write([]byte(err.Error()))
		log.Println("error: render data to json failed, ", err)
		return
	}
	// set application/json header
	r.Ctx.Response.Header.Set("Content-type", "application/json")
	_, err = r.Ctx.Write(rs)
	if err != nil {
		log.Println("error: write json result to client failed,", err)
		return
	}
	return
}
