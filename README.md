# Roll Call

Roll Call gathers and saves meta data relating to a WordPress users site specific accounts.

![Rollcall](rollcall.webp)

## Prerequisites

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

Creation of a variables file with the following values as per your environment:

```go
var (
servers = []string{/* list of servers to test against */}
prodpack = []string{/* Server, Path, and User values for the WordPress Production environment */}
testpack = []string{/* Server, Path, and User values for the WordPress Test environment */}
)
```
## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Windows & Mac:

```bash
go build -o [name] .
```

### Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

```bash
./[program name]
```

## Output

Outputs a file named *compendium.csv*. This file contains a complete list of all WordPress users, any sites they have an account on, their role for that site, and last login time, if found. Login timestamps are stored in Unix epoch time.

## License

Code is distributed under [The Unlicense](https://github.com/nausicaan/free/blob/main/LICENSE.md) and is part of the Public Domain.
