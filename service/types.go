package main

type Cluster struct {
	Clusterspec struct {
		Members []struct {
			Name    string `yaml:"name"`
			Address string `yaml:"address"`
		} `yaml:"members"`
	} `yaml:"clusterspec"`
}
