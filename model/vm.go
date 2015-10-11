package model

type Vm struct {
	ID        GUID   `redis:"id"`
	Name      string `redis:"name"`
	Status    Status `redis:"status"`
	ClusterID GUID   `redis:"cluster_id"`
}
