package mod01_Simple_factory

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

// NewAPI return Api instance by the type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi,%s", name)
}

// helloAPI is one of API implement
type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello,%s", name)
}
