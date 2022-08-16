# padelcuc-shifts
Created for personal purposes, with this web-scrapper we can check the available shifts (in a 4 day period) for the two padel fields in Carmen de Areco.

## About
- You will notice that this repository has **two directories**: `go` and `py`, because i wanted to replicate the same code in two very different languages as Golang and Python to check some metrics, like the speed of execution, lines of code (being the same) , etc.

## Usage
### Golang
- `cd /padelcuc-shifts/go`
- `go run . -field={field}`

Where the `field` flag could be `blindex` or `cerrada`

Example: `go run . -field=blindex`

### Python
- `cd /padelcucshifts/py`
- `python3 main.py {field}`

Example: `python3 main.py blindex`


