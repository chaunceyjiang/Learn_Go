package main

import (
	"golang.org/x/crypto/ssh"
	"time"
	"net"
	"fmt"
	"log"
	"os"
	"flag"
)

func init()  {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}
func main()  {
	var (
		sess *ssh.Session
		err error
		cmd string
		user string
		passwd string
		port int
		host string
	)
	cmd=flagParse(&host,&user,&passwd,&port)
	sess,err=connect(user,passwd,host,port)
	defer sess.Close()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Connection success!")
	sess.Stderr=os.Stderr
	sess.Stdout=os.Stdout

	err=sess.Run(cmd)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Executed successfully")
}
func connect(user, passwd, host string, port int) (*ssh.Session, error) {
	var(
		auth []ssh.AuthMethod
		addr string
		sshClientConfig *ssh.ClientConfig
		sshClient *ssh.Client
		err error
		sshSession *ssh.Session
	)
	auth=[]ssh.AuthMethod{
		ssh.Password(passwd),
	}
	sshClientConfig=&ssh.ClientConfig{
		Auth:auth,
		User:user,
		Timeout:time.Second*30,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr=fmt.Sprintf("%s:%d",host,port)
	sshClient,err=ssh.Dial("tcp",addr,sshClientConfig)
	if err!=nil{
		return nil,err
	}
	if sshSession,err=sshClient.NewSession();err!=nil{
		return nil,err
	}
	return sshSession,nil
}
func flagParse(h, u, p *string, P *int) string {
	var cmd []string
	flag.StringVar(h,"H","","Host")
	flag.StringVar(u,"u","","username")
	flag.StringVar(p,"p","","password")
	flag.IntVar(P,"P",22,"port")
	flag.Parse()
	cmd=flag.Args()
	fmt.Println(cmd)
	if *h==""||*u==""||*p==""|| len(cmd)!=1{
		flag.PrintDefaults()
		fmt.Println("For example: sshexecute -H 192.169.2.50 -u root -p 123456 'ls -l;cd /opt/;touch a.txt'")
		os.Exit(1)
	}
	return cmd[0]
}