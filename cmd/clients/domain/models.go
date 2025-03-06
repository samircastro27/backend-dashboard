package domain

type BaseModel struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt  string `json:"deleted_at"`
}

type ClientsModel struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Progress  float64 `json:"progress"`
	CreatedAt string `json:"created_at"`
}

type ProjectModel struct {
	ProjectId   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	ProjectSlug string `json:"project_slug"`
}

type EnvironmentModel struct {
	EnvironmentName string `json:"environment_name"`
	EnvironmentId   string `json:"environment_id"`
}

type RuntimeModel struct {
	RuntimeId   string `json:"runtime_id"`
	RuntimeName string `json:"runtime_name"`
}

type BillingInfoModel struct {
	CustomerId     string `json:"billing_customer_id"`
	SubscriptionId string `json:"billing_subscription_id"`
}
