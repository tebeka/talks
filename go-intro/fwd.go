// forward proxies traffic between local socket and remote backend
func forward(local net.Conn, remoteAddr string) {
	remote, err := net.Dial("tcp", remoteAddr)
	if remote == nil {
		log.Printf("remote dial failed: %v\n", err)
		local.Close()
		return
	}
	go io.Copy(local, remote)
	go io.Copy(remote, local)
}
