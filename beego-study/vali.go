package main

type User struct {
	Name string
	Age int
}

func main() {
	u := User{"manasdmjkexa", 14}
	valid := validation.Validation{}
	valid.Required(u.Name, "name")
	valid.MaxSize(u.Name, 15, "nameMax")
	valid.Range(u.Age, 15, 18, "age").Message("%d禁", u.Age)

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	// or use like this
	if v := valid.Max(u.Age, 140, "age"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
	}
	// 定制错误信息
	minAge := 14
	valid.Min(u.Age, minAge, "age").Message("少儿不宜！")
	// 错误信息格式化
	valid.Min(u.Age, minAge, "age").Message("%d不禁", minAge)
}
