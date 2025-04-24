package routes

import (
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/Jlenin5/facil_backend/pkg/handlers"
	"github.com/gorilla/mux"
)

// account
func AccountRoutes(router *mux.Router, accountUC *usecase.AccountUseCase) {
	handler := handlers.NewAccountHandler(accountUC)
	router.HandleFunc("/account/{id}", handler.GetAccountById).Methods("GET")
	router.HandleFunc("/account/{id}", handler.UpdateAccount).Methods("PUT")
	router.HandleFunc("/change-password/{id}", handler.GetChangePassowrdById).Methods("GET")
	router.HandleFunc("/change-password/{id}", handler.UpdateChangePassowrd).Methods("PUT")
	router.HandleFunc("/plan-billing/{id}", handler.GetPlanBillingById).Methods("GET")
}

// auth
func AuthRoutes(router *mux.Router, authUC *usecase.DauthUseCase) {
	handler := handlers.NewAuthHandler(authUC)
	router.HandleFunc("/sign-in", handler.Login).Methods("POST")
	router.HandleFunc("/sign-in-with-token", handler.ValidateRefreshToken).Methods("GET")
}

// branch offices
func BranchOfficeRoutes(router *mux.Router, branchOfficeUC *usecase.BranchOfficeUseCase) {
	handler := handlers.NewBranchOfficeHandler(branchOfficeUC)
	router.HandleFunc("/branch-offices", handler.GetAllBranchOffices).Methods("GET")
	router.HandleFunc("/branch-offices/{id}", handler.GetBranchOfficeById).Methods("GET")
	router.HandleFunc("/branch-offices", handler.CreateBranchOffice).Methods("POST")
	router.HandleFunc("/branch-offices/{id}", handler.UpdateBranchOffice).Methods("PUT")
	router.HandleFunc("/branch-offices/{id}", handler.DeleteBranchOfficeById).Methods("DELETE")
	router.HandleFunc("/branch-offices", handler.DeleteBranchOfficesByIds).Methods("DELETE")
}

// brands
func BrandRoutes(router *mux.Router, brandUC *usecase.BrandUseCase) {
	handler := handlers.NewBrandHandler(brandUC)
	router.HandleFunc("/brands", handler.GetAllBrands).Methods("GET")
	router.HandleFunc("/brands/{id}", handler.GetBrandById).Methods("GET")
	router.HandleFunc("/brands", handler.CreateBrand).Methods("POST")
	router.HandleFunc("/brands/{id}", handler.UpdateBrand).Methods("PUT")
	router.HandleFunc("/brands/{id}", handler.DeleteBrandById).Methods("DELETE")
	router.HandleFunc("/brands", handler.DeleteBrandsByIds).Methods("DELETE")
}

// cash movements
func CashMovementRoutes(router *mux.Router, cashMovementUC *usecase.CashMovementUseCase) {
	handler := handlers.NewCashMovementHandler(cashMovementUC)
	router.HandleFunc("/cash-movements", handler.GetAllCashMovements).Methods("GET")
	router.HandleFunc("/cash-movements/{id}", handler.GetCashMovementById).Methods("GET")
	router.HandleFunc("/cash-movements", handler.CreateCashMovement).Methods("POST")
	router.HandleFunc("/cash-movements/{id}", handler.UpdateCashMovement).Methods("PUT")
}

// cash registers
func CashRegisterRoutes(router *mux.Router, cashRegisterUC *usecase.CashRegisterUseCase) {
	handler := handlers.NewCashRegisterHandler(cashRegisterUC)
	router.HandleFunc("/cash-registers", handler.GetAllCashRegisters).Methods("GET")
	router.HandleFunc("/cash-registers/{id}", handler.GetCashRegisterById).Methods("GET")
	router.HandleFunc("/cash-registers", handler.CreateCashRegister).Methods("POST")
	router.HandleFunc("/cash-registers/{id}", handler.UpdateCashRegister).Methods("PUT")
	router.HandleFunc("/cash-registers/{id}", handler.DeleteCashRegisterById).Methods("DELETE")
	router.HandleFunc("/cash-registers", handler.DeleteCashRegistersByIds).Methods("DELETE")
}

