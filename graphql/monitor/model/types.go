package model

//create table if not exists monitor
//(
//    id                 int auto_increment comment 'id'
//        primary key,
//    deployment_name    varchar(64) null comment 'deployment名称',
//    current_replicas   int         null comment '当前副本',
//    available_replicas int         null comment '可用副本',
//    cpu_limit          int         null comment 'cpu_limit值，单位m',
//    cpu_request        int         null comment 'cpu_request值，单位m',
//    mem_limit          mediumtext  null comment '内存limit值，单位byte',
//    mem_request        mediumtext  null comment '内存request值，单位byte',
//    namespace          varchar(64) null comment '命名空间'
//)
//    comment '监控表';

//INSERT INTO graphql.monitor (id, deployment_name, current_replicas, available_replicas, cpu_limit, cpu_request, mem_limit, mem_request, namespace) VALUES (1, 'apiserver', 1, 1, 2000, 2000, '4294967296', '4294967296', 'kube-system');
//INSERT INTO graphql.monitor (id, deployment_name, current_replicas, available_replicas, cpu_limit, cpu_request, mem_limit, mem_request, namespace) VALUES (2, 'controllermanager', 1, 1, 2000, 2000, '2147483648', '2147483648', 'kube-system');
//INSERT INTO graphql.monitor (id, deployment_name, current_replicas, available_replicas, cpu_limit, cpu_request, mem_limit, mem_request, namespace) VALUES (3, 'scheduler', 1, 1, 1000, 1000, '2147483648', '2147483648', 'kube-system');
//INSERT INTO graphql.monitor (id, deployment_name, current_replicas, available_replicas, cpu_limit, cpu_request, mem_limit, mem_request, namespace) VALUES (4, 'platform-api', 2, 2, 4000, 4000, '4294967296', '4294967296', 'platform');
//INSERT INTO graphql.monitor (id, deployment_name, current_replicas, available_replicas, cpu_limit, cpu_request, mem_limit, mem_request, namespace) VALUES (5, 'platform-listener', 1, 1, 1000, 1000, '1073741824', '1073741824', 'platform');
type Tabler interface {
	TableName() string
}
func (Monitor) TableName() string {
	return "monitor"
}

type Monitor struct {
	ID                int64  `json:"id" gorm:"AUTO_INCREMENT"`
	DeploymentName    string `json:"deploymentName"`
	CurrentReplicas   int32  `json:"currentReplicas,omitempty"`
	AvailableReplicas int32  `json:"availableReplicas"`
	CpuLimit          int32  `json:"cpuLimit"`
	CpuRequest        int32  `json:"cpuRequest"`
	MemLimit          int64  `json:"memLimit"`
	MemRequest        int64  `json:"memRequest"`
	Namespace         string `json:"namespace"`
}
