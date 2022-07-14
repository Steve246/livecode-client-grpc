package main

import "enigmacamp.com/lopei_grpc_cnlt/delivery"

func main() {
	delivery.Cli().Run()
	//var conn *grpc.ClientConn
	//conn, err := grpc.Dial("localhost:9898", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("did not connect: %s", err)
	//}
	//defer conn.Close()
	//
	//c := service.NewLopeiPaymentClient(conn)
	//
	//response, err := c.CheckBalance(context.Background(), &service.CheckBalanceMessage{
	//	LopeiId: int32(1),
	//})
	//if err != nil {
	//	log.Fatalf("Error when calling Check Balance: %s", err)
	//}
	//log.Printf("%v", response)
	//
	//response, err = c.DoPayment(context.Background(), &service.PaymentMessage{
	//	LopeiId: int32(1),
	//	Amount:  30000,
	//})
	//if err != nil {
	//	log.Fatalf("Error when calling Check Balance: %s", err)
	//}
	//log.Printf("%v", response)
}
