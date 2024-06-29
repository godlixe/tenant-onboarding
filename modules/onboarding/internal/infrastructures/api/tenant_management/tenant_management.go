package tenantmanagement

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
)

const tenantManagementEndpoint = "https://tm.34d.me/api/tenant"

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

type TenantRequestBody struct {
	ProductID      string `json:"product_id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
}

type TenantResponseJSON struct {
	ID             string `json:"tenant_id"`
	ProductID      string `json:"product_id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	Status         string `json:"tenant_status"`
}

type CreateTenantResponse struct {
	Message string
	Data    map[string]any
}

func (r *TenantManagementRepository) CreateTenant(ctx context.Context, tenant *entities.Tenant) (*entities.Tenant, error) {

	tenantReq := TenantRequestBody{
		ProductID:      tenant.ProductID.String(),
		OrganizationID: tenant.OrganizationID.String(),
		Name:           tenant.Name,
	}

	tenantJSON, err := json.Marshal(tenantReq)
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", ctx.Value("token")))

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
		fmt.Println(tenantResponse)
		return nil, errors.New("http request failed")
	}

	var tenantData TenantResponseJSON = TenantResponseJSON{
		ID:             tenantResponse.Data["tenant_id"].(string),
		ProductID:      tenantResponse.Data["product_id"].(string),
		OrganizationID: tenantResponse.Data["organization_id"].(string),
		Status:         tenantResponse.Data["tenant_status"].(string),
		Name:           tenantResponse.Data["name"].(string),
	}

	tenantIDValueObj, err := valueobjects.NewTenantID(tenantData.ID)
	if err != nil {
		return nil, err
	}

	productIDValueObj, err := valueobjects.NewProductID(tenantData.ProductID)
	if err != nil {
		return nil, err
	}

	organizationIDValueObj, err := valueobjects.NewOrganizationID(tenantData.OrganizationID)
	if err != nil {
		return nil, err
	}

	tenantStatus, err := valueobjects.NewTenantStatus(tenantData.Status)
	if err != nil {
		return nil, err
	}

	tenantEntity := entities.Tenant{
		ID:             tenantIDValueObj,
		ProductID:      productIDValueObj,
		OrganizationID: organizationIDValueObj,
		Name:           tenantData.Name,
		Status:         tenantStatus,
	}

	return &tenantEntity, nil
}
