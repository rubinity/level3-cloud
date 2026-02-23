package api


type ReplicationRequest struct {
    Size int  `json:"size"`
    Namespace string  `json:"namespace"`
	Name string  `json:"name"`
}


type DeleteRequest struct {
    Namespace string  `json:"namespace"`
	Name string  `json:"name"`
}

type RRInfo struct {
    Kind   string      `json:"kind"`
    Name   string      `json:"name"`
    Status interface{} `json:"status"`
}

type Namespace struct {
	// ID       string `json:"id"`
	Namespace string `json:"namespace"`
	Password string `json:"password"`
}


