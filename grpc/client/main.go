package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/javahongxi/golab/grpc/pb"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("=== CreateUser ===")
	createResp, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Username: "hongxi",
		Nickname: "洪熙",
		Gender:   1,
		Age:      30,
	})
	if err != nil {
		log.Fatalf("CreateUser failed: %v", err)
	}
	log.Printf("Created: %+v", createResp.User)

	log.Println("\n=== GetUser ===")
	getResp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: createResp.User.Id})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}
	log.Printf("Got: %+v", getResp.User)

	log.Println("\n=== UpdateUser ===")
	updateResp, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       createResp.User.Id,
		Nickname: "Hongxi",
		Age:      31,
	})
	if err != nil {
		log.Fatalf("UpdateUser failed: %v", err)
	}
	log.Printf("Updated: %+v", updateResp.User)

	log.Println("\n=== ListUsers ===")
	listResp, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
	if err != nil {
		log.Fatalf("ListUsers failed: %v", err)
	}
	log.Printf("Users (%d):", len(listResp.Users))
	for _, user := range listResp.Users {
		log.Printf("  - ID=%d, Username=%s, Nickname=%s", user.Id, user.Username, user.Nickname)
	}

	log.Println("\n=== DeleteUser ===")
	deleteResp, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{Id: createResp.User.Id})
	if err != nil {
		log.Fatalf("DeleteUser failed: %v", err)
	}
	log.Printf("Deleted: %v", deleteResp.Success)

	fmt.Println("\ngRPC client example completed!")
}
