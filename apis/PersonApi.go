package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	."test/models"
)

func IndexApi(c *gin.Context)  {
	c.String(http.StatusOK, "welcome!")
}

func AddPersonApi(c *gin.Context)  {
	name := c.Request.FormValue("name")
	age := c.Request.FormValue("age")
	ages,err := strconv.Atoi(age)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Name:name, Age:ages}

	id, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg":msg,
	})
}

func GetPersonApi(c *gin.Context)  {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id:id}
	person, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

func GetPersonsApi(c *gin.Context)  {
	var p Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"persons":persons,
	})
}

func ModPersonApi(c *gin.Context)  {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil{
		log.Fatalln(err)
	}
	p := Person{Id:id}
	err = c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := p.ModPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("update person %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg":msg,
	})
}

func DelPersonApi(c *gin.Context)  {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id:id}
	ra ,err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("DELETE person %d successful %d", p.Id, ra)

	c.JSON(http.StatusOK, gin.H{
		"msg":msg,
	})
}