// categories
func CategoryRoutes(router *mux.Router, categoryUC *usecase.CategoryUseCase) {
	handler := handlers.NewCategoryHandler(categoryUC)
	router.HandleFunc("/categories", handler.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", handler.GetCategoryById).Methods("GET")
	router.HandleFunc("/categories", handler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", handler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", handler.DeleteCategoryById).Methods("DELETE")
	router.HandleFunc("/categories", handler.DeleteCategoriesByIds).Methods("DELETE")
}

// companies
func CompanyRoutes(router *mux.Router, companyUC *usecase.CompanyUseCase) {
	handler := handlers.NewCompanyHandler(companyUC)
	router.HandleFunc("/companies", handler.GetAllCompanies).Methods("GET")
	router.HandleFunc("/companies/{id}", handler.GetCompanyById).Methods("GET")
	router.HandleFunc("/companies", handler.CreateCompany).Methods("POST")
	router.HandleFunc("/companies/{id}", handler.UpdateCompany).Methods("PUT")
	router.HandleFunc("/companies/{id}", handler.DeleteCompanyById).Methods("DELETE")
	router.HandleFunc("/companies", handler.DeleteCompaniesByIds).Methods("DELETE")
}

// currencies
func CurrencyRoutes(router *mux.Router, currencyUC *usecase.CurrencyUseCase) {
	handler := handlers.NewCurrencyHandler(currencyUC)
	router.HandleFunc("/currencies", handler.GetAllCurrencies).Methods("GET")
	router.HandleFunc("/currencies/{id}", handler.GetCurrencyById).Methods("GET")
	router.HandleFunc("/currencies", handler.CreateCurrency).Methods("POST")
	router.HandleFunc("/currencies/{id}", handler.UpdateCurrency).Methods("PUT")
	router.HandleFunc("/currencies/{id}", handler.DeleteCurrencyById).Methods("DELETE")
	router.HandleFunc("/currencies", handler.DeleteCurrenciesByIds).Methods("DELETE")
}

// customers
func CustomerRoutes(router *mux.Router, customerUC *usecase.CustomerUseCase) {
	handler := handlers.NewCustomerHandler(customerUC)
	router.HandleFunc("/customers", handler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handler.GetCustomerById).Methods("GET")
	router.HandleFunc("/customers", handler.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handler.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handler.DeleteCustomerById).Methods("DELETE")
	router.HandleFunc("/customers", handler.DeleteCustomersByIds).Methods("DELETE")
}

// employees
func EmployeeRoutes(router *mux.Router, employeeUC *usecase.EmployeeUseCase) {
	handler := handlers.NewEmployeeHandler(employeeUC)
	router.HandleFunc("/employees", handler.GetAllEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", handler.GetEmployeeById).Methods("GET")
	router.HandleFunc("/employees", handler.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", handler.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", handler.DeleteEmployeeById).Methods("DELETE")
	router.HandleFunc("/employees", handler.DeleteEmployeesByIds).Methods("DELETE")
}

// exchange rates
func ExchangeRateRoutes(router *mux.Router, exchangeRateUC *usecase.ExchangeRateUseCase) {
	handler := handlers.NewExchangeRateHandler(exchangeRateUC)
	router.HandleFunc("/exchange-rates", handler.GetAllExchangeRates).Methods("GET")
	router.HandleFunc("/exchange-rates/{id}", handler.GetExchangeRateById).Methods("GET")
	router.HandleFunc("/exchange-rates", handler.CreateExchangeRate).Methods("POST")
	router.HandleFunc("/exchange-rates/{id}", handler.UpdateExchangeRate).Methods("PUT")
	router.HandleFunc("/exchange-rates/{id}", handler.DeleteExchangeRateById).Methods("DELETE")
	router.HandleFunc("/exchange-rates", handler.DeleteExchangeRatesByIds).Methods("DELETE")
}

// inventory movements
func InventoryMovementRoutes(router *mux.Router, inventoryMovementUC *usecase.InventoryMovementUseCase) {
	handler := handlers.NewInventoryMovementHandler(inventoryMovementUC)
	router.HandleFunc("/inventory-movements", handler.GetAllInventoryMovements).Methods("GET")
	router.HandleFunc("/inventory-movements/{id}", handler.GetInventoryMovementById).Methods("GET")
	router.HandleFunc("/inventory-movements", handler.CreateInventoryMovement).Methods("POST")
	router.HandleFunc("/inventory-movements/{id}", handler.UpdateInventoryMovement).Methods("PUT")
	router.HandleFunc("/inventory-movements/{id}", handler.DeleteInventoryMovementById).Methods("DELETE")
	router.HandleFunc("/inventory-movements", handler.DeleteInventoryMovementsByIds).Methods("DELETE")
}

// job positions
func JobPositionRoutes(router *mux.Router, jobPositionUC *usecase.JobPositionUseCase) {
	handler := handlers.NewJobPositionHandler(jobPositionUC)
	router.HandleFunc("/job-positions", handler.GetAllJobPositions).Methods("GET")
	router.HandleFunc("/job-positions/{id}", handler.GetJobPositionById).Methods("GET")
	router.HandleFunc("/job-positions", handler.CreateJobPosition).Methods("POST")
	router.HandleFunc("/job-positions/{id}", handler.UpdateJobPosition).Methods("PUT")
	router.HandleFunc("/job-positions/{id}", handler.DeleteJobPositionById).Methods("DELETE")
	router.HandleFunc("/job-positions", handler.DeleteJobPositionsByIds).Methods("DELETE")
}

// keys
func KeyRoutes(router *mux.Router, keyUC *usecase.KeyUseCase) {
	handler := handlers.NewKeyHandler(keyUC)
	router.HandleFunc("/keys", handler.GetAllKeys).Methods("GET")
	router.HandleFunc("/keys/{id}", handler.GetKeyById).Methods("GET")
	router.HandleFunc("/keys", handler.CreateKey).Methods("POST")
	router.HandleFunc("/keys/{id}", handler.UpdateKey).Methods("PUT")
}

// opportunity tracking
func OpportunityTrackingRoutes(router *mux.Router, opportunityTrackingUC *usecase.OpportunityTrackingUseCase) {
	handler := handlers.NewOpportunityTrackingHandler(opportunityTrackingUC)
	router.HandleFunc("/opportunity-tracking", handler.GetAllOpportunityTracking).Methods("GET")
	router.HandleFunc("/opportunity-tracking/{id}", handler.GetOpportunityTrackingById).Methods("GET")
	router.HandleFunc("/opportunity-tracking", handler.CreateOpportunityTracking).Methods("POST")
	router.HandleFunc("/opportunity-tracking/{id}", handler.UpdateOpportunityTracking).Methods("PUT")
	router.HandleFunc("/opportunity-tracking/{id}", handler.DeleteOpportunityTrackingById).Methods("DELETE")
	router.HandleFunc("/opportunity-tracking", handler.DeleteOpportunityTrackingByIds).Methods("DELETE")
}

// payment methods
func PaymentMethodRoutes(router *mux.Router, paymentMethodUC *usecase.PaymentMethodUseCase) {
	handler := handlers.NewPaymentMethodHandler(paymentMethodUC)
	router.HandleFunc("/payment-methods", handler.GetAllPaymentMethods).Methods("GET")
	router.HandleFunc("/payment-methods/{id}", handler.GetPaymentMethodById).Methods("GET")
	router.HandleFunc("/payment-methods", handler.CreatePaymentMethod).Methods("POST")
	router.HandleFunc("/payment-methods/{id}", handler.UpdatePaymentMethod).Methods("PUT")
	router.HandleFunc("/payment-methods/{id}", handler.DeletePaymentMethodById).Methods("DELETE")
	router.HandleFunc("/payment-methods", handler.DeletePaymentMethodsByIds).Methods("DELETE")
}

// plans
func PlanRoutes(router *mux.Router, planUC *usecase.PlanUseCase) {
	handler := handlers.NewPlanHandler(planUC)
	router.HandleFunc("/plans", handler.GetAllPlans).Methods("GET")
	router.HandleFunc("/plans/{id}", handler.GetPlanById).Methods("GET")
	router.HandleFunc("/plans", handler.CreatePlan).Methods("POST")
	router.HandleFunc("/plans/{id}", handler.UpdatePlan).Methods("PUT")
	router.HandleFunc("/plans/{id}", handler.DeletePlanById).Methods("DELETE")
	router.HandleFunc("/plans", handler.DeletePlansByIds).Methods("DELETE")
}

// projects
func ProjectRoutes(router *mux.Router, dashboardUC *usecase.ProjectUseCase) {
	handler := handlers.NewProjectHandler(dashboardUC)
	router.HandleFunc("/dashboard-summary", handler.DashboardSummary).Methods("GET")
	router.HandleFunc("/projects", handler.GetProjectData).Methods("GET")
}

// products
func ProductRoutes(router *mux.Router, productUC *usecase.ProductUseCase) {
	handler := handlers.NewProductHandler(productUC)
	router.HandleFunc("/products", handler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handler.GetProductById).Methods("GET")
	router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handler.DeleteProductById).Methods("DELETE")
	router.HandleFunc("/products", handler.DeleteProductsByIds).Methods("DELETE")
}

// purchase orders
func PurchaseOrderRoutes(router *mux.Router, purchaseOrderUC *usecase.PurchaseOrderUseCase) {
	handler := handlers.NewPurchaseOrderHandler(purchaseOrderUC)
	router.HandleFunc("/purchase-orders", handler.GetAllPurchaseOrders).Methods("GET")
	router.HandleFunc("/purchase-orders/{id}", handler.GetPurchaseOrderByID).Methods("GET")
	router.HandleFunc("/purchase-orders", handler.PostPurchaseOrder).Methods("POST")
}

// purchase requests
func PurchaseRequestRoutes(router *mux.Router, purchaseRequestUC *usecase.PurchaseRequestUseCase) {
	handler := handlers.NewPurchaseRequestHandler(purchaseRequestUC)
	router.HandleFunc("/purchase-requests", handler.GetAllPurchaseRequests).Methods("GET")
	router.HandleFunc("/purchase-requests/{id}", handler.GetPurchaseRequestById).Methods("GET")
	router.HandleFunc("/purchase-requests", handler.CreatePurchaseRequest).Methods("POST")
}

// quotes
func QuoteRoutes(router *mux.Router, quoteUC *usecase.QuoteUseCase) {
	handler := handlers.NewQuoteHandler(quoteUC)
	router.HandleFunc("/quotes", handler.GetAllQuotes).Methods("GET")
	router.HandleFunc("/quotes/{id}", handler.GetQuoteById).Methods("GET")
	router.HandleFunc("/quotes", handler.CreateQuote).Methods("POST")
	router.HandleFunc("/quotes/{id}", handler.UpdateQuote).Methods("PUT")
}

// roles
func RoleRoutes(router *mux.Router, roleUC *usecase.RoleUseCase) {
	handler := handlers.NewRoleHandler(roleUC)
	router.HandleFunc("/roles", handler.GetAllRoles).Methods("GET")
	router.HandleFunc("/roles/{id}", handler.GetRoleById).Methods("GET")
	router.HandleFunc("/roles", handler.CreateRole).Methods("POST")
	router.HandleFunc("/roles/{id}", handler.UpdateRole).Methods("PUT")
	router.HandleFunc("/roles/{id}", handler.DeleteRoleById).Methods("DELETE")
	router.HandleFunc("/roles", handler.DeleteRolesByIds).Methods("DELETE")
}

// sale orders
func SaleOrderRoutes(router *mux.Router, saleOrderUC *usecase.SaleOrderUseCase) {
	handler := handlers.NewSaleOrderHandler(saleOrderUC)
	router.HandleFunc("/sale-orders", handler.GetAllSaleOrders).Methods("GET")
	router.HandleFunc("/sale-orders/{id}", handler.GetSaleOrderById).Methods("GET")
	router.HandleFunc("/sale-orders", handler.CreateSaleOrder).Methods("POST")
	router.HandleFunc("/sale-orders/{id}", handler.UpdateSaleOrder).Methods("PUT")
}

// sales
func SaleRoutes(router *mux.Router, saleUC *usecase.SaleUseCase) {
	handler := handlers.NewSaleHandler(saleUC)
	router.HandleFunc("/sales", handler.GetAllSales).Methods("GET")
	router.HandleFunc("/sales/{id}", handler.GetSaleById).Methods("GET")
	router.HandleFunc("/sales", handler.CreateSale).Methods("POST")
	router.HandleFunc("/sales/{id}", handler.UpdateSale).Methods("PUT")
}

// stock control
func StockControlRoutes(router *mux.Router, stockControlUC *usecase.StockControlUseCase) {
	handler := handlers.NewStockControlHandler(stockControlUC)
	router.HandleFunc("/stock-control", handler.GetAllStockControl).Methods("GET")
	router.HandleFunc("/stock-control/{id}", handler.GetStockControlById).Methods("GET")
	router.HandleFunc("/stock-control", handler.CreateStockControl).Methods("POST")
	router.HandleFunc("/stock-control/{id}", handler.UpdateStockControl).Methods("PUT")
	router.HandleFunc("/stock-control/{id}", handler.DeleteStockControlById).Methods("DELETE")
	router.HandleFunc("/stock-control", handler.DeleteStockControlByIds).Methods("DELETE")
}

// subscriptions
func SubscriptionRoutes(router *mux.Router, subscriptionUC *usecase.SubscriptionUseCase) {
	handler := handlers.NewSubscriptionHandler(subscriptionUC)
	router.HandleFunc("/subscriptions", handler.GetAllSubscriptions).Methods("GET")
	router.HandleFunc("/subscriptions/{id}", handler.GetSubscriptionById).Methods("GET")
	router.HandleFunc("/subscriptions", handler.CreateSubscription).Methods("POST")
	router.HandleFunc("/subscriptions/{id}", handler.UpdateSubscription).Methods("PUT")
	router.HandleFunc("/subscriptions/{id}", handler.DeleteSubscriptionById).Methods("DELETE")
	router.HandleFunc("/subscriptions", handler.DeleteSubscriptionsByIds).Methods("DELETE")
}

// suppliers
func SupplierRoutes(router *mux.Router, supplierUC *usecase.SupplierUseCase) {
	handler := handlers.NewSupplierHandler(supplierUC)
	router.HandleFunc("/suppliers", handler.GetAllSuppliers).Methods("GET")
	router.HandleFunc("/suppliers/{id}", handler.GetSupplierById).Methods("GET")
	router.HandleFunc("/suppliers", handler.CreateSupplier).Methods("POST")
	router.HandleFunc("/suppliers/{id}", handler.UpdateSupplier).Methods("PUT")
	router.HandleFunc("/suppliers/{id}", handler.DeleteSupplierById).Methods("DELETE")
	router.HandleFunc("/suppliers", handler.DeleteSuppliersByIds).Methods("DELETE")
}

// systems
func SystemRoutes(router *mux.Router, supplierUC *usecase.SystemUseCase) {
	handler := handlers.NewSystemHandler(supplierUC)
	router.HandleFunc("/systems", handler.GetAllSystems).Methods("GET")
	router.HandleFunc("/systems/{id}", handler.GetSystemById).Methods("GET")
	router.HandleFunc("/systems", handler.CreateSystem).Methods("POST")
	router.HandleFunc("/systems/{id}", handler.UpdateSystem).Methods("PUT")
}

// taxes
func TaxRoutes(router *mux.Router, taxUC *usecase.TaxUseCase) {
	handler := handlers.NewTaxHandler(taxUC)
	router.HandleFunc("/taxes", handler.GetAllTaxes).Methods("GET")
	router.HandleFunc("/taxes/{id}", handler.GetTaxById).Methods("GET")
	router.HandleFunc("/taxes", handler.CreateTax).Methods("POST")
	router.HandleFunc("/taxes/{id}", handler.UpdateTax).Methods("PUT")
	router.HandleFunc("/taxes/{id}", handler.DeleteTaxById).Methods("DELETE")
	router.HandleFunc("/taxes", handler.DeleteTaxesByIds).Methods("DELETE")
}

// users
func UserRoutes(router *mux.Router, userUC *usecase.UserUseCase) {
	handler := handlers.NewUserHandler(userUC)
	router.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handler.GetUserById).Methods("GET")
	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handler.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/users", handler.DeleteUsersByIds).Methods("DELETE")
}

