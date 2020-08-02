# Ranking Of Web Domains

---

###### Implemented a code to read a CSV file (aka customers.csv) and return sorted data structure of email domains along with the number of customers with e-mail addresses for each domain.  Any errors should be logged (or handled). Performance matters (this is only ~3k lines, but *could* be 1m lines or run on a small machine).

---

## Run

```bash
go run main.go path/of/file.csv
```
## Build

```bash
go build
./interview path/of/file.csv
```
## Run tests

```bash
go test ./... -v
```
## Run tests with coverage output

```bash
go test ./... -v -coverprofile=unittest.out
```

## Coverage output  in HTML

```bash
go tool cover -html=unittest.out
```

## Running with Docker


### Build

```bash
docker build -t ranking_of_web_domains .
```
### Running
```bash
docker run --rm ranking_of_web_domains
```

## Files

- main.go - **Entry point of the program**
- *customerimporter package*
--  interview.go - **Business logic of the program**
-- interview_test.go - **Unit Tests of business logic**

## Improvements
- Make docker image get **file.csv** from a URL/volume
- Increase the test coverage
- Use regex to check a valid email on **getDomainFromEmail** func
- Improve the way the tests are written to mock some func when necessary.
-- An example is when *TestGetListOfDomains* calls **getListOfDomains** and **getListOfDomains** calls **getDomainFromEmail** internally. The func **getDomainFromEmail** should be mocked to not interfere on **getListOfDomains** unittest.

## 3-party libraries
###### I tried to not use any 3-party library to show my skills, but I believe is a good thing to use some specific libraries to help us to avoid of not rewrite the wheel.

Not sure if these libraries would help me in this project, but it's good to know that they exist.

- https://github.com/artonge/go-csv-tag
- https://github.com/stretchr/testify
- https://github.com/wk8/go-ordered-map
