
package main 



func main() {
	store, err := NewPostgresStore()
	server := NewAPIServer(":3000")
	server.Run()
}