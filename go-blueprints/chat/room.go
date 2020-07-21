package main

type room struct {
	// forward는 수신 메시지를 보관하는 채널이며
	// 수신한 메시지는 다른 클라이언트로 전달돼야 한다.
	forward chan []byte
	// join은 방에 들어오려는 클라이언트를 위한 채널이다.
	join chan *client
	// leave는 방을 나가길 원하는 클라이언트를 위한 채널이다.
	leave chan *client
	// clients는 현재 채팅방에 있는 모든 클라이언트를 보유한다.
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// 입장
			r.clients[client] = true
		case client := <-r.leave:
			// 퇴장
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// 모든 클라이언트에게 메시지 전달
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}
