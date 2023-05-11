package handlers

import (
	"fmt"
	"myapp/data"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/hilsonxhero/napoleon"
)

type Handlers struct {
	App    *napoleon.Napoleon
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	myData := "bar"
	h.App.Session.Put(r.Context(), "foo", myData)
	myValue := h.App.Session.GetString(r.Context(), "foo")
	vars := make(jet.VarMap)
	vars.Set("foo", myValue)
	err := h.App.Render.Page(w, r, "home", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

// JSON is the handler to demonstrate json responses
func (h *Handlers) JSON(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID      int64    `json:"id"`
		Name    string   `json:"name"`
		Hobbies []string `json:"hobbies"`
	}

	payload.ID = 10
	payload.Name = "Jack Jones"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

// XML is the handler to demonstrate XML responses
func (h *Handlers) XML(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		ID      int64    `xml:"id"`
		Name    string   `xml:"name"`
		Hobbies []string `xml:"hobbies>hobby"`
	}

	var payload Payload
	payload.ID = 10
	payload.Name = "John Smith"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteXML(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

// DownloadFile is the handler to demonstrate file download reponses
func (h *Handlers) DownloadFile(w http.ResponseWriter, r *http.Request) {
	h.App.DownloadFile(w, r, "./public/images", "celeritas.jpg")
}

func (h *Handlers) TestCrypto(w http.ResponseWriter, r *http.Request) {
	plainText := "Hello, world"
	fmt.Fprint(w, "Unencrypted: "+plainText+"\n")
	encrypted, err := h.encrypt(plainText)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Error500(w, r)
		return
	}

	fmt.Fprint(w, "Encrypted: "+encrypted+"\n")

	decrypted, err := h.decrypt(encrypted)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Error500(w, r)
		return
	}

	fmt.Fprint(w, "Decrypted: "+decrypted+"\n")
}
