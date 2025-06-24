# Parallel File SHA-256 Hasher in Go

This simple Go program computes SHA-256 hashes of all files in a given directory
**in parallel**, speeding up the hashing process by utilizing goroutines.

## How It Works

- Takes a directory path as a command-line argument.
- Reads all entries (files) in the directory.
- For each file (excluding subdirectories), it launches a separate goroutine to:
  - Open the file
  - Compute its SHA-256 hash
  - Print the filename alongside its hash in hex format
- Uses a `sync.WaitGroup` to ensure the main program waits until all file hashes
  are computed before exiting.

## Key Concepts

- **Loop-level parallelism:** Instead of processing files one-by-one 
sequentially, each file is processed concurrently in its own goroutine.
- **No task dependencies:** The hash computation for one file is independent of
  others, making parallelism straightforward and safe.
- **WaitGroup synchronization:** Ensures main goroutine waits until all hashing
  goroutines finish.

Because there is no dependence between tasks, the result of computing the hash
for one file does not affect the computation of another. 
This makes the loop-level parallelism pattern an ideal choice here.

## Usage

```bash
go run main.go /path/to/directory
