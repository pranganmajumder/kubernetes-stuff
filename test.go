type Stu struct {
	Name string
	Age int
	Id int
}

func main()  {
	s := &Stu{
		Name: "Leo",
		Age: 21,
		Id: 1,
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, s)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%q\n", buf)
}