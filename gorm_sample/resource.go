type Resource struct {
    gorm.Model

    Link        string
    Name        string
    Author      string
    Description string
    Tags        pq.StringArray `gorm:"type:varchar(64)[]"`
}