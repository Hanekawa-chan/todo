package v1

import (
	"context"
	"github.com/Hanekawa-chan/todo/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *Db
}

func NewToDoServiceServer(db *Db) v1.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

func (s *toDoServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *toDoServiceServer) Create(ctx context.Context,req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	id := s.db.save(req)

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

func (s *toDoServiceServer) Read(ctx context.Context,req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	toDo, err := s.db.findById(req.Id)
	if err != nil {
		return nil, err
	}

	td := v1.ToDo{
		Id: toDo.id,
		Title: toDo.title,
		Description: toDo.description,
		Check: toDo.check,
	}

	return &v1.ReadResponse{
		Api:  apiVersion,
		ToDo: &td,
	}, nil

}

func (s *toDoServiceServer) Update(ctx context.Context,req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	i, err := s.db.update(req)
	if err != nil {
		return &v1.UpdateResponse{
			Api:     apiVersion,
			Updated: i,
		}, err
	}

	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: i,
	}, nil
}

func (s *toDoServiceServer) Delete(ctx context.Context,req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	rows, err := s.db.delete(req)
	if err != nil {
		return &v1.DeleteResponse{
			Api:     apiVersion,
			Deleted: rows,
		}, err
	} else {
		return &v1.DeleteResponse{
			Api:     apiVersion,
			Deleted: rows,
		}, nil
	}
}

func (s *toDoServiceServer) Check(ctx context.Context,req *v1.CheckRequest) (*v1.CheckResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	rows, err := s.db.check(req)
	if err != nil {
		return &v1.CheckResponse{
			Api:     apiVersion,
			Checked: rows,
		}, err
	} else {
		return &v1.CheckResponse{
			Api:     apiVersion,
			Checked: rows,
		}, nil
	}
}

func (s *toDoServiceServer) ReadAll(ctx context.Context,req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	list, err := s.db.findAll()
	if err != nil {
		return nil, err
	} else {
		var lis []*v1.ToDo
		for _, element := range list {
			lis = append(lis, &v1.ToDo{
				Id: element.id,
				Title: element.title,
				Description: element.description,
				Check: element.check,
			})
		}
		return &v1.ReadAllResponse{
			Api:   apiVersion,
			ToDos: lis,
		}, nil
	}
}