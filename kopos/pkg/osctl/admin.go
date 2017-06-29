package osctl

import (
	"database/sql"
	"fmt"
	"log"
	// "os"
	"strings"

	"github.com/golang/glog"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	_ "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	_ "github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	_ "github.com/gophercloud/gophercloud/openstack/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type AdminInfra struct {
	InfraProvider
}

func Admin() *AdminInfra {
	return new(AdminInfra)
}

type User struct {
	Name string
	Age  int
}

func pq_tut() {
	/*
		 A backslash will escape the next character in values:
		"user=space\ man password='it\'s valid'
	*/
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	// db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	println(rows)

	rows, err = db.Query(`SELECT name FROM users WHERE favorite_fruit = $1
	OR age BETWEEN $2 AND $2 + 3`, "orange", 64)
	println(rows)

	/*
	   pq does not support the LastInsertId() method of the Result type in database/sql.
	   To return the identifier of an INSERT (or UPDATE or DELETE),
	   use the Postgres RETURNING clause with a standard Query or QueryRow call:
	*/
	var userid int
	err = db.QueryRow(`INSERT INTO users(name, favorite_fruit, age)
	VALUES('beatrice', 'starfruit', 93) RETURNING id`).Scan(&userid)

	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}

	// Bulk imports
	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn("users", "name", "age"))
	if err != nil {
		log.Fatal(err)
	}

	users := []User{{"foor", 19}, {"bar", 20}}

	for _, user := range users {
		_, err = stmt.Exec(user.Name, int64(user.Age))
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

type Account struct {
	ID      int `gorm:"primary_key"`
	Balance int
}

func gorm_tut() {
	// ORM
	const addr = "postgresql://maxroach@localhost:26257/bank?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automatically create the "accounts" table based on the Account model.
	db.AutoMigrate(&Account{})

	// Insert two rows into the "accounts" table.
	db.Create(&Account{ID: 1, Balance: 1000})
	db.Create(&Account{ID: 2, Balance: 250})

	// Print out the balances.
	var accounts []Account
	db.Find(&accounts)
	fmt.Println("Initial balances:")
	for _, account := range accounts {
		fmt.Printf("%d %d\n", account.ID, account.Balance)
	}
}

func (admin *AdminInfra) CreateSharedNet(name, tenandid, subnetcidr, gatewayip, description string) {
	adminstateup := true
	shared := true
	network, err := admin.tryto().createNetwork(name, tenandid, adminstateup, shared)
	if nil != err {
		return
	}

	offset := strings.LastIndex(subnetcidr, ".")
	ipprefix := subnetcidr[:offset]
	ipversion := gophercloud.IPv4
	enabledhcp := true
	dnsnameservers := []string{}
	hostroutes := []subnets.HostRoute{}
	subnet, err := admin.tryto().createSubnet(network.ID, subnetcidr, name, tenandid, []subnets.AllocationPool{{ipprefix + ".50", ipprefix + ".200"}}, gatewayip, ipversion, enabledhcp, dnsnameservers, hostroutes)
	if nil != err {
		return
	}
	println(subnet)

	distributed := false
	gatewayinfo := routers.GatewayInfo{}
	router, err := admin.tryto().createRouter(name, adminstateup, distributed, tenandid, gatewayinfo)
	if nil != err {
		return
	}

	secgroup, err := admin.tryto().createSecurityGroup(name, tenandid, description)
	if nil != err {
		return
	}

	// remoteipprefix = "::/0"

	securitygroups := []string{}
	portrangemax := 65535
	portrangemin := 0
	protocol := rules.ProtocolTCP
	remotegroupid := ""
	remoteipprefix := "0.0.0.0/0"
	secgrouprule, err := admin.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	protocol = rules.ProtocolUDP
	secgrouprule, err = admin.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	protocol = rules.ProtocolICMP
	portrangemax = -1
	portrangemin = -1
	secgrouprule, err = admin.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	macaddress := ""
	fixedips := []ports.IP{}
	deviceid := ""
	deviceowner := ""
	allowedaddresspairs := []ports.AddressPair{}
	port, err := admin.tryto().createPort(network.ID, name, adminstateup, macaddress, fixedips, deviceid, deviceowner, tenandid, securitygroups, allowedaddresspairs)
	if nil != err {
		return
	}

	ifinfo, err := admin.tryto().plugRouterIntoSubnet(router.ID, port.ID)
	if nil != err {
		return
	}
	println(ifinfo)
}

func (admin *AdminInfra) tryto() *AdminInfra {
	if nil == admin.providerclient {
		if opts, err := openstack.AuthOptionsFromEnv(); nil != err {
			glog.Errorf("Could not load admin openrc: %v", err)
			admin.lasterr = err
		} else {
			glog.Infof("Load admin openrc: %v %v", opts.IdentityEndpoint, opts.Username)
			if provider, err := openstack.AuthenticatedClient(opts); nil != err {
				glog.Errorf("Could not authenticate admin: %v", err)
				admin.lasterr = err
			} else {
				glog.Infof("Authenticated for token: %v", provider.TokenID)
				admin.lasterr = nil
				admin.providerclient = provider
			}
		}
	}
	return admin
}

func (admin *AdminInfra) IdentityEndpoint(url string) *AdminInfra {
	admin.identityendpoint = url
	return admin
}

func (admin *AdminInfra) BasicAuthCredential(username, password string) *AdminInfra {
	admin.username, admin.password = username, password
	return admin
}
