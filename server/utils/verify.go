package utils

var (
	// system service
	IdVerify               	= Rules{"ID": []string{NotEmpty()}}
	ApiVerify              	= Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             	= Rules{"Path": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         	= Rules{"Title": {NotEmpty()}}
	LoginVerify            	= Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         	= Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify         	= Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         	= Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         	= Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}}
	AutoPackageVerify      	= Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify        	= Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify      	= Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     	= Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   	= Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify 	= Rules{"AuthorityId": {NotEmpty()}}

	// shop service
	CreateShopOrderVerify 	= Rules{"ServiceType": {NotEmpty(), Ge("1"), Le("2")}, "Price": {NotEmpty(), Ge("0.01")}, "Status": {Ge("0"), Le("2")}}
	DeleteShopOrderVerify 	= Rules{"OrderID": {NotEmpty()}}
	GetUserOrdersVerify  	= Rules{"UserID": {NotEmpty()}}

	// tunnel service
	CreateTunnelVerify 		= Rules{"Tunnelname": {NotEmpty()}, "Tunneltype": {NotEmpty()}, "Tunnelpoint": {NotEmpty()}, "Localip": {NotEmpty()}, "Localport": {NotEmpty(), Ge("1"), Le("65535")}, "Remoteip": {NotEmpty()}, "Remoteport": {NotEmpty(), Ge("1"), Le("65535")}, "Token": {NotEmpty()}}
	DeleteTunnelVerify 		= Rules{"Tunnelid": {NotEmpty()}}
	FindUserAllTunnelsVerify = Rules{"Userid": {NotEmpty()}}
)
