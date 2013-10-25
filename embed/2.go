package main

type Id int

func (i Id) IsNew() bool {
	return i == 0
}

type Story struct {
  Id
  ProjectId Id
  Title string
}

func (s *Story) IsNew() bool {
	return s.Id.IsNew() || s.ProjectId.IsNew()
}

func main() {
	s := &Story{Id: Id(3000)}
	println(s.Id, s.ProjectId, s.IsNew())
	s.ProjectId = Id(3000)
	println(s.Id, s.ProjectId, s.IsNew())
}