# ToW

<!-- ABOUT THE PROJECT -->
## About The Project
This Repo houses the version 1 of Tree of wally.

### Built With
1. [Golang](https://go.dev)
2. [Gin-gonic](https://github.com/gin-gonic/gin)
3. [PostgreSQL](https://www.postgresql.org/)
4. [GORM](https://gorm.io/)
5. REST
6. [Mockery](https://pkg.go.dev/github.com/knqyf263/mockery) for testing


<!-- GETTING STARTED -->
## Getting Started
Clone the project from GitHub using the command
```sh
git clone git@github.com:tejiriaustin/ToW.git
```

### Prerequisites

* Go
    ```sh
    https://go.dev
    ```

### Starting the project

1. run command
    ```sh
    go mod tidy
    ```
   to index files and dependencies
2. Create file and name `.env`
3. Contact the admin of this project to get your personal credentials
4. Run command to start the application with docker
    ```sh
   make service
   ```
   
5. If you have Go installed, you can also use
    ```sh
   make api
   ```

6.Test on postman

## Routes avaliable
- Create Regular Account
```
    URL: {{url}}/customer
    Body: {
      "first_name": "Tejiri",
      "last_name": "Dev",
      "phone": "123-456-1234"
      }
```

- Create Admin Account
```
    URL: {{url}}/customer
    Body: {
      "first_name": "Tejiri",
      "last_name": "Dev",
      "phone": "123-456-1234"
      }
```

- Freeze Customer Account
```
    URL: {{url}}/admin/freeze/:accountId
    Method: PUT
```

- Trigger Data Income Distribution
```
    URL: {{url}}/admin/issue-data-income
    Method: POST
```

- Set Minimum Follow Income
```
    URL: {{url}}/admin/set-minimum-follow-spend
    Method: PUT
```

- Customer Subscription
```
    URL: {{url}}/customer/subscribe
    Method: POST
```

- Invest
```
    URL: {{url}}/customer/invest
    Method: POST
```

[Golang-URL]: https://go.dev 
