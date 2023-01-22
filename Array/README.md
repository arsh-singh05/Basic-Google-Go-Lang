## Arrays
Arrays in Go are used to store a fixed-size collection of elements.

Go arrays have a fixed size and can store elements of the same type.

Arrays in Go are used to store a fixed-size collection of elements of the same type. The size of an array is set when it is created and cannot be changed afterwards.

To create an array in Go, use the [size]type syntax, where size is the number of elements in the array and type is the type of the elements. For example, arr := [3]int{1, 2, 3} creates an array of integers with 3 elements, initialized to the values 1, 2 and 3.

Arrays can be accessed and modified using their index, which starts at 0

Go also has a more flexible alternative to arrays called slices, which can be used when the size of the collection is not known or needs to change during the program execution.

In summary, arrays in Go provide a way to store a fixed-size collection of elements and can be useful for situations where the size of the collection is known in advance. However, if the size needs to change dynamically or is unknown, slices might be a better option.
