package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/almaraz333/finance-tracker-proto-files/expense"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewExpenseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var category string
	var amount float64

	fmt.Print("Please enter a category: ")
	fmt.Scan(&category)
	fmt.Print("Please enter an amount: ")
	fmt.Scan(&amount)

	r, err := c.CreateExpense(ctx, &pb.CreateExenseRequest{
		Category:  category,
		CreatedAt: time.Now().UTC().String(),
		Amount:    amount,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Created Expense with the amount: %v", r.GetAmount())
}
