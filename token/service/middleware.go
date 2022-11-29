package service

import service "github.com/IRFAN374/gojwtsvc/token"

type Middleware func(service.Service) service.Service
