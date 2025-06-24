# Directory File Hashing with SHA256 in Go

This Go program computes the SHA256 hash of all files in a specified directory,
with each file being processed in parallel while maintaining a loop-carried
dependence between iterations. The hash values are then combined to generate a 
final hash for the entire directory. This ensures that the order of file hashing 
is preserved across runs, giving consistent results even when the program is 
executed multiple times on the same directory.

## How it Works

1. **Directory Input:** The program expects a directory path as a command-line
   argument. It processes all the files in this directory.
   
2. **File Processing:**
   - For each file in the directory, a new goroutine is spawned to calculate 
     its SHA256 hash concurrently. 
   - The program uses channels to maintain the correct order of execution, ensuring 
     that the hashing process follows the intended sequence.
   
3. **Hashing Logic:** 
   - The hash of each file is calculated individually using the SHA256 algorithm.
   - A synchronization mechanism—using channels—is employed to make sure that
     the hashing process of each file completes before moving on to the next one.
   
4. **Final Output:** Once all files are processed, the program outputs a combined 
   SHA256 hash of the entire directory. This final hash represents a unique 
   fingerprint of the directory's contents.

## Key Concepts

- **Goroutines & Channels:** 
   - The program takes advantage of Go’s concurrency model, where each file's 
     hashing is done in parallel via goroutines. 
   - To maintain the correct sequence of operations (due to loop-carried 
     dependence), channels are used to synchronize the hashing process between 
     different iterations.
  
- **Loop-Carried Dependence:**
   - The order of file hashing matters: each file’s hash depends on the previous 
     file’s hash. This ensures the final hash is calculated correctly. If the 
     order is altered, the final hash will be different.
   
- **Directory Order:** 
   - The program relies on Go’s `os.ReadDir()` function, which returns the 
     directory entries in a consistent order. This is crucial because it ensures 
     that the program will produce the same hash every time it runs, as long as 
     the contents of the directory remain unchanged.
   - Without this consistent order, the hash result could vary on different 
     executions, even if the directory contents have not changed.

- **Synchronization with Channels:**
   - The main goroutine waits for the final iteration to be completed by 
     expecting a "ready" message on the `next` channel. This message serves 
     as a signal that the final file’s hash has been computed.
   - In the implementation, the ready message is simply a `0` sent through the 
     channel, marking the completion of the iteration and signaling that the 
     program can print the final hash.

## Example Usage

```bash
go run main.go /path/to/directory
