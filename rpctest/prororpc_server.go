package main

import (
	"errors"
	"github.com/lurenjia528/study-go/rpctest/pb"
)

// 算术运算结构体
type Arith struct {
}

// 乘法运算方法
func (this *Arith) Multiply(req *pb.ArithRequest, res *pb.ArithResponse) error {
	res.Pro = req.GetA() * req.GetB()
	return nil
}

// 除法运算方法
func (this *Arith) Divide(req *pb.ArithRequest, res *pb.ArithResponse) error {
	if req.GetB() == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.GetA() / req.GetB()
	res.Rem = req.GetA() % req.GetB()
	return nil
}

func main() {
	pb.ListenAndServeArithService("tcp","127.0.0.1:8097",new(Arith))
}
