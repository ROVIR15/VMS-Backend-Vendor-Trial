package constants

// Access constants
type AccessRole string
type AccessPermission string

const (
	// Object
	AccessObjAuth     string = "auth"
	AccessObjBrand    string = "brand"
	AccessObjHotel    string = "hotel"
	AccessObjSegment  string = "segment"
	AccessObjCampaign string = "campaign"
	AccessObjCustomer string = "customer"

	// Action
	AccessActionRead   string = "read"
	AccessActionWrite  string = "write"
	AccessActionUpdate string = "update"
	AccessActionDelete string = "delete"

	// Permission
	AccessPermUserRegister AccessPermission = "user:register"
	AccessPermUserInvite   AccessPermission = "user:invite"

	// Role
	AccessRoleSuperAdmin  AccessRole = "super_admin"
	AccessRoleBrandAdmin  AccessRole = "brand_admin"
	AccessRoleBrandViewer AccessRole = "brand_viewer"
	AccessRoleHotelAdmin  AccessRole = "hotel_admin"
	AccessRoleHotelStaff  AccessRole = "hotel_staff"
	AccessRoleHotelViewer AccessRole = "hotel_viewer"
)
