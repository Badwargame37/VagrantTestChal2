//create code to run a application using golang and gin framework to create a web server and a rest api to get data from a database and display it on a web page
// Author:
// Date: 3/10/2021
// Version: 1.0.0

package main

import (
        "database/sql"
        "net/http"
//        "strings"
        "log"
        "github.com/gin-gonic/gin"
        _ "github.com/go-sql-driver/mysql"
)

type Information struct {
        ID             int    `json:"id"`
        FirstName      string `json:"first_name"`
        LastName       string `json:"last_name"`
        DateOfBirth    string `json:"date_of_birth"`
        DossierNumber  string `json:"dossier_number"`
        PurchaseDate   string `json:"purchase_date"`
        ExpirationDate string `json:"expiration_date"`
        CreationDate   string `json:"creation_date"`
}

func getDB() (*sql.DB, error) {
        db, err := sql.Open("mysql", "pour-sante:P0ur-S@nt3-2023@tcp(127.0.0.1:3306)/carte_pour_ma_sante")
        if err != nil {
                return nil, err
        }
        return db, nil
}
func createInformation(c *gin.Context) {
        var info Information
        if err := c.ShouldBindJSON(&info); err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
        }
        db, err := getDB()
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        defer db.Close()
        result, err := db.Exec("INSERT INTO information (first_name, last_name, date_of_birth, dossier_number, purchase_date, expiration_date, creation_date) VALUES (?, ?, ?, ?, ?, ?, ?)",
          info.FirstName, info.LastName, info.DateOfBirth, info.DossierNumber, info.PurchaseDate, info.ExpirationDate, info.CreationDate)
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        id, err := result.LastInsertId()
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        info.ID = int(id)
        c.JSON(http.StatusCreated, info)
  }
  func getInformationByID(c *gin.Context) {

        id := c.Param("id")
        db, err := getDB()
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        defer db.Close()
        var info Information
        err = db.QueryRow("SELECT id, first_name, last_name, date_of_birth, dossier_number, purchase_date, expiration_date, creation_date FROM information WHERE id=?", id).Scan(&info.ID, &info.FirstName, &info.LastName, &info.DateOfBirth, &info.DossierNumber, &info.PurchaseDate, &info.ExpirationDate, &info.CreationDate)
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        c.JSON(http.StatusOK, info)
  }
  func updateInformation(c *gin.Context) {
        id := c.Param("id")
        var info Information
        if err := c.ShouldBindJSON(&info); err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
        }
        db, err := getDB()
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        defer db.Close()
        _, err = db.Exec("UPDATE information SET first_name=?, last_name=?, date_of_birth=?, dossier_number=?, purchase_date=?, expiration_date=?, creation_date=? WHERE id=?",
          info.FirstName, info.LastName, info.DateOfBirth, info.DossierNumber, info.PurchaseDate, info.ExpirationDate, info.CreationDate, id)
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        c.JSON(http.StatusOK, info)
  }
  func deleteInformation(c *gin.Context) {
        id := c.Param("id")
        db, err := getDB()
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        defer db.Close()
        _, err = db.Exec("DELETE FROM information WHERE id=?", id)
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Information deleted"})
  }
  func listInformation(c *gin.Context) {
        db, err := getDB()
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        defer db.Close()
        rows, err := db.Query("SELECT id, first_name, last_name, date_of_birth, dossier_number, purchase_date, expiration_date, creation_date FROM information")
        if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
        }
        defer rows.Close()
        var infos []Information
        for rows.Next() {
          var info Information
          err := rows.Scan(&info.ID, &info.FirstName, &info.LastName, &info.DateOfBirth, &info.DossierNumber, &info.PurchaseDate, &info.ExpirationDate, &info.CreationDate)
          if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
          }
          infos = append(infos, info)
          }
          if err := rows.Err(); err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
          }
          c.JSON(http.StatusOK, infos)
          }
          func main() {
                router := gin.Default()

                router.GET("/information/list/:id", getInformationByID)
                router.POST("/information/new", createInformation)
                router.POST("/information/update/:id", updateInformation)
                router.POST("/information/delete/:id", deleteInformation)
                router.GET("/information/list", listInformation)

                if err := router.Run(":8080"); err != nil {
                  log.Fatal("Failed to start server: ", err)
                }
          }

		  