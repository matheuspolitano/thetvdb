package resource

type Relation struct {
	ID       *int    `json:"id"`
	TypeName *string `json:"typeName"`
}

type ParentCompany struct {
	ID       *int     `json:"id"`
	Name     *string  `json:"name"`
	Relation Relation `json:"relation"`
}

type CompanyType struct {
	CompanyTypeID   int    `json:"companyTypeId"`
	CompanyTypeName string `json:"companyTypeName"`
}

type Alias struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}

type Companie struct {
	ID                   int           `json:"id"`
	Name                 string        `json:"name"`
	Slug                 string        `json:"slug"`
	NameTranslations     []string      `json:"nameTranslations"`
	OverviewTranslations []string      `json:"overviewTranslations"`
	Aliases              []Alias       `json:"aliases"`
	Country              string        `json:"country"`
	PrimaryCompanyType   int           `json:"primaryCompanyType"`
	ActiveDate           *string       `json:"activeDate"`
	InactiveDate         *string       `json:"inactiveDate"`
	CompanyType          CompanyType   `json:"companyType"`
	ParentCompany        ParentCompany `json:"parentCompany"`
	TagOptions           *string       `json:"tagOptions"`
}

type ResponseCompanie struct {
	Status string     `json:"status"`
	Data   []Companie `json:"data"`
}
