# TODO LIST API

## Summary

Aplikasi ini dapat menggnakan bebrapa database. Aplikasi ini juga support filtering, pagination, dan option untuk menampilkan sub list atau tidak. Filtering dan pagination juga bisa di mix sesuai kebutuhan dengan query setelah tanda "?".

Untuk dokumentasi lengkapnya silahkan buka link berikut :
https://documenter.getpostman.com/view/23649575/2s9YC5xrkP

## Running Application

### 1. Setting Env

Buat database terlebih dahulu. Lalu setting env pada file .env dengan mengubah settingan database sesuai kebutuhan. Contoh :

    DB_HOST_POSTGRES="127.0.0.1"
    DB_USER_POSTGRES=postgres
    DB_PORT_POSTGRES=5432
    DB_PASSWORD_POSTGRES=123
    DB_NAME_POSTGRES=todolist
    DB_SSLMODE_POSTGRES=disable
    DB_TIMEZONE_POSTGRES=Asia/Jakarta

### 2. Setting Database

Anda bisa menambahkan, menghapus atau mengganti database yang akan digunakan dengan mengubah pengaturan file type.go dan databse.go pada folder database.

### 3. Runing App

Untuk menjalankan aplikasi bisa menggunakan command :

    go run main.go

### 4. Testing

Untuk testing dapat masuk ke folder app terlebih dahulu menggunakan command :

    cd internal/app

Lalu untuk menjalankan testing dapat menggunakan command :

    go test ./... -cover

Jika ingin melihat testing step by step dapat menggunakan command :

    go test ./... -v -cover

Contoh hasil nya seperti ini :

    go test ./... -cover
    ok      moonlay-todolist/internal/app/list      0.028s  coverage: 72.6% of statements
    ok      moonlay-todolist/internal/app/sublist   0.027s  coverage: 71.8% of statements

