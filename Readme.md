# Introduction

Use Agently to generate code by calling ChatGPT based on requirements.

## Tech Stack

- Python
- Go
  - gin
  - gorm
- Vue

## Steps

### Install Dependencies

- `pip install -U Agently`
- `pip install -U openai httpx typer loguru`

### Generate Code

- `python3 ./backend.py YOUR_PROJECT_NAME`

The program will read `prd.txt` as the requirements and generate the following files in the `YOUR_PROJECT_NAME` directory:

> `prd.md`: Requirements document  
> `*.go`: Program code  

### Fine-Tuning

- Modify the database connection information in `db.go`.
- Modify the listen port in `main.go`.
- Navigate to the `YOUR_PROJECT_NAME` directory and adjust the code as needed, as the output from the LLM may have some issues.

### Run the Program

Navigate to the `YOUR_PROJECT_NAME` directory and execute the following commands:

> `go mod tidy`  
> `go build`  
> `./YOUR_PROJECT_NAME`
