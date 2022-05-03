<!-- PROJECT LOGO -->
<br />
<p align="center">
  
  <h3 align="center">ALTA Booking Sport Arena</h3>

  <p align="center">
    a Golang REST API with Echo and GORM
    <br />
    <br />
    <a href="https://github.com/ALTA-Booking-Sport-Arena-App/sport-arena-api/issues">Report Bug</a>
    ·
    <a href="https://github.com/ALTA-Booking-Sport-Arena-App/sport-arena-api/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Contributing](#contributing)
* [Contact](#contact)


<!-- ABOUT THE PROJECT -->
## About The Project

### Standard Naming Convention
* FOLDER = camelCase
* FILE = snake_case
* FUNCTION = PascalCase

### Feature

* Register & Login
* CRUD User
* CRUD Facility
* CRUD Category
* CRUD Venue
* CRUD Admin
* CRUD Booking & Payment

### Built With

* [Go as Programming Language](https://golang.org/)
* [Echo as Framework](https://echo.labstack.com/)
* [JWT Auth as Autentication](https://github.com/dgrijalva/jwt-go)
* [Gorm as ORM](https://gorm.io/index.html)
* [MySql as Database](https://www.mysql.com/)
* [Database Stored in RDS Cloud by Amazon Web Services](https://aws.amazon.com/id/?nc2=h_lg)

### Usage
* [Open API](https://app.swaggerhub.com/apis-docs/hafidhirsyad/sport-arena-api/1.0.0#/)
* [ERD](https://drive.google.com/file/d/15USe_PChj1mokGHdY6DdogXZxwvz11j9/view)
* [Postman Collections](https://www.getpostman.com/collections/958e8d44fda92533d2fc)

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running, follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Install Golang, Mysql, and Postman for testing
* create an `.env` file

```bash
  export APP_PORT=""
  export JWT_SECRET="S3CR3T"
  export DB_PORT="3306"
  export DB_DRIVER="mysql"
  export DB_NAME=""
  export DB_ADDRESS="127.0.0.1"
  export DB_USERNAME=""
  export DB_PASSWORD=""
```

### Installation

1. Clone the repo (in Folder htdocs)
```sh
git clone https://github.com/ALTA-Booking-Sport-Arena-App/sport-arena-api.git
```
2. Install module with get
```sh
go get
go mod tidy
```
3. Run
```sh
source .env
go run main.go
```
4. Access via url
```JS
localhost:port
```

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


<!-- CONTACT -->
## Contact

[![](https://img.shields.io/badge/LinkedIn_Haudhi-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/haudhiizza/)
[![](https://img.shields.io/badge/LinkedIn_Hafidh-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/hafidhirsyad/)
![](https://img.shields.io/badge/LinkedIn_Husnul-0077B5?style=flat&logo=linkedin&logoColor=white)

