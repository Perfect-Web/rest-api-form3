package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type AutoGenerated struct {
	Data []struct {
		Type           string `json:"type"`
		ID             string `json:"id"`
		Version        int    `json:"version"`
		OrganisationID string `json:"organisation_id"`
		Attributes     struct {
			Amount           string `json:"amount"`
			BeneficiaryParty struct {
				AccountName       string `json:"account_name"`
				AccountNumber     string `json:"account_number"`
				AccountNumberCode string `json:"account_number_code"`
				AccountType       int    `json:"account_type"`
				Address           string `json:"address"`
				BankID            string `json:"bank_id"`
				BankIDCode        string `json:"bank_id_code"`
				Name              string `json:"name"`
			} `json:"beneficiary_party"`
			ChargesInformation struct {
				BearerCode    string `json:"bearer_code"`
				SenderCharges []struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"sender_charges"`
				ReceiverChargesAmount   string `json:"receiver_charges_amount"`
				ReceiverChargesCurrency string `json:"receiver_charges_currency"`
			} `json:"charges_information"`
			Currency    string `json:"currency"`
			DebtorParty struct {
				AccountName       string `json:"account_name"`
				AccountNumber     string `json:"account_number"`
				AccountNumberCode string `json:"account_number_code"`
				Address           string `json:"address"`
				BankID            string `json:"bank_id"`
				BankIDCode        string `json:"bank_id_code"`
				Name              string `json:"name"`
			} `json:"debtor_party"`
			EndToEndReference string `json:"end_to_end_reference"`
			Fx                struct {
				ContractReference string `json:"contract_reference"`
				ExchangeRate      string `json:"exchange_rate"`
				OriginalAmount    string `json:"original_amount"`
				OriginalCurrency  string `json:"original_currency"`
			} `json:"fx"`
			NumericReference     string `json:"numeric_reference"`
			PaymentID            string `json:"payment_id"`
			PaymentPurpose       string `json:"payment_purpose"`
			PaymentScheme        string `json:"payment_scheme"`
			PaymentType          string `json:"payment_type"`
			ProcessingDate       string `json:"processing_date"`
			Reference            string `json:"reference"`
			SchemePaymentSubType string `json:"scheme_payment_sub_type"`
			SchemePaymentType    string `json:"scheme_payment_type"`
			SponsorParty         struct {
				AccountNumber string `json:"account_number"`
				BankID        string `json:"bank_id"`
				BankIDCode    string `json:"bank_id_code"`
			} `json:"sponsor_party"`
		} `json:"attributes"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}