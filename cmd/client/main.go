package main

import (
	"context"
	"flag"
	v1 "github.com/Hanekawa-chan/todo/pkg/api/v1"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func main() {
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var request string
	var title string
	var description string
	var i int64
	if len(os.Args) > 2 {
		request = os.Args[2]
		switch request {
		case "create":
			{
				title = os.Args[3]
				description = os.Args[4]
				req1 := v1.CreateRequest{
					Api: apiVersion,
					ToDo: &v1.ToDo{
						Title:       title,
						Description: description,
					},
				}
				res1, err := c.Create(ctx, &req1)
				if err != nil {
					log.Fatalf("Create failed: %v", err)
				}
				log.Printf("Create result: <%+v>\n\n", res1)
			}
		case "read":
			{
				i, err = strconv.ParseInt(os.Args[3], 10, 64)
				if err != nil {
					panic(err)
				}
				req2 := v1.ReadRequest{
					Api: apiVersion,
					Id:  i,
				}
				res2, err := c.Read(ctx, &req2)
				if err != nil {
					log.Fatalf("Read failed: %v", err)
				}
				log.Printf("Read result: <%+v>\n\n", res2)
			}
		case "delete":
			{
				i, err = strconv.ParseInt(os.Args[3], 10, 64)
				if err != nil {
					panic(err)
				}
				req5 := v1.DeleteRequest{
					Api: apiVersion,
					Id:  i,
				}
				res5, err := c.Delete(ctx, &req5)
				if err != nil {
					log.Fatalf("Delete failed: %v", err)
				}
				log.Printf("Delete result: <%+v>\n\n", res5)

			}
		case "update":
			{
				i, err = strconv.ParseInt(os.Args[3], 10, 64)
				if err != nil {
					panic(err)
				}
				title = os.Args[4]
				description = os.Args[5]
				req3 := v1.UpdateRequest{
					Api: apiVersion,
					ToDo: &v1.ToDo{
						Id:          i,
						Title:       title,
						Description: description + " + updated",
					},
				}
				res3, err := c.Update(ctx, &req3)
				if err != nil {
					log.Fatalf("Update failed: %v", err)
				}
				log.Printf("Update result: <%+v>\n\n", res3)
			}
		case "readAll":
			{
				req4 := v1.ReadAllRequest{
					Api: apiVersion,
				}
				res4, err := c.ReadAll(ctx, &req4)
				if err != nil {
					log.Fatalf("ReadAll failed: %v", err)
				}
				log.Printf("ReadAll result: <%+v>\n\n", res4)
			}
		case "check":
			{
				i, err = strconv.ParseInt(os.Args[3], 10, 64)
				if err != nil {
					panic(err)
				}
				req6 := v1.CheckRequest{
					Api: apiVersion,
					Id:  i,
				}
				res6, err := c.Check(ctx, &req6)
				if err != nil {
					log.Fatalf("Update failed: %v", err)
				}
				log.Printf("Update result: <%+v>\n\n", res6)
			}
		default:
			log.Printf("please specify something" + os.Args[1] + " " + os.Args[2] + " " + os.Args[3])
		}
	} else {
		log.Printf("please specify type of request: they can be 'create', 'read', 'update', 'delete', 'readAll', 'check'")
	}

}
