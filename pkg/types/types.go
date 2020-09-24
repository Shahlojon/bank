package types

//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д)
type Money int64

//Currency представляет номер карты
type Currency string

//Status представляет собой статус платежа
type Status string

//Предопределенные статусы платежей
const (
	StatusOk Status ="OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

//Category представляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д).
type Category string


//Payment представляет информацию о платеже
type Payment struct{
	ID int
	Amount Money
	Category Category
	Status Status
}
