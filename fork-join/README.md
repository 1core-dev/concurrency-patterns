# Go Code Depth Analyzer

This tool scans a given directory recursively to find the Go source file  
with the deepest level of code block nesting, based on `{}` bracket pairs.  
It uses a **fork/join** concurrency pattern to efficiently process multiple  
files in parallel.

## Purpose

- Analyze `.go` files in a directory tree.
- Detect the file with the deepest nested `{}` blocks.
- Print the file path and nesting depth.

## How It Works

### Fork/Join Pattern

- **Fork**: A goroutine is started for each `.go` file to calculate the  
  nesting depth using `deepestNestedBlock()`.
- **Join**: A separate goroutine collects all results from a shared channel  
  and tracks the file with the highest nesting level.

### Nesting Depth Calculation

- Each `{` increases nesting level by 1.
- Each `}` decreases nesting level by 1.
- Maximum depth is tracked per file.

## main() Execution Flow

1. Reads the root directory from the first CLI argument.
2. Creates a results channel and `WaitGroup`.
3. Uses `filepath.Walk()` to iterate over all files.
4. For each `.go` file found:
   - Logs the discovery (optional for debugging).
   - Starts a goroutine to analyze nesting depth (fork).
5. Starts the join goroutine using `joinResults()`:
   - Listens on the shared results channel.
   - Keeps track of the file with the deepest nesting.
6. Waits for all file-processing goroutines to complete using `WaitGroup`.
7. Closes the results channel once all tasks are done.
8. Receives the final result from the join goroutine.
9. Prints the file path and the nesting depth.