package gcfping

import (
	"fmt"
	"net/http"

	"github.com/go-ping/ping"
)

type Address struct {
	Addr string `json:"addr"`
}

//Ping test conneciton
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starrting....")
	var addr Address
	err := json.NewDecoder(r.Body).Decode(&Address{})
	if err != nil {
		//log.Fatal(err)
		json.NewEncoder(w).Encode(err.Error())
	
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		//panic(err)
		fmt.Println("Error creating pinger")
		json.NewEncoder(w).Encode(err.Error())
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats

	fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
	fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
		stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
		stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())

}
