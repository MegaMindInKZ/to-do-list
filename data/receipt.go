package data

type Receipt struct {
	ID          int
	User_ID     int
	Name        string
	Photo       string
	Duration    int
	Instruction string
	CreatedAt   string
}

type Ingredient struct {
	ID         int
	Name       string
	Receipt_ID int
	Amount     int
	Unit       string
}

func (receipt *Receipt) Create() (err error) {
	st, err := DB.Prepare("INSERT INTO RECEIPTS(USER_ID, NAME, PHOTO, DURATION, INSTRUCTION) VALUES ($1, $2, $3, $4, $5) RETURNING ID, CREATED_AT")
	if err != nil {
		//danger method
		return
	}
	defer st.Close()
	err = st.QueryRow(receipt.User_ID, receipt.Name, receipt.Photo, receipt.Duration, receipt.Instruction).Scan(
		&receipt.ID, &receipt.CreatedAt,
	)
	return
}

func (receipt *Receipt) Delete() (err error) {
	st, err := DB.Prepare("DELETE FROM RECEIPT WHERE ID=$1")
	if err != nil {
		return
	}
	defer st.Close()
	_, err = st.Exec(receipt.ID)
	return
}

func ReceiptsByUserID(userID int) (receipts []Receipt, err error) {
	rows, err := DB.Query(
		"SELECT ID, USER_ID, NAME, PHOTO, DURATION, INSTRUCTION FROM RECEIPTS WHERE USER_ID=$1", userID,
	)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var receipt Receipt
		err = rows.Scan(
			&receipt.ID, &receipt.User_ID, &receipt.Name, &receipt.Photo, &receipt.Duration, &receipt.Instruction,
		)
		if err != nil {
			return
		}
		receipts = append(receipts, receipt)
	}
	return
}

func AllReceipts() (receipts []Receipt, err error) {
	rows, err := DB.Query("SELECT ID, USER_ID, NAME, PHOTO, DURATION, INSTRUCTION FROM RECEIPTS")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var receipt Receipt
		err = rows.Scan(
			&receipt.ID, &receipt.User_ID, &receipt.Name, &receipt.Photo, &receipt.Duration, &receipt.Instruction,
		)
		if err != nil {
			return
		}
		receipts = append(receipts, receipt)
	}
	return
}
