package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPersonInfo(hi *gin.Context) {
	personID := hi.Param("person_id")
	db := getDB()
	var name string
	err := db.QueryRow("SELECT name FROM person WHERE id = ?", personID).Scan(&name)
	if err != nil {
		hi.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	var phoneNumber string
	err = db.QueryRow("SELECT number FROM phone WHERE person_id = ?", personID).Scan(&phoneNumber)
	if err != nil {
		hi.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}

	var addressID int
	err = db.QueryRow("SELECT address_id FROM address_join WHERE person_id = ?", personID).Scan(&addressID)
	if err != nil {
		hi.JSON(http.StatusNotFound, gin.H{"error": "Address join not found"})
		return
	}

	var city, state, street1, street2, zipCode string
	err = db.QueryRow("SELECT city, state, street1, street2, zip_code FROM address WHERE id = ?", addressID).Scan(&city, &state, &street1, &street2, &zipCode)
	if err != nil {
		hi.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	response := gin.H{
		"name":         name,
		"phone_number": phoneNumber,
		"city":         city,
		"state":        state,
		"street1":      street1,
		"street2":      street2,
		"zip_code":     zipCode,
	}

	hi.JSON(http.StatusOK, response)
}
