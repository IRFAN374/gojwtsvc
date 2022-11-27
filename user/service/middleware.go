package service

import service "github.com/IRFAN374/gojwtsvc/user"

type Middleware func(service.Service) service.Service
