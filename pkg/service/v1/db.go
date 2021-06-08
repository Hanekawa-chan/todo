package v1

import (
	v1 "github.com/Hanekawa-chan/todo/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type toDo struct {
	id          int64
	title       string
	description string
	check       bool
}
type Db struct {
	toDos []toDo
}

func (db *Db) findById(id int64) (toDo, error) {
	var item toDo
	for _, element := range db.toDos {
		if element.id == id {
			item = element
		}
	}
	if item.id == 0 {
		return toDo{}, status.Errorf(codes.NotFound, "can't find todo with this id")
	} else {
		return item, nil
	}
}

func (db *Db) findAll() ([]toDo, error) {
	var item []toDo
	for _, element := range db.toDos {
		item = append(item, element)
	}
	if len(item) == 0 {
		return []toDo{}, status.Errorf(codes.NotFound, "can't find todos")
	} else {
		return item, nil
	}
}

func removeById(s []toDo, i int64) []toDo {
	var id int
	for index, element := range s {
		if element.id == i {
			id = index
		}
	}
	s[id] = s[len(s)-1]
	return s[:len(s)-1]
}

func (db *Db) save(r *v1.CreateRequest) int64 {
	var id int64
	if len(db.toDos) > 0 {
		id = db.toDos[0].id
		for _, element := range db.toDos {
			if element.id > id {
				id = element.id
			}
		}
		id += 1
	} else {
		id = 1
	}
	db.toDos = append(db.toDos, toDo{id, r.ToDo.Title, r.ToDo.Description, r.ToDo.Check})
	return id
}

func (db *Db) count(id int64) int64 {
	var i int64 = 0
	for _, element := range db.toDos {
		if element.id == id {
			i += 1
		}
	}
	return i
}

func (db *Db) update(r *v1.UpdateRequest) (int64, error) {
	i := db.count(r.ToDo.Id)
	if i == 0 {
		return 0, status.Errorf(codes.Unknown, "updated nothing")
	} else {
		db.toDos = append(db.toDos, toDo{
			r.ToDo.Id,
			r.ToDo.Title,
			r.ToDo.Description,
			r.ToDo.Check,
		})
		return i, nil
	}
}

func (db *Db) check(r *v1.CheckRequest) (int64, error) {
	i := db.count(r.Id)
	var id int64 = -1
	for index, element := range db.toDos {
		if element.id == r.Id {
			id = int64(index)
		}
	}
	if id == -1 {
		return 0, status.Errorf(codes.Unknown, "updated nothing")
	} else {
		t := db.toDos[id]
		t.check = true
		db.toDos[id] = t
		return i, nil
	}
}

func (db *Db) delete(r *v1.DeleteRequest) (int64, error) {
	i := db.count(r.Id)
	if i == 0 {
		return 0, status.Errorf(codes.Unknown, "deleted nothing")
	} else {
		db.toDos = removeById(db.toDos, r.Id)
		return i, nil
	}
}
