package main

import (
    "log"
    grpc "google.golang.org/grpc"
    context "golang.org/x/net/context"
    pb "etcdlet/etcdlet"
    "github.com/fsnotify/fsnotify"
    "fmt"
    "net"
    "os"
    "strings"
    "github.com/smallfish/simpleyaml"
    "io/ioutil"
    "encoding/json"
    "os/exec"
)

const (
    CONF_DIR = "/tmp/inventory"
    CLUSTER_SPEC = "/tmp/inventory/clusterspec"
)

func main(){
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()
    done := make(chan bool)
    //every node is a member
    s := grpc.NewServer()
    go member(s, "0.0.0.0")
    go func(){
        for{
            select{
                case event := <-watcher.Events:
                    log.Println("Got file event:", event)
                    if _,err := os.Stat(CLUSTER_SPEC); os.IsNotExist(err){
                    } else{
                        log.Println("modified file:", event.Name)
                        log.Println("Assuming role of leader")
                        //one with cluster spec is also a leader
                        leader(address)
                    }
            }
        }
    }()
    err = watcher.Add(CONF_DIR)
    if err != nil {
        log.Fatal(err)
    }
    <-done
}

func getURL(ip string, port string) (string){
    return "http://" + ip + ":" + port
}

func parseClusterSpec(file string) ([]*pb.BootstrapSpec){
    source, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatalf("Could not read cluster spec")
    }
    yaml, err := simpleyaml.NewYaml(source)
    if err != nil {
        log.Fatalf("Could not parse yaml file")
    }
    members := yaml.Get("clusterspec").Get("members")
    num, _ := members.GetArraySize()
    initialCluster := ""
    for i :=0 ; i< num ; i++ {
        member := members.GetIndex(i)
        name,_ := member.Get("name").String()
        address,_ := member.Get("address").String()
        initialCluster += (name + "=" + getURL(address, "2380") +",")
    }
    initialCluster = strings.Trim(initialCluster, ",")
    specArray := make ([]*pb.BootstrapSpec, num)
    for i :=0 ; i< num ; i++ {
        member := members.GetIndex(i)
        name,_ := member.Get("name").String()
        address,_ := member.Get("address").String()
        bootstrapSpec := new (pb.BootstrapSpec)
        bootstrapSpec.Name = name
        bootstrapSpec.Address = address
        bootstrapSpec.InitialAdvertisePeerUrls = getURL(address, "2380")
        bootstrapSpec.ListenPeerUrls = getURL(address, "2380")
        bootstrapSpec.ListenClientUrls = getURL(address, "2379") + "," + getURL("127.0.0.1", "2379")
        bootstrapSpec.AdvertiseClientUrls = getURL(address, "2379")
        bootstrapSpec.InitialClusterToken = "awesome-cluster"
        bootstrapSpec.InitialCluster = initialCluster
        bootstrapSpec.InitialClusterState = "new"
        specArray[i] = bootstrapSpec
    }
    return specArray
}

func leader(address string){
    specs := parseClusterSpec(CLUSTER_SPEC)
    for _, spec := range specs {
        transport, error := grpc.Dial(spec.Address +":5000", grpc.WithInsecure())
        if error != nil {
           log.Fatalf("No can do!")
        }
        client := pb.NewEtcdletClient(transport)
        client.Bootstrap(context.Background(), spec)
        transport.Close()
    }
}

func member(s *grpc.Server, address string){
    transport, error := net.Listen("tcp", address + ":5000")
    if error != nil {
        log.Fatalf("No can do!", error)
    }
    pb.RegisterEtcdletServer(s, &server{})
    s.Serve(transport)
}

type server struct {}

func createCmdString( spec *pb.BootstrapSpec)([]string){
    cmd := fmt.Sprintf("run --net=host -v /tmp/inventory:/tmp/inventory -v /data/etcd:/var/etcd/data gcr.io/google_containers/etcd:3.1.13 /usr/local/bin/etcd --name %s --initial-advertise-peer-urls %s --listen-peer-urls %s --listen-client-urls %s --advertise-client-urls %s --initial-cluster-token %s --initial-cluster %s --initial-cluster-state %s", spec.Name, spec.InitialAdvertisePeerUrls, spec.ListenPeerUrls, spec.ListenClientUrls, spec.AdvertiseClientUrls, spec.InitialClusterToken, spec.InitialCluster, spec.InitialClusterState)
    return strings.Fields(cmd)
}

func (s * server) Bootstrap(ctx context.Context, in *pb.BootstrapSpec) (*pb.Response, error){
    fmt.Println("Got bootstrap request")
    byteslice, _ := json.Marshal(in)
    fmt.Println(string(byteslice))
    args := createCmdString(in)
    fmt.Println(args)
    if err := exec.Command("docker", args[0:]...).Start(); err != nil {
        fmt.Println(err)
    }
    response := new(pb.Response)
    response.Status = "Success"
    return response, nil
}

func (s * server) Reconfigure(ctx context.Context, in *pb.MemberSpec) (*pb.Response, error){
    fmt.Println("Got request to reconfigure")
    response := new(pb.Response)
    response.Status = "Success"
    return response, nil
}

