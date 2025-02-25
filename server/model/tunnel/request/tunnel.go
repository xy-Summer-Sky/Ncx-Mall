package request

// Create Tunnel Json structure
type CreateTunnel struct {
	Tunnelname 	string    	`json:"tunnelname" example:"ssh"`
	Tunneltype 	string    	`json:"tunneltype" example:"tcp"`
	Tunnelpoint string   	`json:"tunnelpoint" example:"shanghai"`
	Localip 	string   	`json:"localip" example:"127.0.0.1"`
	Localport 	uint16   	`json:"localport" example:"22"`
	Remoteip 	string   	`json:"remoteip" example:"139.254.14.13"`
	Remoteport 	uint16   	`json:"remoteport" example:"23"`
	Token 		string   	`json:"token" example:"123456"`
}

// Delete Tunnel Json structure
type DeleteTunnel struct {
	Tunnelid 	uint    `json:"tunnelId"`
}

// Find User All Tunnels Json structure
type FindUserAllTunnels struct {
	Userid		uint    `json:"userId"`
}