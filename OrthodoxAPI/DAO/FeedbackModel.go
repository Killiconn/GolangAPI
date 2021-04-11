package DAO


type Feedback struct {
	Username  string `gorm:"column:Username"`
	Feedback  string `gorm:"column:Feedback"`
}

// This function is for GORM library which takes in the name of the database

func (Feedback) TableName() string {
	return "feedbacks"
   }