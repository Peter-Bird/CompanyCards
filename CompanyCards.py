import csv
import os

def process_csv(input_file, output_dir):
    # Ensure the output directory exists
    os.makedirs(output_dir, exist_ok=True)
    
    # Read the CSV file
    with open(input_file, mode='r', encoding='utf-8') as file:
        reader = csv.DictReader(file)
        
        for row in reader:
            # Format the output string with each key-value pair on a new line
            output_string = (
                f"Symbol: {row['Symbol']}\n"
                f"Security: {row['Security']}\n"
                f"GICS Sector: {row['GICS Sector']}\n"
                f"GICS Sub-Industry: {row['GICS Sub-Industry']}\n"
                f"Headquarters Location: \"{row['Headquarters Location']}\"\n"  # Ensure quotes around location
                f"Date added: {row['Date added']}\n"
                f"CIK: {row['CIK']}\n"
                f"Founded: {row['Founded']}\n"
            )
            
            # Define the output file name
            output_file = os.path.join(output_dir, f"{row['Symbol']}.txt")
            
            # Write to the output file
            with open(output_file, mode='w', encoding='utf-8') as out_file:
                out_file.write(output_string)
    
    print(f"Processed entries. Files saved in '{output_dir}'.")

# Example usage
process_csv('sap500Info.csv', 'output_files')

