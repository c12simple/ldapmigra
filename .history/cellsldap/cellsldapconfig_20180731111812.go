package cellsldap

import (
	"strings"
)


type Config struct {	
				
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
	Mappings []Mapping

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
	Mappings []Mapping
}


type Mapping struct {
	LeftAttribute  string
	RightAttribute string
	RuleString     string
	RolePrefix     string
}

func (c *Config) ConvertFromPydioConfig(pydioConfig map[string]string ){
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
	c.Schedule = "daily"
	c.SchedulerDetails = "3:00"

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
	c.RolePrefix = pydioConfig["LDAP_GROUP_PREFIX"]


	var index string
	// LDAP DNs
	for i:= 0; i < 10; i++{
		if i == 0 {
			index = ""
		}else{
			index = string(i)
		}
		if _,ok := pydioConfig["LDAP_DN" + index]; ok {
			c.User.DNs = append(c.User.DNs, pydioConfig["LDAP_DN" + index])
		}
	}
	c.User.Filter = pydioConfig["LDAP_FILTER"]
	c.User.IDAttribute = pydioConfig["LDAP_USERATTR"]
	c.User.Scope = "sub"

	// Mapping 
	const LOCAL_TYPE = "MAPPING_LOCAL_TYPE"
	const LOCAL_PARAM = "MAPPING_LOCAL_PARAM"
	const LDAP_PARAM = "MAPPING_LDAP_PARAM"
		
	for i := 0; i < 10; i++ {
		if i == 0 {
			index = ""
		}else{
			index = string(i)
		}
		MappingLocalParam 	:= LOCAL_PARAM + index
		MappingLocalType	:= LOCAL_TYPE + index
		MappingLdapParam 	:= LDAP_PARAM + index

		if strings.ToLower(pydioConfig[MappingLdapParam]) == "memberof"{
			// Mapping Group
			c.MemberOfMapping.GroupFilter.Filter = pydioConfig["LDAP_GROUP_FILTER"]
			c.MemberOfMapping.GroupFilter.IDAttribute = pydioConfig["LDAP_GROUPATTR"]
			gindex := ""
			for j:= 0; j < 10; j++ {
				if j == 0 {
					gindex = ""
				}else{
					gindex = string(j)
				}
				c.MemberOfMapping.GroupFilter.DNs = append(c.MemberOfMapping.GroupFilter.DNs, pydioConfig["LDAP_GDN" + gindex])
			}
			var mappingMemberOf Mapping
			mappingMemberOf.LeftAttribute = "memberOf"
			mappingMemberOf.RightAttribute = "Roles"
			mappingMemberOf.RolePrefix = c.RolePrefix
			mappingMemberOf.RuleString = pydioConfig[MappingLocalParam]
			c.MemberOfMapping.Mappings = append(c.MemberOfMapping.Mappings, mappingMemberOf)

			if _,ok := pydioConfig["LDAP_FAKE_MEMBEROF"]; !ok {
				c.MemberOfMapping.RealMemberOf = true
			}else{
				c.MemberOfMapping.
				"PydioMemberOfAttribute" = pydioConfig["LDAP_FAKE_MEMBEROF"]
				if _,ok := pydioConfig["LDAP_VALUE_MEMBERATTR_IN_GROUP"].(bool; ok && pydioConfig["LDAP_VALUE_MEMBERATTR_IN_GROUP"] == ""
				"PydioMemberOfValueFormat": "dn"				
			}
		}

		if _,ok := pydioConfig[MappingLocalType]; ok {
			c.Mappings[i].RightAttribute = convert(pydioConfig[MappingLocalType])
		}
		if _,ok := pydioConfig[MappingLdapParam]; ok {
			c.Mappings[i].LeftAttribute = convert(pydioConfig[MappingLdapParam])
		}
		if _,ok := pydioConfig[MappingLocalParam]; ok {
			c.Mappings[i].RuleString = convert(pydioConfig[MappingLocalParam])		
		}
	}
}

func convert(pydioldap string) string {
	switch pydioldap {
	case "role_id":
		return "Roles"
	case "core.conf/email":
		return "email"
	case "core.conf/USER_DISPLAY_NAME":
		return "displayName"
	default:			
		return pydioldap
	}
}