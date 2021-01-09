package enquiry

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {

}

func (s *Server) ReturnTrainDetails (ctx context.Context, message *Message) (*EnquiryResponse, error){
	log.Printf("Recieved message from the grpc server %v", message)
	t := &TrainSchedule{
		StationName:   "DDL",
		Arrival:       "8:00",
		Departure:     "12:00",
	}
	respSlice := make([]*TrainSchedule, 0, 2)
	respSlice = append(respSlice, t)
	return &EnquiryResponse{
		Resp: respSlice,
	}, nil
}
