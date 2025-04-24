package server

type App struct {
	GrpcServer *GRPCServer
	HttpServer *HTTPServer
}

func NewApp(grpcServer *GRPCServer, httpServer *HTTPServer) *App {
	return &App{
		GrpcServer: grpcServer,
		HttpServer: httpServer,
	}
}
