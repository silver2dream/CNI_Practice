package main


import (
	"encoding/json"
	"fmt"
	"syscall"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/vishvananda/netlink"
	//"github.com/containernetworking/plugins/pkg/ip"
)

type SimpleBridge struct {
	BridgeName string `json:"bridgeName"`
	IP string `json:"ip"`
}


func cmdAdd(args *skel.CmdArgs) error {
	sb:= SimpleBridge{}
	if err:= json.Unmarshal(args.StdinData, &sb); err!=nil {
	               return err
	}
	fmt.Println(sb)	

	br:= &netlink.Bridge {
	        LinkAttrs: netlink.LinkAttrs {
		           Name:sb.BridgeName,
			   q
		           MTU:1500,
		           TxQLen:-1,
		}, 
	}    

       err:= netlink.LinkAdd(br)
       if err!= nil && err != syscall.EEXIST {
               return err
       }
      																	      
       
       if err :=netlink.LinkSetUp(br); err!=nil{																								                
		return err
	}

	fmt.Printf("interface Name: %s\n",args.IfName)
	fmt.Printf("netns path: %s\n", args.Netns)
	fmt.Printf("the config data: %s\n", args.StdinData)
	return nil
}

func cmdCheck(args *skel.CmdArgs) error {
	return nil

}

func cmdDel(args *skel.CmdArgs) error {
	return nil
}

func init(){
//	runtime.LockOSThread()
}

func main() {
	
	skel.PluginMain(cmdAdd, cmdCheck , cmdDel, version.All, "CNI noop plugin v0.7.0")

/*	netns, err:= ns.GetNS(args.Netns) 
	if err!=nil{
		return err
	}

	err:=netns.Do(func(hostNS ns.NetNS))error {
		hostVeth, containerVeth, err:= ip.SetupVeth(ifName,mtu,hostNS)
		if err!=nil{
			return err
		}
		guestIface.Name = containerVeth.Name
		guestIface.Mac = containerVeth.HardwareAddr.String()
		guestIface.Sandbox = netns.Path()
		hostIface.Name = hostVeth.Name
		return nil
	})

	link, err:=netlink.LinkByName(args.IfName)
	if err!=nil {
		return err
	}

	ipv4Addr, ipv4Net, err := net.ParseCIDR(sb.IP)
	addr := &netlink.Addr{
		IPNet:ipv4Net,
		Label:""
	}
	ipv4Net.IP = ipv4Addr
	if err = netlink.AddrAdd(link, addr); err != nil {
		return fmt.Errorf("failed to add IP addr %v to %q:%v",ipv4Net, args.IfName, err)
	}
	return nil*/
}

func createBridge(name string) (*netlink.Bridge, error) {
	br:= &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs {
			Name:name,
			MTU:1500,
			TxQLen:-1,
		},
	}

	err:= netlink.LinkAdd(br)
	if err!= nil && err != syscall.EEXIST {
		return nil,err
	}

	if err :=netlink.LinkSetUp(br); err!=nil{
		return nil,err
	}

	return br,nil
}

