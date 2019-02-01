package main

import "fmt"

type Methods interface {
	GetArticle() string
	BeforeSave()
	Save()
	AfterSave()
	Execute()
}

func ExecuteTemplateMethod(T Methods) {
	fmt.Println("[Article:",T.GetArticle(),"]")
	T.BeforeSave()
	T.Save()
	T.AfterSave()
	fmt.Println("[Session closed!]")
}

//Implementation

type SaveToDatabase struct {
	dbhost string
	dbuser string
	article string
}

func (s *SaveToDatabase) GetArticle() string {
	return s.article
}

func (s *SaveToDatabase) BeforeSave() {
	fmt.Println("Connecting to Database: ",s.dbuser,"@",s.dbhost,":80")
	fmt.Println("Connected to Database: Session Started!")
}

func (s *SaveToDatabase) Save() {
	fmt.Println("Saving: '",s.article,"' to database!")
}

func (s *SaveToDatabase) AfterSave() {
	fmt.Println("[Database connection is closed!]")
}

func (s *SaveToDatabase) Execute() {
	ExecuteTemplateMethod(s)
}
func main() {
	SDB := SaveToDatabase{"localhost","root","This is a sample article!"}
	SDB.Execute()
}