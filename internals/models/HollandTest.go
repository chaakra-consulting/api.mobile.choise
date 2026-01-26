package models

type HollandTest struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	QuestionText string `json:"question_text"`
	Type         string `json:"type"`
	Identifier   string `json:"identifier"`
}
