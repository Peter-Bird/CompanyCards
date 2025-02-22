# CompanyCards CSV Processor

This project processes a CSV file containing company information and generates individual text files for each record. Each output file contains neatly formatted details about a company, making it easy to access and review specific data.

## Overview

The project reads a CSV file with the following columns:
- **Symbol**
- **Security**
- **GICS Sector**
- **GICS Sub-Industry**
- **Headquarters Location**
- **Date added**
- **CIK**
- **Founded**

For each row in the CSV file, the script creates a text file named after the company's symbol (e.g., `AAPL.txt`). Each file contains the company's details, with each field formatted on a new line.

## Project Structure

- `CompanyCards.py`: The Python script that processes the CSV and writes the output files.
- `sap500Info.csv`: Example CSV file containing company information (this file should be provided by the user).
- `output_files/`: The directory where the generated text files are stored.
- `README.md`: This file.

## Features

- **CSV Parsing:** Reads and parses CSV files with headers.
- **Dynamic File Generation:** Automatically creates an output directory (if it doesn't exist) and generates text files for each company.
- **Formatted Output:** Each text file contains the company's information in a clear, readable format.
- **Error Handling:** Basic error handling to manage file I/O and CSV parsing issues.

## Getting Started

### Prerequisites

- **Python 3.6+**: The script uses Python's built-in modules, so no additional packages are required.
- A CSV file with the required columns (e.g., `sap500Info.csv`).

### Running the Python Script

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Place your CSV file:**

   Ensure that your CSV file (e.g., `sap500Info.csv`) is in the same directory as `CompanyCards.py` or update the file path accordingly.

3. **Run the script:**

   ```bash
   python CompanyCards.py
   ```

   The script will process the CSV file and generate individual text files in the `output_files` directory.

### Golang Version

For users who prefer Golang, a translated version of the script is also provided. The Golang code replicates the functionality of the Python script.

#### Running the Golang Version

1. **Install Golang:**

   Download and install Golang from [golang.org](https://golang.org/).

2. **Build and Run:**

   ```bash
   go run main.go
   ```

   Ensure that the input CSV file (`sap500Info.csv`) is in the same directory or update the file path in the code accordingly.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your improvements.

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

For any questions or suggestions, please open an issue in the repository or contact the maintainer.

---

This project aims to simplify the process of handling and reviewing company data by automating the generation of easy-to-read text files from a CSV source. Enjoy!