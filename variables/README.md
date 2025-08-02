# Golang Basics: Syntax and Basic Types

## Overview

This section covers the fundamental syntax and basic types in Go. Understanding these concepts is essential for getting started with Go programming.

## 1. Basic Types in Go

- **Integer Types (`int`, `int8`, `int16`, `int32`, `int64`)**: Go provides several integer types of varying sizes (8, 16, 32, 64 bits). The default `int` type is platform-dependent (32-bit or 64-bit). Unsigned integer types (`uint`, `uint8`, `uint16`, `uint32`, `uint64`) are also available and cannot represent negative values.

- **Floating-Point Types (`float32`, `float64`)**: Go supports floating-point numbers with `float32` and `float64`. The `float64` type is the default and offers more precision.

- **String Type (`string`)**: Strings in Go are immutable sequences of bytes. They can be defined using double quotes (`"..."`) for a regular string or backticks (`` `...` ``) for a raw string literal, which does not interpret escape sequences.

- **Boolean Type (`bool`)**: Booleans in Go are either `true` or `false`. They are primarily used in conditional expressions and control structures.

## 2. Variable Declaration

- **Basic Declaration**: Use the `var` keyword followed by the variable name and type. Example: `var x int = 10`.

- **Type Inference**: Go can infer the type of a variable based on its initial value. Example: `var x = 10`.

- **Short Variable Declaration**: Use `:=` for declaring and initializing a variable in a single statement within a function. Example: `x := 10`.

## 3. Constants

- **Declaration**: Constants are declared using the `const` keyword and must be assigned a value at the time of declaration. They cannot be changed later in the program. Example: `const Pi = 3.14`.

- **Typed and Untyped Constants**: Constants can be typed or untyped. Typed constants are given a specific type like `int` or `float32`, while untyped constants have a kind and can be used in different contexts without a specific type.

## 4. Zero Values

- **Zero Values**: When variables are declared without an initial value, they are assigned a "zero value" by default. The zero value depends on the type:
    - `int` → `0`
    - `float64` → `0.0`
    - `string` → `""` (empty string)
    - `bool` → `false`

## 5. Type Conversions

- **Explicit Conversion**: Unlike some other languages, Go does not perform implicit type conversions. If you want to convert one type to another, you must do so explicitly using a conversion function. Example: `float64(x)` converts an integer `x` to a `float64`.

## Conclusion

Understanding the syntax and basic types in Go provides a strong foundation for learning more advanced topics. Mastering these concepts will help you write clear, idiomatic, and efficient Go code.
