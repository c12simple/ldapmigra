package cellsldap

import (
	"github.com/gogo/protobuf/test/indeximport-issue72"
	
)


type CellsLdapConfig struct {	
				
	SkipVerifyCertificate bool   `protobuf:"varint,7,opt,name=SkipVerifyCertificate" json:"SkipVerifyCertificate,omitempty"`
	RootCA                string `protobuf:"bytes,8,opt,name=RootCA" json:"RootCA,omitempty"`
	// To be converted to []byte
	RootCAData       string               `protobuf:"bytes,9,opt,name=RootCAData" json:"RootCAData,omitempty"`
	
	User struct {
		DNs []string
		Filter string
		IDAttribute string
		Scope string
	}
	MemberOfMapping  *MemberMapping

	//ConfigId		string //: "b81c85b7-6764-4cb8-885f-b7b2b49d2f7f",
	//DomainName example.org
	DomainName            string `protobuf:"bytes,2,opt,name=DomainName" json:"DomainName,omitempty"`
	Host                  string `protobuf:"bytes,3,opt,name=Host" json:"Host,omitempty"`

	// Connection: "normal"
	Connection            string `protobuf:"bytes,4,opt,name=Connection" json:"Connection,omitempty"`
	BindDN                string `protobuf:"bytes,5,opt,name=BindDN" json:"BindDN,omitempty"`
	BindPW                string `protobuf:"bytes,6,opt,name=BindPW" json:"BindPW,omitempty"`

	PageSize         int32                `protobuf:"varint,10,opt,name=PageSize" json:"PageSize,omitempty"`
	RolePrefix       string               `protobuf:"bytes,14,opt,name=RolePrefix" json:"RolePrefix,omitempty"`

	// Set default value
	Schedule         string               `protobuf:"bytes,15,opt,name=Schedule" json:"Schedule,omitempty"`
	SchedulerDetails string               `protobuf:"bytes,16,opt,name=SchedulerDetails" json:"SchedulerDetails,omitempty"`	

}

type MemberMapping struct {
	RealMemberOf bool
	GroupFilter struct {
		DNs []string
		DisplayAttribute string
		Filter string
		IDAttribute string
		Scope string
	}
	Mapping []struct {
		LeftAttribute string
		RightAttribute string
		RolePrefix string
	}
}



func (c *CellsLdapConfig) ConvertFromPydioConfig(pydioConfig map[string]string ){
	c.DomainName = ""
	// Connection
	if _, ok := pydioConfig["LDAP_PORT"]; ok {
		c.Host = pydioConfig["LDAP_URL"] + ":" + pydioConfig["LDAP_HOST"]	
	}else{
		c.Host = pydioConfig["LDAP_URL"] + ":389"
	}	
	c.BindDN = pydioConfig["LDAP_USER"]
	c.BindPW = pydioConfig["LDAP_PASSWORD"]

	// Todo
	c.DomainName = ""

	// Connection Type	
	switch pydioConfig["LDAP_PROTOCOL"] {
	case "ldap":
		c.Connection = "normal"	
	case "ldaps":
		c.Connection = "ssl"
	case "tls":
		c.Connection = "tls"
	default:
		c.Connection = "normal"
	}

	c.PageSize = 500


	// Mapping 
	const LOCAL_TYPE = "MAPPING_LOCAL_TYPE"
	const LOCAL_PARAM = "MAPPING_LOCAL_PARAM"
	const LDAP_PARAM = "MAPPING_LDAP_PARAM"
	
	var index string
	for i := 0; i < 10; i++ {
		if i == 0 {
			index = ""
		}else{
			index = string(i)
		}
		MappingLocalParam := LOCAL_PARAM + index

		if _,ok := pydioConfig[MappingLocalParam]; ok {
			c.MemberOfMapping.
		}
	}


}