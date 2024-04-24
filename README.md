"# GOLANG_AS_3_2-3"  "# GOLANG_AS_3_2-3" 
OK THERE IS 3 EXERCISES 
1 JWT IS MADE BY USING COOKIE IF IN COOKIE HAS AUTH IN NEWS SITE WILL OPEN OTHERWISE WILL SHOW JUST U ARE NOT AUTORIZED 

2 The downloadFile function is defined. It takes four parameters: url, filePath, wg, and ch.
url is the web address of the file to download.
filePath is the location where the downloaded file will be saved.
wg is a WaitGroup, a tool for waiting until all downloads are finished.
ch is a channel used for communicating errors back to the main function.
Inside downloadFile:
It starts by deferring the call to wg.Done(), which signals that this download is finished.
It creates a new file with the given filePath.
Then it tries to download the file from the provided url.
After downloading, it copies the content of the downloaded file into the newly created file.
Finally, it sends any encountered errors to the error channel (ch).
In the main function:
An array of URLs to download is defined.
A channel for errors (errorChannel) is created.
A WaitGroup (wg) is initialized.
A loop iterates through each URL in the array:
For each URL, a goroutine (a lightweight thread) is started to download the file concurrently.
The wg.Add(1) call increments the WaitGroup to keep track of the number of downloads.
Another goroutine is started to wait until all downloads are completed and then closes the error channel.
Finally, a loop reads from the error channel, printing any errors encountered during the downloads.

3.1This code is like a calculator. It adds, subtracts, multiplies, and divides numbers. It starts by printing "Starting the program..." Then, it does some math: adding, subtracting, multiplying, and dividing two numbers. After each math operation, it prints the result. If there's a problem with dividing (like trying to divide by zero), it prints an error message instead of the result. Finally, it prints "Program finished."

3.2This code is similar to the first one, but it's like a bit more organized. It still does the same math operations: adding, subtracting, multiplying, and dividing two numbers. But this time, after each operation, it prints what it did (like "I added these numbers and got this result"). It also handles errors better, so if there's a problem with dividing (like trying to divide by zero), it prints a message saying what went wrong. Finally, it prints "Program finished." when it's done.
