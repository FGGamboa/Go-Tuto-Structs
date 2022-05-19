package main

import (
	"fmt"

	"structs/users"
)

func main() {
	martha := users.User{Id: 1, Name: "Martha", Age: 30}
	pedro := users.User{Id: 2, Name: "Pedro", Age: 31}
	juan := users.User{Id: 3, Name: "Juan", Age: 32}
	jane := users.User{Id: 4, Name: "Jane", Age: 33}

	martha.SayHello()
	martha.AddFriend(pedro)
	martha.AddFriend(juan)
	martha.AddFriend(jane)
	martha.ListFiends()

	martha.RemoveFriend(juan.Id)
	fmt.Println("===========================")
	martha.ListFiends()
}
