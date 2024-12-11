package main

type Message struct {
	Action    string `json:"action"`
    Name    string `json:"name"`
    Content *string `json:"content,omitempty"`
}

type DeleteFriendsMessage struct {
	Action    string `json:"action"`
    CustomerList []string `json:"customer_list"`
}

type AddGroupMessage struct {
	Action    string `json:"action"`
    Name string `json:"name"`
    StaffList []string `json:"staff_list"`
}

// 你可以在这里定义更多的数据结构