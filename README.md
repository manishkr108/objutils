# Hi 

[![Build Status]([https://github.com/manishkr108/objutils/actions/workflows/curl](https://github.com/manishkr108/objutils/blob/main/.github/workflows/main.yml)/badge.svg?branch=main)](https://github.com/manishkr108/objutils/actions/workflows/main.yml?query=main)

[![Build Status](https://github.com/manishkr108/objutils/actions/workflows/tests.yaml/badge.svg?branch=main)](https://github.com/manishkr108/objutils/actions/workflows/tests.yaml?query=main)

[![Build Status](https://github.com/manishkr108/objutils/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/avelino/awesome-go/actions/workflows/main.yml?query=branch%3Amain)


# objutils: Go Object Utility Library

ObjUtils is a powerful and flexible Go library designed to handle operations on JSON objects, structs, and maps. It simplifies complex data manipulation tasks such as sorting keys, preserving duplicates, extracting values, and more.

## Features
## 🚀 Key Features
- **Sort JSON Keys (Supports Duplicates)**: Sorts the keys in a JSON object while preserving duplicate keys.
- **MaxValue**: Finds the maximum value in a map (supports numeric types).
- **GetAllKeys**: Extracts all keys from a map or struct.
- **GetAllValues**: Extracts all values from a map or struct.
- **CombineMaps**: Merges two maps into one.
- **Sort Struct Keys**: Sorts the keys of a struct by their field names.
- **Preserve and Sort Duplicates**: Handles maps with duplicate keys, sorts them by key, and preserves all instances of duplicates.
- **Utility Functions for Reflection**: Generic handling of maps and structs using reflection for flexible data manipulation.
## Installation

You can install this package using `go get`:

```bash
go get github.com/manishkr108/objutils
```


## API Reference
- SortJSONKeys(input string) ([]KeyValue, error)
- Sorts the keys of a JSON input string and returns the sorted key-value pairs. Supports duplicate keys.

- MaxValue(m interface{}) interface{}
- Finds the maximum value in a map of numeric types.

- GetAllKeys(m interface{}) []string
- Retrieves all keys from a map or struct and returns them in a sorted order.

- GetAllValues(m interface{}) []interface{}
- Retrieves all values from a map or struct.

- CombineMaps(m1, m2 interface{}) map[string]interface{}
- Combines two maps, with values from the second map overriding the first in case of duplicate keys.

- StructKeyShortener(s interface{}) map[string]interface{}
- Returns a map of shortened struct keys and their associated values.

## Usage
check my main.go file, there are some common use cases and examples of how to use ObjUtils in your Go project.


## Contributing
Contributions are welcome! If you'd like to contribute to ObjUtils, please follow these steps:

Fork the repository.
Create a new feature branch: git checkout -b my-feature.
Commit your changes: git commit -m 'Add a new feature'.
Push to the branch: git push origin my-feature.
Submit a pull request.
We appreciate your feedback and contributions!

