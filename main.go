package main

import (
	"MVC_member/Model"
	"fmt"
)

func main() {
	Model.InsertMember(1, "Mikhael", "Surabaya", 90)
	Model.InsertMember(2, "Robby", "Surabaya", 89)
	Model.InsertMember(3, "Ivan", "Surabaya", 85)

	Model.SearchMember(3)

	SearchValue := Model.CariMember(2)
	fmt.Println("Func CariMember :", SearchValue.Data)

	Model.UpdateMember(3, "Rifan", "Surabaya")
	Model.UpdateMember(4, "Ikhsan", "Surabaya")
	Model.DeleteMember(2)

	members := Model.ReadAllMember()
	fmt.Println("--------------------------")
	if members != nil {
		for members.Next != nil {
			fmt.Println(members.Next.Data)
			members = members.Next
		}
	}
	fmt.Println("--------------------------")
}
