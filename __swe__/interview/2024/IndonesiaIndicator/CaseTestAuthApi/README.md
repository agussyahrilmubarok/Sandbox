# Backend Technical Test

## Overview

### User Authentication API

1. Scenario
   Web service is Restful API for user authentication that allows users to register, log in and
   access protected resources.
2. Requirements
    - Registration: As a visitor, I want to register a user account with username, phone number and password.
      The acceptance criteria is as follows:
        - The username must not already exist in the system.
        - The phone number must valid format phone number.
        - The password should hash when store into database.
        - The password must be at least eight characters.
        - The password must contain at least one number.
        - The password must contain at least one letter.
        - The response should return message success when is user registered.
    - Login: As a registered user, I can use my username with the password to log into application.
      The acceptance criteria is as follows:
        - The username must is already exist in the system.
        - The password must be at least eight characters.
        - The password must contain at least one number.
        - The password must contain at least one letter.
        - The response should return token access when is valid credentials.
    - Resource: As a registered user, I can access protected resource.
      The acceptance criteria is as follows:
        - The response should return message success access protected resource.

## How to run

1. Setup MySQL database (required)

      ```bash
     MYSQL_DATABASE=demo
     MYSQL_USER=developer
     MYSQL_PASSWORD=my-secret-pw
      ```
   or run on docker container with run command `docker compose -f docker-mysql.yml up`.

2. Maven install
   ```bash
   ./mvnw install
   ```

3. Run testing

   ```bash
   ./mvnw clean test
   ```

3. Run application without testing

   ```bash
   ./mvnw spring-boot:run -DskipTests
   ```

5. Run testing with postman
   Import file `Demo API Testing.postman_collection.json` into postman application then run collection.

## Description

Dalam membuat API Autentikasi Pengguna, saya memutuskan untuk membuat web service dengan teknologi dari Java.
Spring framework digunakan untuk memanfaatkan kemudahan dengan Dependency Injection. Saya mengimplementasikan
Three-Layer Architecture yang terdiri dari repository untuk interaksi basis data, service untuk logika dan
rest controller untuk menangani permintaan.

Terdapat beberapa library yang digunakan yaitu:

1. spring-web untuk menangani pemintaan HTTP dan response dalam aplikasi
   dilengkapi dengan dokumentasi dalam Swagger dengan library dari springdoc-openapi-starter-webmvc-ui.
2. spring-data-jpa untuk interaksi dengan database (MySQL).
3. spring-validation untuk memudahkan validasi pada property json.
4. spring-security untuk menelola authentikasi dan otorisasi yang saya gunakan dengan JWT dari library io.jsonwebtoken.
5. untuk testing memanfaatkan testcontaners untuk menggunakan container mysql secara dinamis.

Tantangan yang saya hadapi yaitu melakukan penanganan token dan mengkostumiasi respone pada
token yang tidak sesuai kategori.
