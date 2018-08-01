package cellsldap


type CellsLdapConfig struct {	
		
	DomainName            string `protobuf:"bytes,2,opt,name=DomainName" json:"DomainName,omitempty"`
	Host                  string `protobuf:"bytes,3,opt,name=Host" json:"Host,omitempty"`
	//Connection            string `protobuf:"bytes,4,opt,name=Connection" json:"Connection,omitempty"`
	BindDN                string `protobuf:"bytes,5,opt,name=BindDN" json:"BindDN,omitempty"`
	BindPW                string `protobuf:"bytes,6,opt,name=BindPW" json:"BindPW,omitempty"`
	SkipVerifyCertificate bool   `protobuf:"varint,7,opt,name=SkipVerifyCertificate" json:"SkipVerifyCertificate,omitempty"`
	RootCA                string `protobuf:"bytes,8,opt,name=RootCA" json:"RootCA,omitempty"`
	// To be converted to []byte
	RootCAData       string               `protobuf:"bytes,9,opt,name=RootCAData" json:"RootCAData,omitempty"`
	PageSize         int32                `protobuf:"varint,10,opt,name=PageSize" json:"PageSize,omitempty"`
	User             *LdapSearchFilter    `protobuf:"bytes,11,opt,name=User" json:"User,omitempty"`
	MappingRules     []*LdapMapping       `protobuf:"bytes,12,rep,name=MappingRules" json:"MappingRules,omitempty"`
	MemberOfMapping  *LdapMemberOfMapping `protobuf:"bytes,13,opt,name=MemberOfMapping" json:"MemberOfMapping,omitempty"`
	RolePrefix       string               `protobuf:"bytes,14,opt,name=RolePrefix" json:"RolePrefix,omitempty"`
	Schedule         string               `protobuf:"bytes,15,opt,name=Schedule" json:"Schedule,omitempty"`
	SchedulerDetails string               `protobuf:"bytes,16,opt,name=SchedulerDetails" json:"SchedulerDetails,omitempty"`	
}

type MemberMapping struct {
	RealMemberOf bool
	GroupFIlter struct {
		DNs []string
		
	}
}

func (c *CellsLdapConfig) ConvertFromPydioConfig(pydioConfig map[string]string ){
	c.DomainName = ""
	// Connection
	c.Host = pydioConfig["LDAP_URL"].(string) 
	c.BindDN = pydioConfig[""]
}