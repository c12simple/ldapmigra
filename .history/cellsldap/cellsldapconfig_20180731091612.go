package cellsldap

import (
	
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
	c.Host = pydioConfig["LDAP_URL"]
	c.BindDN = pydioConfig["LDAP_USER"]
	c.BindPW = pydioConfig["LDAP_PASSWORD"]

	// Todo
	c.DomainName = ""

	// Connection Type
	var connectionType string
	switch pydioConfig["LDAP_PROTOCOL"] {
	case "ldap":
		c.Connection = "normal"	
	case "ldaps":
		c.Connection = "ssl"
	case "tls":
		c.Connection = "tls"
	}

	c.PageSize = int32500


	c.Connection = 


}