package model

type Roommate struct {
    Name        string
    Email       string
    PhoneNumber string
}

func NewRoommate(name, email, phoneNumber string) *Roommate{
    return &Roommate {
        Name: name,
        Email: email,
        PhoneNumber: phoneNumber,
    }
}



