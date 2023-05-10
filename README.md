# Simple GraphQL MongoDB API

This is a simple web API that performs basic CRUD in Go 
Using GraphQL and MongoDB

### How to run the application

- Clone the repository
- Navigate to the project directory: `cd simpleMongoAPI`
- Create an `.env` file in the root directory of the project.
- Populate the `.env` file with the required environment variables.
Here's an example: 
`PORT=:8080
DB_URI={mongodb connection string}`

#### Running using Docker

- run `docker compose up`
- to stop the application run `docker compose down`

#### Running Locally

- run `go build -o {binaryName} .\cmd\app\main.go` to build the binary
- run `.\{binary name}` to run the binary (eg .\simpleAPI)

### Testing

- run `go test -v ./...` to run the unit tests
- run `go test -v ./... -coverprofile=c.out` to create the cover profile
- run `go tool cover -html c` to see the
  coverage of the unit tests

## Endpoints

### Create Student

This Endpoint Creates a Student

#### Request
    mutation {
      createStudent (
        student: {
          firstname: "Daniel"
          lastname: "Riccardo"
          year: 1
        }
      ) {
        status
        data {
          id
          firstname
          lastname
          year
        }
        message
      }
    }

#### Response

    {
      "data": {
        "createStudent": {
          "status": "Success",
          "data": {
              "id": "645776de7e33a836bc282e14",
              "firstname": "Daniel",
              "lastname": "Riccardo",
              "year": 1
            },
          "message": "Students Created Successfully"
        }
      }
    }

### Get All Students

This Endpoint Returns all the Students

#### Request

    query {
      getAllStudents {
        status
        data {
          id
          firstname
          lastname
          year
        }
        message
      }
    }
#### Response

    {
      "data": {
        "getAllStudents": {
          "status": "Success",
          "data": [
            {
              "id": "64548462893590bd8ef53a75",
              "firstname": "Charles",
              "lastname": "Leclerc",
              "year": 1
            },
            {
              "id": "6454d4dd088ecfbcf32718c6",
              "firstname": "Lewis",
              "lastname": "Hamilton",
              "year": 2
            }
          ],
          "message": "Students Queried Successfully"
        }
      }
    }

### Get Specific Student

This Endpoint Returns a specific Student

#### Request

    query {
      getStudent (studentID:"64548462893590bd8ef53a75") {
        status
        data {
          id
          firstname
          lastname
          year
        }
        message
      }
    }

#### Response

    {
      "data": {
        "getStudent": {
          "status": "Success",
          "data": {
              "id": "64548462893590bd8ef53a75",
              "firstname": "Charles",
              "lastname": "Leclerc",
              "year": 1
            },
          "message": "Students Queried Successfully"
        }
      }
    }

### Update Student

This Endpoint Updates Data of a Student

#### Request

    mutation {
      updateStudent (
        student: {
          id: "6457738effef30308908717c"
          firstname: "Daniel"
          lastname: "Riccardo"
          year: 5
        }
      ) {
        status
        data {
          id
          firstname
          lastname
          year
        }
        message
      }
    }

#### Response

    {
      "data": {
        "updateStudent": {
          "status": "Success",
          "data": {
              "id": "6457738effef30308908717c",
              "firstname": "Daniel",
              "lastname": "Riccardo",
              "year": 5
            },
          "message": "Students Updated Successfully"
        }
      }
    }

### Delete Student

This Endpoint Deletes a Student

#### Request

    mutation {
      deleteStudent (studentID:"6457738effef30308908717c") {
        status
        message
      }
    }

#### Response

    {
      "data": {
        "deleteStudent": {
          "status": "Success",
          "message": "Student Deleted Successfully"
        }
      }
    }

### Search Student

This Endpoint can be used to search a student
based on their firstname or lastname and sort
the results. This endpoint give a paginated
response

#### Request

    query {
      searchStudent (
        searchString: "a"
        sortBy: {
          column: "lastname"
          direction: ASC
        }
        pagination: {
          page: 1
          pageSize: 2
        }
      ) {
        status
        data {
          totalElements
          data {
            id
            firstname
            lastname
            year
          }
        }
       message
      }
    }

#### Response

    {
      "data": {
        "searchStudent": {
          "status": "Success",
          "data": {
            "totalElements": 0,
            "data": [
              {
                "id": "6454d4dd088ecfbcf32718c6",
                "firstname": "Lewis",
                "lastname": "Hamilton",
                "year": 2
              },
              {
                "id": "64548462893590bd8ef53a75",
                "firstname": "Charles",
                "lastname": "Leclerc",
                "year": 1
              }
            ]
          },
          "message": "Students Queried Successfully"
        }
      }
    }

