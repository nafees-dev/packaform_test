
# Packaform | Order Details

Monolithic Architecture is used in this project. There are two folders web and api and one migration.py file which will be use to migrate all the data from csv file (which is in test_data folder) to postgres database. 
The Database name is mentioned in migration.py file





## Tech Stack 

**Client:** Vue JS, Bootstrap

**Server:** Golang, Python





## Installation

Install and Setup Vue.js Project

```bash
  cd web
  npm i
  npm run serve
```

Run Migration.py file, First install the required library
```bash
pip insall -r requirement.txt
```
OR directly install the package using below command
```bash
pip install psycopg2
```
then run below command
```bash
python3 migration.py
```
Now Run the Go Server
```bash
  cd api
  go run main.go
```


## Application Page
URL: 
http://localhost:8080/order
![App Screenshot](https://raw.githubusercontent.com/nafees-dev/packaform_test/master/screenshot1.png)

