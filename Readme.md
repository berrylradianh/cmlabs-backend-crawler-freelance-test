# cmlabs-backend-crawler-freelance-test

## Description
Repository yang digunakan untuk Pre-assessment Test  Intermediate Back-end Developer cmlabs

## Local Installation

1. Clone this project
    ```
    git clone https://github.com/berrylradianh/cmlabs-backend-crawler-freelance-test.git
    ```

2. Copy `.env.example` to `.env`
    ```
    cp .env.example .env
    ```
3. Configure environment variables for database connection
    ```
    DB_CONNECTION=mysql
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=cmlabs_crawler
    DB_USERNAME=root
    DB_PASSWORD=
    ```

4.  Run the application
    ```
    go run main.go
    ```