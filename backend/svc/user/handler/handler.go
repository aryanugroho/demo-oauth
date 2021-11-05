package handler

import (
	"github.com/aryanugroho/demo-oauth/backend/svc/user/repository"
	hydraAdmin "github.com/ory/hydra-client-go/client/admin"
)

type Handler struct {
	HydraAdmin hydraAdmin.ClientService
	UserRepo   repository.Repository
}
