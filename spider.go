package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"regexp"
	"strings"
)

type SubjectType struct {
	Name string `gorm:"column:name"`
	Type string `gorm:"column:type"`
}

func main() {

	c := colly.NewCollector(

	//colly.Async(true),
	//colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),

	)

	c.AllowedDomains = []string{"www.youzy.cn"}

	c.OnHTML("div[class=content]", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.DOM.Find("div[class=font]").Eq(0).Text())
		re, _ := regexp.Compile("（[\\S\\s]+?）")
		text = re.ReplaceAllString(text, "")
		fmt.Println(text)
		subject := e.ChildTexts("li a[href]")
		subjectType := e.ChildTexts("div[class=major-num] a")
		insertData := []SubjectType{}
		for _, name := range subjectType {
			newName := re.ReplaceAllString(name, "")
			s := SubjectType{newName, text}
			insertData = append(insertData, s)

		}
		for _, name := range subject {
			s := SubjectType{name, text}
			insertData = append(insertData, s)
		}
		fmt.Println(insertData)
		db, err := gorm.Open("mysql", "root:asdf930516@tcp(127.0.0.1:3306)/volunteer?charset=utf8&parseTime=True&loc=Local")
		if err != nil {

			panic(err)
		}
		db.SingularTable(true)
		BatchSave(db, insertData)

	})

	/*c.OnHTML("div[class=content] li a", func(e *colly.HTMLElement) {

		link := e.Attr("href")

		fmt.Printf("Link found: %q -> %s\n", e.Text, link)


		c.Visit(e.Request.AbsoluteURL(link))

	})*/

	c.OnRequest(func(r *colly.Request) {

		//fmt.Println("Visiting", r.URL.String())

	})

	//c.Limit(&colly.LimitRule{DomainGlob: "*.youzy.*", Parallelism: 5})

	// Start scraping on https://hackerspaces.org

	c.Visit("https://www.youzy.cn/tzy/search/majors/homepage")

}

func BatchSave(db *gorm.DB, emps []SubjectType) error {
	var buffer bytes.Buffer
	sql := "insert into `smart_subject_type` (`name`,`type`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range emps {
		//fmt.Println(e.Name)
		if i == len(emps)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s');", e.Name, e.Type))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s'),", e.Name, e.Type))
		}
	}
	fmt.Println(buffer.String())
	//return
	return db.Exec(buffer.String()).Error
}
