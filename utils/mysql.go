package utils

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Age  int8
}

func init() {
	//mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beego?charset=utf8", maxIdle, maxConn)
	orm.RegisterModel(new(User))

	// RegisterDataBase 里面如果写过了，这里就不需要在设置了。
	orm.SetMaxIdleConns("default", maxIdle)
	orm.SetMaxOpenConns("default", maxConn)
	//orm.RunSyncdb("default", false, true)
}

// 支持 ORM 操作
func doOrm() {
	o := orm.NewOrm()
	user := User{Name: "wanyang", Age: 28}

	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	user.Name = "wanyang3"
	num, err := o.Update(&user, "Name")
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

// 支持执行原生 sql 语句
func doSql() {
	o := orm.NewOrm()
	// Exec 返回 sql.Result 对象
	rs, err := o.Raw("update user set age = ? where id = ?", "30", 5).Exec()
	if err == nil {
		rowsAffected,_ := rs.RowsAffected()
		lastInsertId,_ := rs.LastInsertId()
		fmt.Printf("rs.RowsAffected: %v, rs.LastInsertId: %v\n", rowsAffected, lastInsertId )
	}
	fmt.Println("0: --------------------------------------------")

	// Values 返回结果集的 key => value 值
	var value []orm.Params
	num, err := o.Raw("select * from user").Values(&value)
	if err == nil && num > 0 {
		for _, term := range value {
			fmt.Printf("Id: %s, Name: %s, Age: %s\n", term["id"], term["name"], term["age"])
		}
	}
	fmt.Println("1: --------------------------------------------")

	// ValuesList 返回结果集 slice
	var lists []orm.ParamsList
	num, err = o.Raw("select * from user where id > ?", 3).ValuesList(&lists)
	if err == nil && num > 0 {
		for _,list := range lists {
			fmt.Printf("User: %v\n", list)
		}
	}
	fmt.Println("2: --------------------------------------------")

	// ValuesFlat 返回单一字段的平铺 slice 数据
	var list orm.ParamsList
	num, err = o.Raw("select name from user where id > ?", 2).ValuesFlat(&list)
	if err == nil && num > 0 {
		for _,val := range list {
			fmt.Printf("User.Name: %s\n", val)
		}
	}
	fmt.Println("3: --------------------------------------------")

	// RowsToMap 查询结果匹配到 map 里
	res := make(orm.Params)
	num, err = o.Raw("select id, name from user").RowsToMap(&res, "id", "name")
	if err == nil && num > 0 {
		for k, v := range res {
			fmt.Printf("id: %s, name: %s\n", k, v)
		}
	}
	fmt.Println("4: --------------------------------------------")

	// RowsToStruct 查询结果匹配到 struct 里
	str := new(User)
	num, err = o.Raw("select id, name from user").RowsToStruct(&res, "id", "name")
	if err == nil && num > 0 {
		fmt.Printf("Id: %v\n", str.Id)
		fmt.Printf("Name: %v\n", str.Name)
	}
	fmt.Println("5: --------------------------------------------")


	// QueryRow 提供高级 sql mapper 功能，支持 struct
	var user User
	//err := o.Raw("select * from user where id = ?", 5).QueryRow(&user.Id, &user.Name, &user.Age)
	err = o.Raw("select * from user where id = ?", 5).QueryRow(&user)
	if err == nil {
		fmt.Printf("Id: %d, Name: %s, Age: %d\n", user.Id, user.Name, user.Age)
	}
	fmt.Println("6: --------------------------------------------")

	// QueryRows 提供高级 sql mapper 功能，支持 struct、map 但都是 slice 类型。
	var id []int
	var name []string
	var age []int8
	num, err = o.Raw("select * from user where id >= 3").QueryRows(&id, &name, &age)
	fmt.Printf("Total row: %d\n", num)
	if err == nil {
		fmt.Printf("Id: %v, Name: %v, Age: %v\n", id, name, age)
		for index, value := range id {
			fmt.Printf("Row: %d, Id: %d, Name: %s, Age: %d\n", index+1, value, name[index], age[index])
		}
	}

	var users []User
	num1, err := o.Raw("select * from user where id >= 3").QueryRows(&users)
	fmt.Printf("Total row: %d\n", num1)
	if err == nil {
		fmt.Printf("User: %v\n", users)
		for index, u := range users {
			fmt.Printf("Row: %d, Id: %d, Name: %s, Age: %d\n", index+1, u.Id, u.Name, u.Age)
		}
	}
	fmt.Println("7: ---------------------------------------------")

	// Prepare 一次 prepare 多次 exec，提高批量执行的速度
	rp, err := o.Raw("update user set name= ? where id = ?").Prepare()
	rs, err = rp.Exec("wanyang33", 4)
	rs, err = rp.Exec("Beyta", 5)
	rp.Close()
}

// 支持事务处理
func doTransaction() {
	o := orm.NewOrm()
	o.Begin()
	user := User{Id: 5, Name: "Lucy", Age: 19}
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
}

// 支持 QueryBuilder，功能类似 ORM，但 ORM 更适用于简单的 CRUD 操作，而 QueryBuilder 则更适用于复杂的查询
func doQueryBuilder() {
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("id,name,age").
		From("user").
		Where("age > ?").
		And("id < ?").
		OrderBy("age").
		Asc().
		Limit(2).
		Offset(0)
	sql := qb.String()
	fmt.Printf("sql: %s\n", sql)

	var users []User
	o := orm.NewOrm()
	num, err := o.Raw(sql, 20, 100).QueryRows(&users)
	if err == nil && num > 0 {
		for _, user := range users {
			fmt.Printf("Id: %d, Name: %s, Age: %d\n", user.Id, user.Name, user.Age)
		}
	}
}

func main() {
	//dorm()
	orm.Debug = true
	//doSql()
	//doTransaction()
	doQueryBuilder()
}
