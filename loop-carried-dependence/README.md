# Directory File Hashing with SHA256 in Go

This Go program computes the SHA256 hash of all files in a specified directory,
with each file being processed in parallel while maintaining a loop-carried
dependence between iterations. The hash values are combined to generate a 
final hash for the entire directory.

## How it Works

1. **Directory Input:** The program expects a directory path as a command-line
   argument.
   
2. **File Processing:** For each file in the directory:
   - A goroutine is spawned to compute the SHA256 hash of the file.
   - A synchronization mechanism is used to ensure that the order of hashing
     is maintained (due to loop-carried dependence).
   
3. **Hashing Logic:** 
   - The hash of each file is calculated using SHA256.
   - A channel is used to coordinate the sequential processing of the hash.
   
4. **Final Output:** The program outputs a combined SHA256 hash of all files
   in the directory.

## Key Concepts

- **Goroutines & Channels:** Each file is hashed concurrently using
  goroutines, with channels used to synchronize their execution in the
  correct order.
  
- **Loop-Carried Dependence:** The hashing order matters; each file's hash
  depends on the previous file's hash being computed first.
  
- **Directory Order:** The program relies on Go’s `os.ReadDir()` function,
  which returns directory entries in a consistent order. This is crucial
  for ensuring that the final hash result remains the same every time the
  program is run, even if the directory is re-scanned. Without this ordered
  listing, the hash result could change on every run, even if the directory
  contents haven't changed.

## Example Usage

```bash
go run main.go /path/to/directory
