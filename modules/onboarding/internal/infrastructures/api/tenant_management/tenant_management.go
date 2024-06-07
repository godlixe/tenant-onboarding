package tenantmanagement

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
)

const tenantManagementEndpoint = "https://"

type TenantManagementRepository struct {
	client *http.Client
}

func NewTenantManagementRepository(
	client *http.Client,
) *TenantManagementRepository {
	return &TenantManagementRepository{
		client: client,
	}
}

type TenantResponseJSON struct {
	ID             string `json:"id"`
	ProductID      string `json:"product_id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
}

type CreateTenantResponse struct {
	Message string
	Data    TenantResponseJSON
}

func (r *TenantManagementRepository) CreateTenant(ctx context.Context, tenant *entities.Tenant) (*entities.Tenant, error) {
	tenantJSON, err := json.Marshal(tenant)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		tenantManagementEndpoint,
		bytes.NewBuffer(tenantJSON),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tenantResponse CreateTenantResponse
	err = json.Unmarshal(body, &tenantResponse)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("http request failed: " + tenantResponse.Message)
	}

	tenantIDValueObj, err := valueobjects.NewTenantID(tenantResponse.Data.ID)
	if err != nil {
		return nil, err
	}

	productIDValueObj, err := valueobjects.NewProductID(tenantResponse.Data.ProductID)
	if err != nil {
		return nil, err
	}

	organizationIDValueObj, err := valueobjects.NewOrganizationID(tenantResponse.Data.OrganizationID)
	if err != nil {
		return nil, err
	}

	tenantStatus, err := valueobjects.NewTenantStatus(tenantResponse.Data.Status)
	if err != nil {
		return nil, err
	}

	tenantEntity := entities.Tenant{
		ID:             tenantIDValueObj,
		ProductID:      productIDValueObj,
		OrganizationID: organizationIDValueObj,
		Name:           tenantResponse.Data.Name,
		Status:         tenantStatus,
	}

	return &tenantEntity, nil
}
