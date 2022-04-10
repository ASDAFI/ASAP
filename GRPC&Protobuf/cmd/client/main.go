package main

import (
	"ASAP/client"
	ASAP "ASAP/pb/proto"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"time"
)

const (
	refreshDuration = 30 * time.Second
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	fmt.Print("Enter username: ")
	// var then variable name then variable type
	var username string
	// Taking input from user
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	var password string
	fmt.Scanln(&password)
	authClient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}
	cc2, err := grpc.Dial(*serverAddress,
		grpc.WithTransportCredentials(tlsCredentials),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	//log.Println(cc2)
	conn := client.NewConnectionClient(cc2)
	testConnection(conn)
	//testSearchLaptop(laptopClient)
}

func testConnection(conn *client.ConnectionClient) {
	conn.Connect(&ASAP.Connection{Request: "salam"})
}

func authMethods() map[string]bool {
	const farmServicePath = "/GRPC&Protobuf.ConnectionService/"
	const AuthServicePath = "/GRPC&Protobuf.AuthService/"
	return map[string]bool{
		farmServicePath + "Connect": true,
		AuthServicePath + "Login":   true,
	}
}
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}
	return credentials.NewTLS(config), nil
}
