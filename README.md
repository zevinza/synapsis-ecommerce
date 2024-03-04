
# Simple Online Shop 

This Project is created to make a RESTful API for online shop application
written with go programming language and go-fiber framework.

before running this application, make sure you have this Prequisites: 
- Docker Compose
- Database Viewer such as Navicat

## How to Run
1. Run Docker Compose
```bash
docker-compose up -d
```
2. Run Go fiber API
```bash
docker-compose exec go go run .
```
3. Visit swagger documentation at 
```
localhost:8080
```
## Authorization
There's two type of Authorization, Access Token and Header Token.
to get access token, you can login with this account,
| Username | Password     | Role                |
| :-------- | :------- | :------------------------- |
| `admin@mail.com` | `password` | `admin` |
| `armadamuhammads@mail.com` | `password` | `user`|

Header Token is used pre-login action, please fill 