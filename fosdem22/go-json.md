# JSON Serialization - The Fine Print

A quiz! Consider the following code:

	a := 1
	data, _ := json.Marshal(a)
	var b interface{}
	json.Unmarshal(data, &b)
	fmt.Println(a == b)

What will be printed? And why it's "1 1 false"? :)

We all work with JSON, and even though it's a simple format, there are several sharp edges you should should avoid.
In this talk, we'll journey from the common well lit areas of JSON serialization and end in some darker corners. I hope you'll enjoy the ride.

---
- types
- streaming
- missing values
- `json:""` language
    - "-,"
    - ",string"
    - "omitempty"
- time format
- binary
- mapstructure
- split encoding types from internal
- validate
- `json.NewEncoder(os.Stdout).Encode("<script>")` (same with Unmarshal)
- MarshalText
- https://pkg.go.dev/encoding/json#Marshal
