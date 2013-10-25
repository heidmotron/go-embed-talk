package main

type Id int

func (i Id) IsNew() bool {
	return i == 0
}

type Story struct {
  Id // embedding
  Title string
  State string
}


func main() {
	s := Story{Title: "Andre 3000", State: "Hey Ya"}
	println(s.Id, s.IsNew())
	s.Id = Id(3000)
	println(s.Id, s.IsNew())
}

