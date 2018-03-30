package main

/*
import (
	"flag"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"time"
	"path"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}
func main() {
	var (
		err            error
		sftpClient     *sftp.Client
		username       string
		passwd         string
		host           string
		port           int
		localFileName  string
		remoteFilename string
	)
	flagParse(&host, &username, &passwd, &port, &localFileName, &remoteFilename)
	sftpClient, err = connect(username, passwd, host, port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	srcFile, err := sftpClient.Open(remoteFilename)
	defer srcFile.Close()
	if err != nil {
		log.Fatal(err, remoteFilename)
	}
	fmt.Println("Connection success!")
	if func(filename string) bool {
		_, err = os.Stat(filename)
		return err == nil || os.IsExist(err)
	}(localFileName) {
		os.Remove(localFileName)
	}
	destFile, err := os.Create(localFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	_, err = srcFile.WriteTo(destFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Copy file from remote server finished!")
	//distFile,err:=os.OpenFile(localFileName,os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	//checkErr(err)
	//srcFile.WriteTo(distFile)
	//var remoteFileName = path.Base(localFilePath)
	//dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer dstFile.Close()
	//
	//buf := make([]byte, 1024)
	//for {
	//	n, _ := srcFile.Read(buf)
	//	if n == 0 {
	//		break
	//	}
	//	dstFile.Write(buf)
	//}
}
func connect(user, passwd, host string, port int) (*sftp.Client, error) {
	var (
		//auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	//auth = make([]ssh.AuthMethod, 0)
	//auth = append(auth, ssh.Password(passwd))
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil

}
func flagParse(h, u, p *string, P *int, l, r *string) {
	flag.StringVar(h, "H", "127.0.0.1", "-H Host")
	flag.StringVar(p, "p", "", "passworld")
	flag.StringVar(u, "u", "", "username")
	flag.StringVar(l, "l", "/opt/dongxin.tar.gz", "local filename")
	flag.StringVar(r, "r", "/home/qacd/Release_Folder/BMS/Daily/"+time.Now().Format("20060102")+"/dongxin.tar.gz", "remote filename")
	flag.IntVar(P, "P", 22, "-P 22")

	flag.Parse()
	if *u == "" || *p == "" || *l == "" || *r == "" {
		flag.PrintDefaults()
		fmt.Println(`For example: sftpClient -H 192.168.1.1 -P 22 -u admin -p 123456 -r /opt/a.txt -l /home/b.txt`)
		os.Exit(1)
	}
}
*/
