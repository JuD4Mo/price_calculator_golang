# Price Calculator (Go)

> Small Go utility that reads a list of prices from `prices.txt`, calculates tax-included prices for several tax rates, and writes the results to JSON files.

## Requirements
- Go 1.20+ installed

## Project structure
- `main.go` — orchestrates running the job for multiple tax rates.
- `prices.txt` — input file with one price per line (e.g. `9.99`).
- `fileManager/` — handles reading lines and writing JSON output.
- `prices/` — contains logic to compute tax-included prices.
- `conversion/` — helper to convert strings to floats.

## Usage
From the project root (PowerShell):

```powershell
# Run the program
go run .

# View a result file (example)
Get-Content .\result_0.json
```

The program generates files named `result_*.json` (for example `result_0.json`, `result_70.json`) containing the job output.

## What it does
- Reads `prices.txt` and converts each line to a `float64`.
- For each tax rate (defined in `main.go`) it calculates the tax-included price and writes a JSON file with the job data.

---