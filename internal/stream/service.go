package stream


import (
	"fmt"
	"io"
	"os"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/thienchuong/golang-file-upload/proto"
)

type Service struct {
	proto.UnimplementedFileUploadServiceServer
}

func (s *Service) DownloadFile(req *proto.DownloadFileRequest, stream proto.FileUploadService_DownloadFileServer) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}

	// open file
	file, err:= os.Open(req.Name)
	if err != nil {
		if os.IsNotExist(err) {
			return status.Error(codes.NotFound, "file not found")
		}
		return err
	}
	defer file.Close()

	const bufferSize = 5*1024 // send in 5kb chunks
	buff := make([]byte, bufferSize)

	for {
		bytesRead, err:= file.Read(buff)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
        // stream to cliens
		err:= stream.Send(&proto.DownloadFileResponse{Content: buff[:bytesRead]}) 
		if err != nil {
			return err
		}
	}


}

func (s *Service) UploadFile(stream proto.FileUploadService_UploadFileServer) error {
	
}