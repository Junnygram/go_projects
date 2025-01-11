1. **Run the program**:  
   Start the program by running:

   ```bash
   go run main.go
   ```

2. **Test SSH connection**:  
   Open a new terminal and connect to the SSH server:

   ```bash
   ssh localhost -p 2222 < main.go
   ```

3. **Test with cURL**:  
   In another terminal, send an HTTP request using the tunnel ID:

   ```bash
   curl "http://localhost:8080/tunnel?id=8723886496792943456"
   ```

4. **If you encounter issues reconnecting**:  
   Open the `~/.ssh/known_hosts` file using a text editor:
   ```bash
   nano ~/.ssh/known_hosts
   ```
   Remove the old key entry for `localhost`, then regenerate the key and run the program again.
