package main

import (
	"errors"
	"fmt"

	"github.com/zhaoqiang0201/design-pattern/chain/pkg"
)

//  挂号--》诊室看病--》收费处缴费--》药房拿药方
//  挂号--》--》初诊--》影响科拍片--》诊室看病--》收费处缴费--》药房拿药方

// 病人、患者
type PatientHandler interface {
	Execute(*patient) error
	SetNext(PatientHandler) PatientHandler
	Do(*patient) error
}

// 充当抽象类型，实现公共方法，抽象方法不实现留给实现类自己实现
type Next struct {
	nextHandelr PatientHandler
}

func (n *Next) SetNext(handler PatientHandler) PatientHandler {
	n.nextHandelr = handler
	return handler
}

func (n *Next) Execute(patient *patient) error {
	// 调用不到外部类型的 Do 方法，所以 Next 不能实现 Do 方法
	if n.nextHandelr != nil {
		if err := n.nextHandelr.Do(patient); err != nil {
			return err
		}

		return n.nextHandelr.Execute(patient)
	}
	return errors.New("nextHandler is null")
}

//流程中的请求类--患者
type patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

// Reception 挂号处处理器
type ReceptionAction struct {
	Next
}

func (r *ReceptionAction) Do(p *patient) error {
	if p.RegistrationDone {
		fmt.Println("挂号处处理器ReceptionAction already done")
		return nil
	}
	fmt.Println("挂号处处理器doing")
	p.RegistrationDone = true
	return nil
}

// Clinic 诊室处理器--用于医生给病人看病
type Clinic struct {
	Next
}

func (d *Clinic) Do(p *patient) (err error) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	return
}

// Cashier 收费处处理器
type Cashier struct {
	Next
}

func (c *Cashier) Do(p *patient) (err error) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
		return
	}
	fmt.Println("Cashier getting money from patient patient")
	p.PaymentDone = true
	return
}

// Pharmacy 药房处理器
type Pharmacy struct {
	Next
}

func (m *Pharmacy) Do(p *patient) (err error) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		return
	}
	fmt.Println("Pharmacy giving medicine to patient")
	p.MedicineDone = true
	return
}

// StartHandler 不做操作，作为第一个Handler向下转发请求
type StartHandler struct {
	Next
}

// Do 空Handler的Do
func (h *StartHandler) Do(c *patient) (err error) {
	// 空Handler 这里什么也不做 只是载体 do nothing...
	return
}

func main() {
	patient := &patient{Name: "abc"}
	var s PatientHandler = &StartHandler{}
	s.SetNext(&ReceptionAction{}).
		SetNext(&Clinic{}).
		SetNext(&Cashier{}).
		SetNext(&Pharmacy{})
	if err := s.Execute(patient); err != nil {
		fmt.Println(err)
	}

	a := pkg.NewChan("ttt", &pkg.One{}, &pkg.Two{}, &pkg.Thread{}, &pkg.Four{}, &pkg.Five{})
	if err := a.Execute(a); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.Slice())
}
