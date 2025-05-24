# 📄 Go Code Explanation – Using Interfaces with Structs

## ✅ Objective

This Go program demonstrates:

- How to use **interfaces** to define behavior.
- How **structs** can implement that behavior.
- How to **call methods** polymorphically through the interface.

---

## 🔧 Code Summary

```go
package main
```

### `package main`

- Tells the Go compiler this is an **executable** program.
- Execution begins at the `main()` function.

---

```go
import "fmt"
```

### `import "fmt"`

- Imports the `fmt` package (standard Go library).
- Used for formatted output, such as `fmt.Println()`.

---

```go
type shape interface {
	area() float32
	perimeter() float32
}
```

### `shape` Interface

Defines a contract that any type (like `rect`) must follow if it wants to be considered a `shape`.

- `area()` and `perimeter()` are method **signatures**.
- Any struct that **implements both** these methods satisfies the `shape` interface.

This enables **polymorphism** — treating different shapes (circle, triangle, etc.) the same way if they implement the same methods.

---

```go
type rect struct {
	width  float32
	height float32
}
```

### `rect` Struct

Defines a **rectangle** shape with:

- `width`: horizontal size
- `height`: vertical size

The `rect` struct will now be used to implement the `shape` interface.

---

```go
func (r rect) area() float32 {
	return r.width * r.height
}
```

### `area()` Method for `rect`

- Defined on the `rect` type using `(r rect)` — `r` is the receiver (like `this` in other languages).
- Implements the `area()` method required by the `shape` interface.
- Calculates area using `width × height`.

---

```go
func (r rect) perimeter() float32 {
	return 2 * (r.width + r.height)
}
```

### `perimeter()` Method for `rect`

- Also defined on the `rect` type.
- Returns the perimeter: `2 × (width + height)`.

Now `rect` implements **both methods**, so it satisfies the `shape` interface.

---

```go
func main() {
	var s shape
	s = rect{width: 20, height: 10}

	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}
```

### `main()` Function – Program Entry Point

#### Step-by-step:

1. `var s shape`

   - Declares `s` as a variable of type `shape` (interface type).
   - `s` can hold any type that implements the `shape` interface.

2. `s = rect{width: 20, height: 10}`

   - Creates a `rect` instance.
   - Since `rect` implements both `area()` and `perimeter()`, it can be assigned to a `shape` interface variable.

3. `fmt.Println("Area:", s.area())`

   - Calls the `area()` method on the `shape` variable.
   - Internally, Go calls `rect.area()`.

4. `fmt.Println("Perimeter:", s.perimeter())`
   - Similarly, calls the `perimeter()` method.

---

## 🧠 What You Learn from This Example

| Concept         | What It Teaches                                                |
| --------------- | -------------------------------------------------------------- |
| Interface       | Defines behavior, not data                                     |
| Struct          | Holds data and implements methods                              |
| Method receiver | Connects functions to specific types                           |
| Polymorphism    | You can use different types the same way via interfaces        |
| Encapsulation   | Methods hide the logic for area/perimeter behind the interface |

---

## 📌 Output of the Program

```
Area: 200
Perimeter: 60
```

Because:

- Area = `20 × 10 = 200`
- Perimeter = `2 × (20 + 10) = 60`

---

## ✅ In Summary

This is a foundational example of how Go uses **interfaces + structs** to model behavior. It prepares you for:

- More complex polymorphic code.
- Working with slices of shapes (multiple types).
- Using pointers when modifying struct fields.
