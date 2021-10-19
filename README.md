# conhash
A simple implementation of Consistent Hashing in Go.

## Usage
**Creating a New Ring**  
```go
nodes := []conhash.Node{
    {Host: "localhost", Port: 8000},
    {Host: "localhost", Port: 8001},
    {Host: "localhost", Port: 8002},
}

ring := conhash.New(nodes)
```
  
**Find a Target Node**  
```go
target := ring.Find("your-key")
```
  
**Add a New Node**  
```
node := conhash.Node{Host: "localhost", Port: 8003}
err := ring.Add(node)
if err != nil { 
    // ...
}
```
  
**Remove a Node**  
```go
node := conhash.Node{Host: "localhost", Port: 8003}
err := ring.Remove(node)
if err != nil { 
    // ...
}
```
## Contributing
All contributions are welcome, just open a pull request.
