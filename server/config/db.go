package config


import (
    "crypto/tls"
    "crypto/x509"
    "database/sql"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    
    "github.com/joho/godotenv"
    "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    er := godotenv.Load()
    if er != nil {
        log.Fatal("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    dbSSLCA := os.Getenv("DB_SSL_CA")

    rootCertPool := x509.NewCertPool()
    pem, err := ioutil.ReadFile(dbSSLCA)
    if err != nil {
        log.Fatalf("Failed to read CA cert file: %v", err)
    }
    if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
        log.Fatal("Failed to append CA cert")
    }

    tlsConfig := &tls.Config{
        RootCAs: rootCertPool,
    }

    err = mysql.RegisterTLSConfig("custom", tlsConfig)
    if err != nil {
        log.Fatalf("Failed to register TLS config: %v", err)
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=custom",
        dbUser,
        dbPassword,
        dbHost,
        dbPort,
        dbName,
    )

    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        fmt.Print("Error: ", err)
        panic("Cannot connect to the database")
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    fmt.Print("Successfully connected to the database!!")
}