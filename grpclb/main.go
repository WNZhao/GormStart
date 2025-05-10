/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-05-10 13:47:24
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-05-10 13:59:13
 * @FilePath: /GormStart/grpclb/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"GormStart/grpclb/proto"
	"context"
	"log"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"consul://192.168.1.7:8500/user-srv?wait=14s&tag=joyshop",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userSrvClient := proto.NewUserClient(conn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Page:     1,
		PageSize: 2,
	})
	if err != nil {
		log.Fatal(err)
	}
	for index, data := range rsp.Data {
		log.Printf("第%d条数据: %v", index, data)
	}
}
