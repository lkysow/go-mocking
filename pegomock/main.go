package pegomock

import "fmt"

// User is what we want to test.
type User struct{
	D Doer
}
// Run just calls DoIt()
func (u *User) Run() error {
	return u.D.DoIt()
}

// Doer is the interface we need to mock.
type Doer interface {
	DoIt() error
}
//go:generate pegomock generate Doer

// Client implements the Doer interface.
type Client struct{}

// DoIt just returns nil.
func (c *Client) DoIt() error {
	return nil
}

func main() {
	d := Client{}
	u := User{
		D: &d,
	}
	err := u.Run()
	fmt.Printf("err: %s\n", err)
}
