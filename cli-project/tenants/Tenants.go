package tenants

type Tenant struct {
	Firstname   string
	Lastname    string
	Gender      string
	HouseNumber int
}

var tenants []Tenant

var tenant1 Tenant = Tenant{
	Firstname:   "Tom",
	Lastname:    "Henry",
	Gender:      "male",
	HouseNumber: 1,
}

func GetTenants() []Tenant {
	tenants = append(tenants, tenant1, tenant1, tenant1, tenant1, tenant1, tenant1)
	return tenants
}

func AddTenant(tenant Tenant) string {
	tenants = append(tenants, tenant)
	return "Tenant " + tenant.Firstname + " " + tenant.Lastname + "added."
}
