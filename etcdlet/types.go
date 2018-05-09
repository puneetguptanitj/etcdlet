package main

type Cluster struct {
	Clusterspec ClusterspecType `yaml:"clusterspec"`
}

type ClusterspecType struct {
	Members []SpecMember `yaml:"members"` 
}

type SpecMember struct {
	Name string `yaml:"name"`
	Address string `yaml:"address"`
}
