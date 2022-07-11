package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc/contact/contactpb"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	connectStr := "root:123456@tcp(127.0.0.1:3306)/contact?charset=utf8"
	// maxIdle := 100
	// maxConn := 100
	err := orm.RegisterDataBase("default", "mysql", connectStr)
	if err != nil {
		log.Panic("register db err %v", err)
	}
	orm.RegisterModel(new(ContactInfo))

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Panicf("run migrate db err %v", err)
	}
	fmt.Println("register connect db sucessfully!")
}

type server struct{}

func (server) Insert(ctx context.Context, req *contactpb.InsertRequest) (*contactpb.InsertResponse, error) {
	log.Printf("calling insert %+v\n", req.Contact)
	ci := ConvertPbContact2ContactInfo(req.Contact)

	err := ci.Insert()

	if err != nil {
		resp := &contactpb.InsertResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("insert err %v", err),
		}
		return resp, nil
		// return nil, status.Errorf(codes.InvalidArgument, "Insert %+v err %v", ci, err)
	}

	resp := &contactpb.InsertResponse{
		StatusCode: 1,
		Message:    "OK",
	}

	return resp, nil
}

func (server) Read(ctx context.Context, req *contactpb.ReadRequest) (*contactpb.ReadResponse, error) {
	log.Printf("calling read %s\n", req.GetPhoneNumber())
	ci, err := Read(req.GetPhoneNumber())
	if err == orm.ErrNoRows {
		return nil, status.Errorf(codes.InvalidArgument, "Phone %s not exist", req.GetPhoneNumber())
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Read phone %s err %v", req.GetPhoneNumber(), err)
	}

	return &contactpb.ReadResponse{
		Contact: ConvertContactInfo2PbContact(ci),
	}, nil

}

func (server) Update(ctx context.Context, req *contactpb.UpdateRequest) (*contactpb.UpdateResponse, error) {
	if req.GetNewContact() == nil {
		return nil, status.Error(codes.InvalidArgument, "update req with nil contact")
	}
	log.Printf("calling update with data %v\n", req.GetNewContact())
	ci := ConvertPbContact2ContactInfo(req.GetNewContact())
	err := ci.Update()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "update %+v err %v", req.GetNewContact(), err)
	}

	updateContact, err := Read(req.GetNewContact().GetPhoneNumber())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "try to read update contact %+v err %v", req.GetNewContact(), err)
	}

	return &contactpb.UpdateResponse{
		UpdateContact: ConvertContactInfo2PbContact(updateContact),
	}, nil
}

func (server) Delete(ctx context.Context, req *contactpb.DeleteRequest) (*contactpb.DeleteResponse, error) {
	log.Printf("calling delete %s\n", req.GetPhoneNumber())
	if len(req.GetPhoneNumber()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Request delete with empty phone number")
	}

	ci := &ContactInfo{
		PhoneNumber: req.GetPhoneNumber(),
	}
	err := ci.Delete()
	if err != nil {
		return &contactpb.DeleteResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("delete contact %s err %v", req.GetPhoneNumber(), err),
		}, nil
	}

	return &contactpb.DeleteResponse{
		StatusCode: 1,
		Message:    fmt.Sprintf("delete contact %s successfully", req.GetPhoneNumber()),
	}, nil
}

func (server) Search(ctx context.Context, req *contactpb.SearchRequest) (*contactpb.SearchResponse, error) {
	log.Printf("calling search %s\n", req.GetSearchName())
	if len(req.GetSearchName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Request search with empty phone number")
	}

	listCi, err := SearchByName(req.GetSearchName())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Request search %s err %v", req.GetSearchName(), err)
	}

	listPbContact := []*contactpb.Contact{}
	for _, ci := range listCi {
		pbContact := ConvertContactInfo2PbContact(ci)
		listPbContact = append(listPbContact, pbContact)
	}

	return &contactpb.SearchResponse{
		Results: listPbContact,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50070")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()
	contactpb.RegisterContactServiceServer(s, &server{})

	fmt.Println("contact service is running....")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while server %v", err)
	}

}

///////// Models

type ContactInfo struct {
	PhoneNumber string `orm:"size(15);pk"`
	Name        string
	Address     string `orm:"type(text)"`
}

func ConvertPbContact2ContactInfo(pbContact *contactpb.Contact) *ContactInfo {
	return &ContactInfo{
		PhoneNumber: pbContact.PhoneNumber,
		Name:        pbContact.Name,
		Address:     pbContact.Address,
	}
}

func ConvertContactInfo2PbContact(ci *ContactInfo) *contactpb.Contact {
	return &contactpb.Contact{
		PhoneNumber: ci.PhoneNumber,
		Name:        ci.Name,
		Address:     ci.Address,
	}
}

func (c *ContactInfo) Insert() error {
	o := orm.NewOrm()

	_, err := o.Insert(c)
	if err != nil {
		log.Print("Insert contact %+v err %v", c, err)
		return err
	}
	log.Printf("Insert %+v successfully", c)
	return nil
}

func Read(phoneNumber string) (*ContactInfo, error) {
	o := orm.NewOrm()
	ci := &ContactInfo{
		PhoneNumber: phoneNumber,
	}
	err := o.Read(ci)
	if err != nil {
		log.Printf("Read contact %+v err %v\n", ci, err)
		return nil, err
	}

	return ci, nil
}

func (c *ContactInfo) Update() error {
	o := orm.NewOrm()

	num, err := o.Update(c, "Name")
	if err != nil {
		log.Printf("Update %+v err %v\n", c, err)
		return err
	}
	log.Printf("update contact %+v, affect %d row\n", c, num)
	return nil
}

func (c *ContactInfo) Delete() error {
	o := orm.NewOrm()

	num, err := o.Delete(c)
	if err != nil {
		log.Printf("delete %+v err %v\n", c, err)
		return err
	}
	log.Printf("delete contact %+v, affect %d row\n", c, num)
	return nil
}

func SearchByName(name string) ([]*ContactInfo, error) {
	result := []*ContactInfo{}
	o := orm.NewOrm()

	num, err := o.QueryTable(new(ContactInfo)).Filter("name__icontains", name).All(&result)

	if err == orm.ErrNoRows {
		log.Printf("search %s found no rows\n", name)
		return result, nil
	}

	if err != nil {
		log.Printf("search %s err %v\n", name, err)
		return nil, err
	}

	log.Printf("search %s found %d rows\n", name, num)
	return result, nil
}
