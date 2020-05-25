package prototype

type Customer struct {
	Name  string
	Email string
}

func NewCustomer(name, email string) Customer {
	return Customer{
		Name:  name,
		Email: email,
	}
}

func (c Customer) WithName(name string) Customer {
	c.Name = name
	return c
}

func (c Customer) WithEmail(email string) Customer {
	c.Email = email
	return c
}
