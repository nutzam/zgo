package z

import (
	"bytes"
	"net"
	"os"
	"time"
)

const (
	ICMP_ECHO_REQUEST = 8
	ICMP_ECHO_REPLY   = 0
)

// Ping Request
func makePingRequest(id, seq, pktlen int, filler []byte) []byte {
	p := make([]byte, pktlen)
	copy(p[8:], bytes.Repeat(filler, (pktlen-8)/len(filler)+1))

	p[0] = ICMP_ECHO_REQUEST // type
	p[1] = 0                 // code
	p[2] = 0                 // cksum
	p[3] = 0                 // cksum
	p[4] = uint8(id >> 8)    // id
	p[5] = uint8(id & 0xff)  // id
	p[6] = uint8(seq >> 8)   // sequence
	p[7] = uint8(seq & 0xff) // sequence

	// calculate icmp checksum
	cklen := len(p)
	s := uint32(0)
	for i := 0; i < (cklen - 1); i += 2 {
		s += uint32(p[i+1])<<8 | uint32(p[i])
	}
	if cklen&1 == 1 {
		s += uint32(p[cklen-1])
	}
	s = (s >> 16) + (s & 0xffff)
	s = s + (s >> 16)

	// place checksum back in header; using ^= avoids the
	// assumption the checksum bytes are zero
	p[2] ^= uint8(^s & 0xff)
	p[3] ^= uint8(^s >> 8)

	return p
}

func parsePingReply(p []byte) (id, seq int) {
	id = int(p[4])<<8 | int(p[5])
	seq = int(p[6])<<8 | int(p[7])
	return
}

// Ping
func Ping(addr string, i int) bool {

	// *IPAddr
	raddr, e := net.ResolveIPAddr("ip4", addr)
	if e != nil {
		return false
	}

	// *IPConn
	ipconn, ee := net.DialIP("ip4:icmp", nil, raddr)
	if ee != nil {
		return false
	}

	// 保证连接正常关闭
	defer ipconn.Close()

	// PID
	sendid := os.Getpid() & 0xffff
	sendseq := 1
	pingpktlen := 64

	for {

		sendpkt := makePingRequest(sendid, sendseq, pingpktlen, []byte("Go Ping"))

		// 发送请求
		n, err := ipconn.WriteToIP(sendpkt, raddr)
		if err != nil || n != pingpktlen {
			break
		}

		// 超时
		ipconn.SetDeadline(time.Now().Add(5 * time.Second))

		// 返回数据
		resp := make([]byte, 1024)
		for {

			// 读取返回
			_, _, err := ipconn.ReadFrom(resp)
			if err != nil {
				break
			}

			// 判断状态
			if resp[0] != ICMP_ECHO_REPLY {
				continue
			}

			// 判断状态
			rcvid, rcvseq := parsePingReply(resp)
			if rcvid != sendid || rcvseq != sendseq {
				break
			}

			// 成功返回
			return true

		}

		// 执行次数内未成功返回
		if i == sendseq {
			break
		}

		// 计数器
		sendseq++

	}

	// 失败返回
	return false
}
