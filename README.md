# BE-Mini-Project_Delosaqua
This is the technical test for applying internship at Delosaqua. Create the CRUD, consume api, unit test, and using clean code, design pattern, e.g. Clean Architecture

# DELOSAQUA - Aquafarm Management Applications Prototype REST API using Go Language

This application provides CRUD APIs for managing farms and ponds for aquaculture/fish farms. 

The system allows fish farm owners to track information about their different farms and the ponds located within each farm. Key information such as location, capacity, construction details, stocking densities, and inventory/yields can be stored for farms and associated ponds.

The API server provides the following endpoints:

**Farms**

- Create new farms
- View list of all farms
- View details of a single farm 
- Update information for existing farms
- Delete farms using SoftDelete

**Ponds**

- Create new ponds associated with a farm 
- View list of all ponds
- View details of a single pond
- Update information for existing ponds 
- Delete ponds using SoftDelete

**Relationships**

The system enforces a 1:N relationship between farms and ponds. A pond can only be associated with one farm. The APIs validate this relationship when creating or updating entities.

**Statistics**

Usage statistics are tracked for each API endpoint, containing number of requests and unique clients.

The system is designed with scalability and maintainability in mind. It uses a Dockerized MySQL database that can handle larger datasets as the system scales to additional farms. The Golang application also incorporates design patterns such as repository interfaces to minimize coupling.

## Prerequisites

- Golang 1.21 or higher
- Docker 
- Postman

## Environment Setup

### Golang

Install the latest Golang per https://golang.org/doc/install

Set GOPATH environment variable to point to your workspace.

### Docker 

Install Docker engine and docker-compose per https://docs.docker.com/get-docker/

## Running the Application
To start the Dockerized Postgresql database and run Golang Apps:
```docker-compose up -d```

The API server will start on port 3000.

## API Routes

The API routes are versioned:
```http://localhost:3000/v1```

**Farms**

- Get all farms - `GET /farms-delosaqua/viewAll`
- Create farm - `POST /farms-delosaqua/manage/store`

**Ponds**

- Get ponds - `GET /ponds-delosaqua/viewAll` 
- Create pond - `POST /ponds-delosaqua/manage/store`

**Metrics**

- `GET /get-stats` - Return API metrics

## API Documentation
You can access the Postman API Documentation [here](https://documenter.getpostman.com/view/20251635/2s9YeAAZkY)

