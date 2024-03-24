package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/almaraz333/finance-tracker-proto-files/expense"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExpenseServer
	db    *sql.DB
	mutex sync.Mutex
}

func (s *server) GetExpenses(ctx context.Context, _ *pb.Empty) (*pb.GetExpensesResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var res = []*pb.ExpenseObject{}

	rows, err := s.db.Query("select * from expenses")
	defer rows.Close()

	for rows.Next() {
		var category string
		var amount float64
		var createdAt string
		var id int32

		err = rows.Scan(&id, &createdAt, &category, &amount)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, &pb.ExpenseObject{
			Category:  category,
			CreatedAt: createdAt,
			Amount:    amount,
			Id:        id,
		})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &pb.GetExpensesResponse{
		Expenses: res,
	}, nil
}

func (s *server) CreateExpense(ctx context.Context, in *pb.CreateExenseRequest) (*pb.CreateExpenseResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("Received: %v, %v", in.GetAmount(), in.GetCategory())

	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into expenses(category, amount) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	stmt.Exec(in.GetCategory(), in.GetAmount())

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	return &pb.CreateExpenseResponse{
			Amount: in.GetAmount(),
		},
		nil
}

func main() {
	db, err := sql.Open("sqlite3", "expense.db")

	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	// Create Table if data has been/needs to be deleted

	// sqlStmt := `
	// create table expenses (
	//  id integer not null primary key autoincrement,
	//  createdAt datetime default current_timestamp,
	//  category text not null,
	//  amount real not null
	//  );
	// `
	// _, err = db.Exec(sqlStmt)
	//
	// if err != nil {
	// 	log.Printf("%q: %s\n", err, sqlStmt)
	// 	return
	// }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", 50051))

	if err != nil {
		log.Fatalf("Could not lister on port: %v, with error: %v", 50051, err)
	}

	s := grpc.NewServer()

	pb.RegisterExpenseServer(s, &server{
		db: db,
	})

	log.Printf("Server listening on port %v ...", 50051)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
