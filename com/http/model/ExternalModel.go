package model


type StudentModel struct{
	Status  int   "status"
	Message string   "message"
}


type LendAplication struct{
	Status  int   "status"
	Message string   "message"
	Data Data "data"
}


type Data struct{
	Count int
	Timestamp int64
	SystemID string
	Records []Records
}

type Records struct{
	Status  int   "status"
	Message string   "message"
	Timestamp int64
	Args Args
	System string
	Msg string
	Data []InternalData

}
type Args struct{
	Ids string
	Sn string
}

type InternalData struct{
	User_state string
	Id_no string
	Lends []Lend
}
type Lend struct{
	Main_reason_code string
	Product_type string
	Minor_reason string
	Mobile string
	Lend_id string
	Minor_reason_code string
	State string
	Create_time string
	Main_reason string
}