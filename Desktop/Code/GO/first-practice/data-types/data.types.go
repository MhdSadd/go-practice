package datatype

import "fmt"

func GoDataTypes() {
	//*** STRINGS
	StringDataType := "Any data sorrounded by double quotes only"

	//*** BOOL
	isLoggedIn := false

	/*
		INTEGERS: golang has many forms of integer types depending on precision and platform you're on;
			****THE INTs (whole numbers)****
			int: The int type represents signed integers and is implementation-dependent in terms of its size. It's typically 32 or 64 bits, depending on the platform.
			int8: Represents a signed 8-bit integer with values ranging from -128 to 127.
			int16: Represents a signed 16-bit integer with values ranging from -32,768 to 32,767.
			int32: Represents a signed 32-bit integer with values ranging from approximately -2 billion to 2 billion.
			int64: Represents a signed 64-bit integer with a wide range of values.

			***THE UINTs (positive whole numbers)***
			uint: The uint type represents unsigned integers and is implementation-dependent, similar to int. It's typically 32 or 64 bits, depending on the platform.
			uint8: Represents an unsigned 8-bit integer with values ranging from 0 to 255.
			uint16: Represents an unsigned 16-bit integer with values ranging from 0 to 65,535.
			uint32: Represents an unsigned 32-bit integer with values ranging from 0 to approximately 4.3 billion.
			uint64: Represents an unsigned 64-bit integer with a wide range of values.
			uintptr: This type is an unsigned integer type that's large enough to hold the bit pattern of a pointer. Its size varies based on the platform (32-bit or 64-bit).

			The default type, as mentioned earlier, is *int* for integer literals when no specific type is provided.
			For example, if you write x := 42, x will be of type int. which basically just aliases the platform you're on 32 or 64.
			If you need a specific-sized integer, it's better to explicitly specify the type, such as int32 or int64, to ensure consistent behavior across different systems.
	*/
	age := 30
	debit := -10000

	/*
		FLOATS: for representing decimal/floating point numbers
		float32: The float32 type is a 32-bit floating-point number, also known as single-precision.
		float64: The float64 type is a 64-bit floating-point number, also known as double-precision. It offers a wider range and higher precision compared to float32.
	*/

	trafficTicket := 123.10

	/*
		COMPLEXS: for representing numbers with real and imaginary parts
		complex64: uses 32 bit for each number part while
		complex128: uses 64 bit for each number part
	*/
	complexNumber := complex(5, 1)
	realComplexPart := real(complexNumber)
	imaginaryComplexPart := imag(complexNumber)

	fmt.Printf("***DATA TYPES IN GOLANG INCLUDE***\n%T : %s\n", StringDataType, StringDataType)
	fmt.Printf("-> %T : %v\n", isLoggedIn, isLoggedIn)
	fmt.Printf("-> [unsigned] %T : %d\n", age, age)
	fmt.Printf("-> [signed] %T : %d\n", debit, debit)
	fmt.Printf("-> %T : %f\n", trafficTicket, trafficTicket)
	fmt.Printf("-> %T : %f\n", complexNumber, complexNumber)
	fmt.Printf("real and imaginary part of the complex number are %f, %f\n\n", realComplexPart, imaginaryComplexPart)
}
