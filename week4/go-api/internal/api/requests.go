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