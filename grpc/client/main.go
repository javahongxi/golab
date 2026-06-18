package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"go.uber.org/zap"

	"github.com/javahongxi/golab/grpc/pb"
)

func main() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sugar.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sugar.Info("=== CreateUser ===")
	createResp, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Username: "hongxi",
		Nickname: "洪熙",
		Gender:   1,
		Age:      30,
	})
	if err != nil {
		sugar.Fatalf("CreateUser failed: %v", err)
	}
	sugar.Infof("Created: %+v", createResp.User)

	sugar.Info("=== GetUser ===")
	getResp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: createResp.User.Id})
	if err != nil {
		sugar.Fatalf("GetUser failed: %v", err)
	}
	sugar.Infof("Got: %+v", getResp.User)

	sugar.Info("=== UpdateUser ===")
	updateResp, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       createResp.User.Id,
		Nickname: "Hongxi",
		Age:      31,
	})
	if err != nil {
		sugar.Fatalf("UpdateUser failed: %v", err)
	}
	sugar.Infof("Updated: %+v", updateResp.User)

	sugar.Info("=== ListUsers ===")
	listResp, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
	if err != nil {
		sugar.Fatalf("ListUsers failed: %v", err)
	}
	sugar.Infof("Users (%d):", len(listResp.Users))
	for _, user := range listResp.Users {
		sugar.Infof("  - ID=%d, Username=%s, Nickname=%s", user.Id, user.Username, user.Nickname)
	}

	sugar.Info("=== DeleteUser ===")
	deleteResp, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{Id: createResp.User.Id})
	if err != nil {
		sugar.Fatalf("DeleteUser failed: %v", err)
	}
	sugar.Infof("Deleted: %v", deleteResp.Success)

	fmt.Println("\ngRPC client example completed!")
}
