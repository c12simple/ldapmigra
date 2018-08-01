package cellsldap

import (
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	SkipVerifyCertificate bool   `protobuf:"varint,7,opt,name=SkipVerifyCertificate" json:"SkipVerifyCertificate,omitempty"`
	RootCA                string `protobuf:"bytes,8,opt,name=RootCA" json:"RootCA,omitempty"`
	// To be converted to []byte
	RootCAData string `protobuf:"bytes,9,opt,name=RootCAData" json:"RootCAData,omitempty"`

	User struct {
		DNs         []string
		Filter      string
		IDAttribute string
		Scope       string
	}
	MemberOfMapping MemberMapping
	Mappings        []Mapping

	//ConfigId		string //: "b81c85b7-6764-4cb8-885f-b7b2b49d2f7f",
	//DomainName example.org
	DomainName string `protobuf:"bytes,2,opt,name=DomainName" json:"DomainName,omitempty"`
	Host       string `protobuf:"bytes,3,opt,name=Host" json:"Host,omitempty"`

	// Connection: "normal"
	Connection string `protobuf:"bytes,4,opt,name=Connection" json:"Connection,omitempty"`
	BindDN     string `protobuf:"bytes,5,opt,name=BindDN" json:"BindDN,omitempty"`
	BindPW     string `protobuf:"bytes,6,opt,name=BindPW" json:"BindPW,omitempty"`

	PageSize   int32  `protobuf:"varint,10,opt,name=PageSize" json:"PageSize,omitempty"`
	RolePrefix string `protobuf:"bytes,14,opt,name=RolePrefix" json:"RolePrefix,omitempty"`

	// Set default value
	Schedule         string `protobuf:"bytes,15,opt,name=Schedule" json:"Schedule,omitempty"`
	SchedulerDetails string `protobuf:"bytes,16,opt,name=SchedulerDetails" json:"SchedulerDetails,omitempty"`
}

type MemberMapping struct {
	RealMemberOf bool
	GroupFilter  struct {
		DNs              []string
		DisplayAttribute string
		Filter           string
		IDAttribute      string
		Scope            string
	}
	Mappings                 []Mapping
	PydioMemberOfAttribute   string
	PydioMemberOfValueFormat string
}

type Mapping struct {
	LeftAttribute  string
	RightAttribute string
	RuleString     string
	RolePrefix     string
}

