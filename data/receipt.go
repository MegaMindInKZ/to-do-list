package data

type Receipt struct {
	ID          int
	User_ID     int
	Name        string
	Photo       string
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
	st, err := DB.Prepare("INSERT INTO RECEIPTS(USER_ID, NAME, PHOTO, INSTRUCTION) VALUES ($1, $2, $3, $4) RETURNING ID, CREATED_AT")
	if err != nil {
		//danger method
		return
	}
	defer st.Close()
	err = st.QueryRow(receipt.User_ID, receipt.Name, receipt.Photo, receipt.Instruction).Scan(
		&receipt.ID, &receipt.CreatedAt,
	)
	return
}

func AllReceipts() (receipts []Receipt, err error) {
	rows, err := DB.Query("SELECT ID, USER_ID, NAME, PHOTO, INSTRUCTION FROM RECEIPTS")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var receipt Receipt
		err = rows.Scan(&receipt.ID, &receipt.User_ID, &receipt.Name, &receipt.Photo, &receipt.Instruction)
		if err != nil {
			return
		}
		receipts = append(receipts, receipt)
	}
	return
}
