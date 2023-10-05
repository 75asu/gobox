# CLI Calculator

A command-line calculator application for performing basic arithmetic operations.

### Features

- **Addition:** Add two numbers.
- **Subtraction:** Subtract two numbers.
- **Multiplication:** Multiply two numbers.
- **Division:** Divide two numbers. Handles division by zero gracefully.

### Usage

To use the CLI calculator app, follow these steps:

1. **Build the Application:**

   ```shell
   go build calc.go
   ```

2. **Run the Calculator:**

   Run the `calc` executable with one of the following command-line flags to perform the desired operation:

   - `-add`: Addition
   - `-sub`: Subtraction
   - `-mul`: Multiplication
   - `-div`: Division

   For example, to perform addition:

   ```shell
   ./calc -add
   ```

   You will be prompted to enter two numbers, and the calculator will display the result.

3. **Example:**

   ```shell
   gitpod /workspace/gobox/3-calculator (main) $ go build calc.go 
   gitpod /workspace/gobox/3-calculator (main) $ ./calc -add
   Enter 1st Number: 
   4
   Enter 2nd Number: 
   6
   Addition: 10 
   ```

### Supported Operations

- `-add`: Perform addition.
- `-sub`: Perform subtraction.
- `-mul`: Perform multiplication.
- `-div`: Perform division. Handles division by zero gracefully.

### Error Handling

- If you attempt to divide by zero, the calculator will display an error message: "Error: Can't Divide By Zero."

Feel free to use this calculator for basic arithmetic operations in your command-line environment.

