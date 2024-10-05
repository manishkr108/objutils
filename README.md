# objutils

# objutils: Go Object Utility Library

ObjUtils is a lightweight utility library for working with JSON objects in Go. It provides functions for sorting keys, preserving duplicates, and more, making it easier for developers to manipulate JSON data and other key-value structures.


Here’s a more detailed and comprehensive README template that highlights all the features of your objutils library. It includes examples and explanations of how each function can be used, along with installation instructions, contribution guidelines, and more.

ObjUtils - JSON and Object Utility Library for Go

ObjUtils is a powerful and flexible Go library designed to handle operations on JSON objects, structs, and maps. It simplifies complex data manipulation tasks such as sorting keys, preserving duplicates, extracting values, and more.

## Features
## 🚀 Key Features

- **Sort JSON Keys (Supports Duplicates): Sorts the keys in a JSON object while preserving duplicate keys.
- **MaxValue: Finds the maximum value in a map (supports numeric types).
- **GetAllKeys: Extracts all keys from a map or struct.
- **GetAllValues: Extracts all values from a map or struct.
- **CombineMaps: Merges two maps into one.
- **Sort Struct Keys: Sorts the keys of a struct by their field names.
- **Preserve and Sort Duplicates: Handles maps with duplicate keys, sorts them by key, and preserves all instances of duplicates.
- **Utility Functions for Reflection: Generic handling of maps and structs using reflection for flexible data manipulation.
## Installation

You can install this package using `go get`:

```bash
go get github.com/manishkr108/objutils
