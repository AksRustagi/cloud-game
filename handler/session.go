package handler

import "github.com/giongto35/cloud-game/webrtc"

// Session represents a session connected from the browser to the current server
// It requires one connection to browser and one connection to the overlord
// connection to browser is 1-1. connection to overlord is n - 1
// Peerconnection can be from other server to ensure better latency
type Session struct {
	ID string
	//BrowserClient *BrowserClient
	//OverlordClient *OverlordClient
	peerconnection *webrtc.WebRTC

	// TODO: Decouple this
	//handler *Handler

	//ServerID    string
	//GameName    string
	RoomID string
	//PlayerIndex int
}
