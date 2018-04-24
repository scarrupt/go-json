# go-json
Source code for my talk Read/Write JSON streams in Go at Nulab Drinking Code Singapore

## Getting Started
* Clone the repository :

    ``` git clone https://github.com/scarrupt/go-json.git ```
* Download the test data file from https://www.kaggle.com/fda/fda-enforcement-actions
* Extract the file device-510k-0001-of-0001.json into this project
* Run in a terminal ```go run main.go```
* Execute the following commands:
    - Load in memory: ```curl -vX POST http://localhost:8081/memory -d @device-510k-0001-of-0001.json --header "Content-Type: application/json"```
    - Stream: ```curl -vX POST http://localhost:8081/stream -d @device-510k-0001-of-0001.json --header "Content-Type: application/json"```


