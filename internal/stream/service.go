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