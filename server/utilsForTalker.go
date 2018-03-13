package server

import "log"

// Operator methods
const (
	actionGetAllRooms = "getAllRooms"
	actionEnterRoom   = "enterRoom"
	actionLeaveRoom   = "leaveRoom"
)

// CheckError checks errors and print log
func CheckError(err error, message string, fatal bool) bool {
	if err != nil {
		if fatal {
			log.Fatalln(message + ": " + err.Error())
		} else {
			log.Println(message + ": " + err.Error())
		}
	}
	return err == nil
}

//// Operator messages

type OperatorResponseAddToRoom struct {
	Room int `json:"roomID"`
}

type OperatorResponseRooms struct {
	Pair map[int]*Pair `json:"rooms"`
	Size int           `json:"size"`
}

type RequestActionWithRoom struct {
	ID int `json:"rid"`
}

type OperatorSendMessage struct {
	Message string `json:"message"`
}
