## File handling in Go 
It is used to read and write data to files.

Go provides a built-in package called os which has functions to work with files, like Open, Create, Read, Write, Close.

File handling in Go allows you to read and write data to files on disk. Go's standard library provides functions for reading and writing files.

To read a file, you can use the ioutil.ReadFile() function. This function reads the entire contents of a file and returns it as a byte slice.

To write to a file, you can use the ioutil.WriteFile() function. This function writes data to a file and creates the file if it does not exist.

Alternatively, you can also use the os.Open(), os.Create(), os.OpenFile() functions to open a file, and file.Read(), file.ReadAt(), file.Write(), file.WriteAt() functions to read and write data to the file.

Go also provides a higher-level package called ioutil which has functions to read and write files more easily.

Go has built-in support for reading and writing CSV, JSON and XML formatted files as well.

In summary, File handling in Go allows you to read and write data to files on disk, it includes various functions such as ioutil.ReadFile(), ioutil.WriteFile(), os.Open(), os.Create(), os.OpenFile(), file.Read(), file.ReadAt(), file.Write(), file.WriteAt(). They are a fundamental concept for creating programs that can interact with the file system.
