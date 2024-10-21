# compression-tool

File Compression Tool CLI

A simple command-line tool built with Go and Cobra CLI to compress and decompress files or folders efficiently.
Features

    Compress Files/Folders: Supports compressing both individual files and entire directories.
    Decompress Files: Extract compressed files to their original form.

    Command-line Interface: Easy to use with intuitive commands for compression and decompression.


Here’s a structured README.md for your file compression tool:
File Compression Tool CLI

A simple command-line tool built with Go and Cobra CLI to compress and decompress files or folders efficiently.
Features

    Compress Files/Folders: Supports compressing both individual files and entire directories.
    Decompress Files: Extract compressed files to their original form.
    Fast and Lightweight: Leverages Go’s concurrency features for performance.
    Command-line Interface: Easy to use with intuitive commands for compression and decompression.


Usage
1. Compress a File or Folder

        go run main.go compress --filename <filename> 

        --filename: Path to the file to compress

    <!-- or -->

        go run main.go compress --foldername <foldername>
        --foldername: Path to the folder to compress

Example:

    go run main.go --foldername input

2. Decompress a File

        go run main.go decompress --filename <filename> 

        --filename: Path to the file to decompress

    <!-- or -->

        go run main.go decompress --foldername <foldername>
        --foldername: Path to the folder to decompress or extract the zipoed folder

Example:

    go run main.go --foldername input
