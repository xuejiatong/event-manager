package main

type Referrer struct {
	ID       int    `json:"id"`
	Referrer string `json:"referrer"`
}

type Support struct {
	ID      int    `json:"id"`
	Support string `json:"support"`
}

func getReferrers() ([]Referrer, error) {
	var referrers []Referrer
	query := `SELECT id, referrer FROM referrers`
	rows, err := db.Query(query)
	if err != nil {
		return referrers, err
	}
	defer rows.Close()
	for rows.Next() {
		var r Referrer
		err := rows.Scan(&r.ID, &r.Referrer)
		if err != nil {
			return referrers, err
		}
		referrers = append(referrers, r)
	}

	return referrers, err
}

func addReferrer(referrer string) (int, error) {
	insertStatement := `
		INSERT INTO referrers (referrer)
		VALUES ($1)
		RETURNING id`
	var id int
	err := db.QueryRow(insertStatement, referrer).Scan(&id)
	return id, err
}

func getSupport() ([]Support, error) {
	var support []Support
	query := `SELECT id, support FROM supports`
	rows, err := db.Query(query)
	if err != nil {
		return support, err
	}
	defer rows.Close()
	for rows.Next() {
		var s Support
		err := rows.Scan(&s.ID, &s.Support)
		if err != nil {
			return support, err
		}
		support = append(support, s)
	}

	return support, err
}

func addSupport(support string) (int, error) {
	insertStatement := `
		INSERT INTO supports (support)
		VALUES ($1)
		RETURNING id`
	var id int
	err := db.QueryRow(insertStatement, support).Scan(&id)

	return id, err
}
