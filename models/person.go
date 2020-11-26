package models
//
//import (
//	db "test/database"
//)
//
//type Person struct {
//	Id   int    `json:"id" form:"id"`
//	Name string `json:"name" form:"name"`
//	Age  int    `json:"age" form:"age"`
//}
//
//func (p *Person) AddPerson () (id int64, err error)  {
//	rs, err := db.SqlDB.Exec("INSERT INTO test(name, age) VALUE (?, ?)", p.Name, p.Age)
//	if err != nil {
//		return
//	}
//
//	id, err = rs.LastInsertId()
//	return
//}
//
//func (p *Person) GetPerson () (person Person, err error) {
//	err = db.SqlDB.QueryRow("SELECT id,name,age FROM test WHERE id=?", p.Id).Scan(&person.Id, &person.Name, &person.Age)
//
//	if err != nil {
//		return
//	}
//	return
//}
//
//func (p *Person) GetPersons () (persons []Person, err error)  {
//	persons = make([]Person, 0)
//	rows, err := db.SqlDB.Query("SELECT id,name,age FROM test")
//	if err != nil {
//		return
//	}
//	defer rows.Close()
//	for rows.Next(){
//		var person Person
//		rows.Scan(&person.Id, &person.Name, &person.Age)
//		persons = append(persons, person)
//	}
//	if err = rows.Err(); err != nil {
//		return
//	}
//	return
//}
//
//func (p *Person) ModPerson () (ra int64, err error)  {
//	stmt, err := db.SqlDB.Prepare("UPDATE test SET name=?,age=? WHERE id=?")
//	defer stmt.Close()
//	if err != nil {
//		return
//	}
//	rs, err := stmt.Exec(p.Name, p.Age, p.Id)
//	if err != nil {
//		return
//	}
//	ra, err = rs.RowsAffected()
//	return
//}
//
//func (p *Person) DelPerson () (ra int64, err error)  {
//	stmt, err := db.SqlDB.Prepare("DELETE FROM  test WHERE id=?")
//	defer stmt.Close()
//	if err != nil {
//		return
//	}
//	rs, err := stmt.Exec(p.Id)
//	if err != nil {
//		return
//	}
//	ra, err = rs.RowsAffected()
//	return
//}