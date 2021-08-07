package student

import (
	"awesomeProject/practice/mockery/pkg/model"
	"awesomeProject/practice/mockery/pkg/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Student struct {
}

func NewStudentService() *Student {
	return &Student{}
}

func (p *Student) ListStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	res, err := repository.Repo.ListStudent(ctx)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (p *Student) GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	res, err := repository.Repo.GetStudent(ctx, id)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (p *Student) CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Student model.StudentDetails
	err := Student.DecodeFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Failed to Decode", http.StatusBadRequest)
		return
	}
	if err = Student.Validate(); err != nil {
		http.Error(w, "failed to validate struct", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	res, err := repository.Repo.CreateStudent(ctx, &Student)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	fmt.Println(res)

	w.WriteHeader(http.StatusCreated)
	Student.EncodeToJSON(w)
	return

}

func (p *Student) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	err := repository.Repo.DeleteStudent(ctx, id)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	res := model.StudentDetails{

		Msg: "Student details deleted",
	}
	json.NewEncoder(w).Encode(res.Msg)
}

func (p *Student) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var Student model.StudentDetails
	w.Header().Set("Content-Type", "application/json")
	err := Student.DecodeFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Failed to Decode", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	res, err := repository.Repo.GetStudent(ctx, id)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	Student.Id = res.Id
	res, err = repository.Repo.UpdateStudent(ctx, &Student)
	if err != nil {
		http.Error(w, "Failed to create in database", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(res)
}
