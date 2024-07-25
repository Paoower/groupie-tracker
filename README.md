# Groupie Tracker

Groupie server is a fully fonctionning web app working on a Golang server. It uses HTML template to render client front-end. The information from the groupie API is taken and displayed.

## Project Architecture

- `datatypes/`
    - `structs.go`
- `handlers/`
    - `artist.go`
    - `artists.go`
    - `error.go`
    - `filters.go`
    - `home.go`
    - `searchBar.go`
- `json/`
    - `artists.json`	
    - `dates.json`	
    - `locations.json`	
    - `relation.json`
- `templates/`
    - `scripts/`
        - `artist_block.html`
        - `artist.html`
        - `artists.html`
        - `error.html`
        - `index.html`
- `utils/`
    - `artists.go`
    - `locations.go`
    - `readjson.go`
    - `variables.go`
- `go.mod`
- `main.go`
- `README.md`

## Usage

To clone the repository:

```bash
git clone https://github.com/itsliamine/groupie-tracker.git
```

To run the Program:
```bash
cd groupie-tracker

go run . 
```