func (c *Config) ConvertFromPydioConfig(pydioConfig map[string]interface{}) {

	c.DomainName = ""
	// Connection
	if _, ok := pydioConfig["LDAP_PORT"].(string); ok {
		c.Host = pydioConfig["LDAP_URL"].(string) + ":" + pydioConfig["LDAP_PORT"].(string)
	} else {
		c.Host = pydioConfig["LDAP_URL"].(string) + ":389"
	}
	c.BindDN = pydioConfig["LDAP_USER"].(string)
	c.BindPW = pydioConfig["LDAP_PASSWORD"].(string)

	// Todo
	c.DomainName = ""
	c.Schedule = "daily"
	c.SchedulerDetails = "3:00"

	// Connection Type
	switch pydioConfig["LDAP_PROTOCOL"].(string) {
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
	c.RolePrefix = pydioConfig["LDAP_GROUP_PREFIX"].(string)

	var index string
	// LDAP DNs
	for i := 0; i < 10; i++ {
		if i == 0 {
			index = ""
		} else {
			index = strconv.Itoa(i)
		}
		if _, ok := pydioConfig["LDAP_DN"+index]; ok {
			c.User.DNs = append(c.User.DNs, pydioConfig["LDAP_DN"+index].(string))
		}
	}
	c.User.Filter = pydioConfig["LDAP_FILTER"].(string)
	c.User.IDAttribute = pydioConfig["LDAP_USERATTR"].(string)
	c.User.Scope = "sub"

	// Mapping
	const LOCAL_TYPE = "MAPPING_LOCAL_TYPE"
	const LOCAL_PARAM = "MAPPING_LOCAL_PARAM"
	const LDAP_PARAM = "MAPPING_LDAP_PARAM"

	mappingLoop := true
	for i := 0; mappingLoop; i++ {
		if i == 0 {
			index = ""
		} else {
			index = "_" + strconv.Itoa(i)
		}
		MappingLocalParam := LOCAL_PARAM + index
		MappingLocalType := LOCAL_TYPE + index
		MappingLdapParam := LDAP_PARAM + index

		fmt.Println("loop: " + index + " : " + MappingLdapParam)

		if _, ok := pydioConfig[MappingLdapParam]; !ok {
			mappingLoop = false
			break
		}

		if strings.ToLower(pydioConfig[MappingLdapParam].(string)) == "memberof" {
			// Mapping Group

			c.MemberOfMapping.GroupFilter.Filter = pydioConfig["LDAP_GROUP_FILTER"].(string)
			c.MemberOfMapping.GroupFilter.IDAttribute = pydioConfig["LDAP_GROUPATTR"].(string)
			gindex := ""
			for j := 0; j < 10; j++ {
				if j == 0 {
					gindex = ""
				} else {
					gindex = "_" + strconv.Itoa(j)
				}
				if pydioConfig["LDAP_GDN"+gindex] !=  nil {
					c.MemberOfMapping.GroupFilter.DNs = append(c.MemberOfMapping.GroupFilter.DNs, pydioConfig["LDAP_GDN"+gindex].(string))
				}
			}
			var mappingMemberOf Mapping
			mappingMemberOf.LeftAttribute = "memberOf"
			mappingMemberOf.RightAttribute = "Roles"
			mappingMemberOf.RolePrefix = c.RolePrefix
			mappingMemberOf.RuleString = pydioConfig[MappingLocalParam].(string)
			c.MemberOfMapping.Mappings = append(c.MemberOfMapping.Mappings, mappingMemberOf)

			if _, ok := pydioConfig["LDAP_FAKE_MEMBEROF"]; !ok {
				c.MemberOfMapping.RealMemberOf = true
			} else {
				c.MemberOfMapping.PydioMemberOfAttribute = pydioConfig["LDAP_FAKE_MEMBEROF"].(string)
				if _, ok := pydioConfig["LDAP_VALUE_MEMBERATTR_IN_GROUP"]; ok && pydioConfig["LDAP_VALUE_MEMBERATTR_IN_GROUP"].(bool) {
					c.MemberOfMapping.PydioMemberOfValueFormat = "dn"
				} else {
					c.MemberOfMapping.PydioMemberOfValueFormat = "cn"
				}
			}
		}

		fmt.Println(MappingLdapParam + ":" + pydioConfig[MappingLdapParam].(string))
		fmt.Println(MappingLocalType + ":" + pydioConfig[MappingLocalType].(string))
		fmt.Println(MappingLocalParam + ":" + pydioConfig[MappingLocalParam].(string))

		if _, ok := pydioConfig[MappingLocalParam]; ok {
			fmt.Println("==>" + pydioConfig[MappingLocalParam].(string))
			//c.Mappings[i].RuleString = pydioConfig[MappingLocalParam].(string)
		}

		if _, ok := pydioConfig[MappingLdapParam]; ok {
			fmt.Println("==>" + pydioConfig[MappingLdapParam].(string))
			//c.Mappings[i].LeftAttribute = convert(pydioConfig[MappingLdapParam].(string))
		}
		if _, ok := pydioConfig[MappingLocalType]; ok {
			fmt.Println("==>" + pydioConfig[MappingLocalType].(string))
			//c.Mappings[i].RuleString = convert(pydioConfig[MappingLocalType].(string))
		}

	}
	fmt.Printf("%v", c)
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

func castString(str interface{}) string {
	switch str.(type) {
	case string:
		return str.(string)
	default:
		return ""
	}
}
