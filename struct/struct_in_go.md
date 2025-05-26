# 📘 Structs in Go

## 🧱 What is a Struct?

A **struct** in Go is a composite data type (also known as a _record_ or _object_ in other languages) that groups together variables under a single name. These variables are called **fields**.

### Syntax:

```go
type StructName struct {
    Field1 Type1
    Field2 Type2
    ...
}
```

## ✅ Key Features:

- Used to create user-defined data types.
- Fields can be accessed using dot notation.
- Structs can be passed by value or by reference (using pointers).
- Zero values apply: integers default to `0`, strings to `""`, etc.

---

## 📦 Code Example

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    p := &v
    p.X = 1e9
    fmt.Println(v)
}
```

---

## 🔍 Step-by-Step Explanation

### 1. `type Vertex struct { X int; Y int }`

Defines a struct type named `Vertex` with two integer fields: `X` and `Y`.

### 2. `v := Vertex{1, 2}`

Creates a value of type `Vertex` with `X=1` and `Y=2`.

### 3. `p := &v`

Creates a **pointer** to `v`. Now `p` holds the memory address of `v`.

### 4. `p.X = 1e9`

This line is **syntactic sugar** in Go. When you write `p.X`, Go automatically dereferences the pointer `p` and accesses the `X` field of the underlying struct.  
`1e9` is scientific notation and equals `1,000,000,000` (a float64, but it's implicitly cast to int here).

Effectively, this line does:

```go
(*p).X = 1000000000
```

### 5. `fmt.Println(v)`

Prints the updated struct:

```
{1000000000 2}
```

---

## 💡 Notes:

- Go structs are mutable, and using pointers lets you modify the original struct directly.
- This is efficient and commonly used for method receivers or shared state.

---

## 📌 Summary

- **Structs** in Go are used to group related data together.
- They are value types, but you can use pointers for efficient mutation.
- `p.X` is shorthand for `(*p).X`.
