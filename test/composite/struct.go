package composite

import "fmt"

type User struct {
	name   string
	gender string
}

type Group struct {
	name   string
	detail []User
}

type ApproveBodyDirect struct {
	ApplyID    string `json:"applyId"`
	ExamID     string `json:"examId"`
	ExamSource int64  `json:"examSource"`
	IsOver     int64  `json:"isOver"`
	Remark     string `json:"remark"`
	State      int64  `json:"state"`
}

func TestStruct() {
	user1 := User{"Leson", "Male"}
	user2 := User{"vicky", "Female"}
	group := Group{"family", []User{user1, user2}}
	fmt.Printf("user:%v\ngroup:%v\n", user1, group.detail[0])
	bodyDirect := ApproveBodyDirect{"applyId", "examId", 1, 1, "Approved", 0}
	fmt.Printf("%v\n", bodyDirect.IsOver)
}
