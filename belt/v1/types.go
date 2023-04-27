package v1

import (
	metav1 "github.com/zgjin/beltctl/apimachinery/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Config struct {
	metav1.TypeMeta `json:",inline"`
	Users           []*User    `json:"users,omitempty"`
	Contexts        []*Context `json:"contexts,omitempty"`
	CurrentContext  string     `json:"current_context,omitempty"`
}

type Cluster struct {
	Name   string `json:"name,omitempty"`
	Cert   string `json:"cert,omitempty"`
	Server string `json:"server,omitempty"`
}

type User struct {
	Name string `json:"name,omitempty"`
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}

type Context struct {
	Name        string `json:"name,omitempty"`
	ClusterName string `json:"cluster_name,omitempty"`
	UserName    string `json:"user_name,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Policy struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`
}

type PolicySpec struct {
	Description string            `json:"description,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}
