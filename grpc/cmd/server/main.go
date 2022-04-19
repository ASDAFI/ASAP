package main

import (
	ASAP "ASAP/pb/proto"
	"ASAP/service"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"time"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemClientCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}
	sercetCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{sercetCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	return credentials.NewTLS(config), nil
}
func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port: %d", *port)
	farmServer:= service.NewConnectionServer()
	userStore := service.NewInMemoryUserStore()
	err := seedUsers(userStore)
	if err != nil {
		log.Fatal("cannot seed users: ")
	}
	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)
	interceptor := service.NewAuthInterceptor(jwtManager)
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load tls credentials: ", err)
	}
	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	ASAP.RegisterAuthServiceServer(grpcServer, authServer)
	ASAP.RegisterConnectionServiceServer(grpcServer, farmServer)
	reflection.Register(grpcServer)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)

	}
}

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/practice.LaptopService/"
	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
	}
}
func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "sayna", "secret")
	if err != nil {
		return fmt.Errorf("sayna's error: %v",err)
	}
	err = createUser(userStore, "armita", "secret")
	if err != nil {
		return fmt.Errorf("armita's error: %v",err)
	}
	err = createUser(userStore, "parnian", "secret")
	if err != nil {
		return fmt.Errorf("parnian's error: %v",err)
	}
	err = createUser(userStore, "ali", "secret")
	if err != nil {
		return fmt.Errorf("ali's error: %v",err)
	}
	return nil
}
func createUser(userStore service.UserStore, username, password string) error {
	user, err := service.NewUser(username, password)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}