// warehouses
func WarehouseRoutes(router *mux.Router, warehouseUC *usecase.WarehouseUseCase) {
	handler := handlers.NewWarehouseHandler(warehouseUC)
	router.HandleFunc("/warehouses", handler.GetAllWarehouses).Methods("GET")
	router.HandleFunc("/warehouses/{id}", handler.GetWarehouseById).Methods("GET")
	router.HandleFunc("/warehouses", handler.CreateWarehouse).Methods("POST")
	router.HandleFunc("/warehouses/{id}", handler.UpdateWarehouse).Methods("PUT")
	router.HandleFunc("/warehouses/{id}", handler.DeleteWarehouseById).Methods("DELETE")
	router.HandleFunc("/warehouses", handler.DeleteWarehousesByIds).Methods("DELETE")
}

// work areas
func WorkAreaRoutes(router *mux.Router, workAreaUC *usecase.WorkAreaUseCase) {
	handler := handlers.NewWorkAreaHandler(workAreaUC)
	router.HandleFunc("/work-areas", handler.GetAllWorkAreas).Methods("GET")
	router.HandleFunc("/work-areas/{id}", handler.GetWorkAreaById).Methods("GET")
	router.HandleFunc("/work-areas", handler.CreateWorkArea).Methods("POST")
	router.HandleFunc("/work-areas/{id}", handler.UpdateWorkArea).Methods("PUT")
	router.HandleFunc("/work-areas/{id}", handler.DeleteWorkAreaById).Methods("DELETE")
	router.HandleFunc("/work-areas", handler.DeleteWorkAreasByIds).Methods("DELETE")
}