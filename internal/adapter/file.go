package adapter

import (
	filespb "github.com/Confialink/wallet-files/rpc/files"
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type File struct {
	rpcFiles *connection.RpcFiles
}

func NewFile() *File {
	return &File{
		rpcFiles: connection.GetRpcFiles(),
	}
}

func (s *File) FindById(ctx context.Context, id uint64) (*model.File, error) {
	request := &filespb.FileReq{
		Id: id,
	}

	response, err := s.rpcFiles.Client.GetFile(ctx, request)
	if err != nil {
		return nil, internal_errors.CreateError(err, internal_errors.FileNotFound, "")
	}

	file := s.convertFile(*response)
	return &file, nil
}

func (s File) convertFile(file filespb.FileResp) model.File {
	object := model.File{
		Id:       file.Id,
		Location: file.Location,
	}
	return object
}
