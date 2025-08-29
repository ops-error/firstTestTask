package models

type Item struct {
	Chrt_id     int64   `db:"chrt_id" json:"chrt_id"`
	Price       float64 `db:"price" json:"price"`
	Rid         string  `db:"rid" json:"rid"`
	Name        string  `db:"name" json:"name"`
	Sale        int64   `db:"sale" json:"sale"`
	Size        string  `db:"size" json:"size"`
	Total_price float64 `db:"total_price" json:"total_price"`
	Nm_id       int64   `db:"nm_id" json:"nm_id"`
	Brand       string  `db:"brand" json:"brand"`
	Status      int64   `db:"status" json:"status"`
}
