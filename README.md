# du
DUSocial backend repository.

## Getting Started
### Prerequisites

- Go >= 1.22.1
- A MongoDB Database

### Installation
1. Clone the repository
```sh
git clone https://github.com/gdscduzceuniversity/du.git
```

2. Install dependencies
```sh
go get
```

3. Run the mongo container

```bash
docker run -d -p 27017:27017 --name mongodb mongo
```

### environment variables

```bash
export MONGO_URI=mongodb://localhost:27017
```

or rename the `example.env` file to `du.env` and just edit the file with connection url.

### Usage
```sh
go run main.go
```
Open [localhost:8080/ping](http://localhost:8080/ping) to check if the server is running.

## License
Distributed under the GPL-3.0 License. See [LICENSE](/LICENSE) for more information.