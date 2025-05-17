package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/Jlenin5/facil_backend/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

type SaleHandler struct {
	SaleUC *usecase.SaleUseCase
}

func NewSaleHandler(SaleUC *usecase.SaleUseCase) *SaleHandler {
	return &SaleHandler{SaleUC: SaleUC}
}

// Crear una venta
func (h *SaleHandler) CreateSale(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos del usuario del contexto
	userData, ok := r.Context().Value(middleware.UserContextKey).(map[string]interface{})
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verificar que employee exista en userData
	employeeData, exists := userData["employee"]
	if !exists || employeeData == nil {
		http.Error(w, "Employee not found in user data", http.StatusBadRequest)
		return
	}

	// Asegurarse que employee sea un map[string]interface{}
	employeeMap, ok := employeeData.(map[string]interface{})
	if !ok {
		http.Error(w, "Invalid employee data", http.StatusBadRequest)
		return
	}

	// Verificar que warehouse_id exista en el map
	warehouseIDRaw, exists := employeeMap["warehouse_id"]
	if !exists || warehouseIDRaw == nil {
		http.Error(w, "Warehouse ID not found", http.StatusBadRequest)
		return
	}

	// Convertir warehouse_id a float64 y luego a int
	warehouseIDFloat, ok := warehouseIDRaw.(float64)
	if !ok {
		http.Error(w, "Invalid warehouse_id type", http.StatusBadRequest)
		return
	}

	warehouseId := int(warehouseIDFloat)

	// Decodificar el cuerpo de la solicitud (el JSON con los datos de la venta)
	var sale domain.Sales
	err := json.NewDecoder(r.Body).Decode(&sale)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Asignar los datos del usuario a la venta
	sale.Warehouse_Id = warehouseId

	// Validar user_id del usuario autenticado
	userIDRaw, exists := userData["id"]
	if !exists {
		http.Error(w, "User ID not found", http.StatusBadRequest)
		return
	}

	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	sale.User_Id = int(userIDFloat)

	// Reiniciar los IDs de los detalles de la venta (opcional si usas auto increment en la base)
	for i := range sale.SaleDetails {
		sale.SaleDetails[i].Id = 0
	}

	// Crear la venta en la base de datos
	err = h.SaleUC.CreateSale(&sale, sale.SaleDetails)
	if err != nil {
		http.Error(w, "Failed to create sale", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sale created successfully"})
}


// Obtener todas las ventas
func (h *SaleHandler) GetAllSales(w http.ResponseWriter, r *http.Request) {
	sales, err := h.SaleUC.GetAllSales()
	if err != nil {
		http.Error(w, "Failed to fetch sales", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sales)
}

// Obtener una venta por ID
func (h *SaleHandler) GetSaleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid sale Id", http.StatusBadRequest)
		return
	}

	sale, err := h.SaleUC.GetSaleById(id)
	if err != nil {
		http.Error(w, "Sale not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sale)
}

// Actualizar una venta
func (h *SaleHandler) UpdateSale(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid sale ID", http.StatusBadRequest)
		return
	}

	var sale domain.Sales
	err = json.NewDecoder(r.Body).Decode(&sale)
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	sale.Id = id

	err = h.SaleUC.UpdateSale(&sale, sale.SaleDetails)
	if err != nil {
		http.Error(w, "Failed to update sale", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sale updated successfully"})
}

// Visualizar una venta en Pdf por Documento
func (h *SaleHandler) OpenPDF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bill := vars["bill"]

	// Obtener la venta por el número de factura
	sale, err := h.SaleUC.GetSaleByBill(bill)
	if err != nil {
		http.Error(w, "Sale not found", http.StatusNotFound)
		return
	}

	// Crear el PDF
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 100, Ht: 258},
	})

	pdf.AddPage()
	pdf.AddUTF8Font("ArialUnicode", "", "uploads/fonts/Arial_Unicode/Arial-Unicode-Regular.ttf")
	pdf.AddUTF8Font("ArialUnicode", "B", "uploads/fonts/Arial_Unicode/Arial-Unicode-Bold.ttf")

	// Centrar el texto
	centerText := func(spaceLn float64, fontStyle string, Thickness string, text string, fontSize float64) {
		pdf.SetFont(fontStyle, Thickness, fontSize)
		width, _ := pdf.GetPageSize()
		textWidth := pdf.GetStringWidth(text)
		x := (width - textWidth) / 2
		pdf.SetX(x)
		pdf.Cell(40, 10, text)
		pdf.Ln(spaceLn)
	}

	// Título del PDF
	centerText(7, "ArialUnicode", "B", sale.Warehouse.Branch_Office.Company.Name, 16)
	// RUC
	centerText(7, "ArialUnicode", "", "RUC: "+sale.Warehouse.Branch_Office.Company.Ruc, 12)
	// Phone
	centerText(7, "ArialUnicode", "", "Teléfono: +51 "+*sale.Warehouse.Branch_Office.Company.Phone.String, 12)
	// Address
	centerText(7, "ArialUnicode", "", sale.Warehouse.Branch_Office.Company.Address, 12)
	// hr
	centerText(7, "Arial", "", "----------------------------------------------------", 10)
	// Fecha
	centerText(7, "Arial", "", sale.Issue_Date.Format("02/01/2006 03:04:05 PM"), 12)
	// Almacén
	centerText(7, "ArialUnicode", "", sale.Warehouse.Name, 12)
	// Usuario
	centerText(7, "ArialUnicode", "", sale.User.Employee.First_Name+" "+*sale.User.Employee.Second_Name.String+" "+*sale.User.Employee.Surname.String+" "+*sale.User.Employee.Second_Surname.String, 12)
	// Documento
	centerText(7, "ArialUnicode", "B", *sale.Bill.String, 14)
	// hr
	centerText(7, "Arial", "", "----------------------------------------------------", 10)
	// Cliente
	centerText(7, "ArialUnicode", "", "Cliente: "+*sale.Customer.First_Name.String+" "+*sale.Customer.Surname.String, 12)
	// Documento
	centerText(7, "ArialUnicode", "", "Documento: "+sale.Customer.Document_Number, 12)
	// Phone
	centerText(7, "ArialUnicode", "", "Teléfono: +51 "+*sale.Customer.Phone.String, 12)
	// hr
	centerText(4, "Arial", "", "--------------------------------------------------------------------", 10)
	// Título detalles de venta
	centerText(4, "Arial", "", "Cant.               Precio               Desc.               Total", 10)
	// hr
	centerText(4, "Arial", "", "--------------------------------------------------------------------", 10)
	// Productos
	for _, detail := range sale.SaleDetails {
		documentType := ""
		discountValue := "0"
		if detail.Discount.Float != nil && *detail.Discount.Float != 0 {
			if detail.Discount_Method == 0 {
				documentType = "%"
				discountValue = strconv.FormatFloat(*detail.Discount.Float, 'f', 2, 64)
			} else {
				documentType = sale.Currency.Symbol
				discountValue = strconv.Itoa(int(*detail.Discount.Float)) // convierte a entero
			}
		} else {
			// Si no hay descuento o es cero, definir el símbolo correctamente
			if detail.Discount_Method == 0 {
				documentType = "%"
			} else {
				documentType = sale.Currency.Symbol
			}
		}
		centerText(6, "ArialUnicode", "", detail.Product_Name, 10)
		pdf.SetX(5) // Ajusta la posición inicial
		pdf.CellFormat(20, 6, sale.Currency.Symbol + strconv.FormatFloat(detail.Quantity, 'f', 2, 64), "0", 0, "C", false, 0, "")
		pdf.CellFormat(25, 6, sale.Currency.Symbol + strconv.FormatFloat(detail.Price, 'f', 2, 64), "0", 0, "C", false, 0, "")
		pdf.CellFormat(25, 6, documentType + discountValue, "0", 0, "C", false, 0, "")
		pdf.CellFormat(25, 6, sale.Currency.Symbol + strconv.FormatFloat(detail.Total, 'f', 2, 64), "0", 1, "C", false, 0, "")
		pdf.Ln(4)
	}
	// hr
	centerText(7, "Arial", "", "--------------------------------------------------------------------", 10)
	// Subtotal
	pdf.SetLeftMargin(57)
	pdf.CellFormat(25, 6, "SUBTOTAL:     "+sale.Currency.Symbol+strconv.FormatFloat(sale.Subtotal, 'f', 2, 64), "0", 0, "C", false, 0, "")
	pdf.Ln(6)
	pdf.SetLeftMargin(60)
	// IGV
	pdf.CellFormat(20, 6, "IGV(18%):     "+sale.Currency.Symbol+strconv.FormatFloat(sale.Total-sale.Subtotal, 'f', 2, 64), "0", 0, "C", false, 0, "")
	pdf.Ln(3)
	// hr
	centerText(7, "Arial", "", "--------------------------------------------------------------------", 10)
	// Total
	pdf.SetX(53)
	pdf.CellFormat(25, 6, "TOTAL A PAGAR:     "+sale.Currency.Symbol+strconv.FormatFloat(sale.Total, 'f', 2, 64), "0", 0, "C", false, 0, "")
	pdf.Ln(6)
	// Total
	pdf.SetX(52)
	pdf.CellFormat(25, 6, "TOTAL A PAGADO:     "+sale.Currency.Symbol+strconv.FormatFloat(sale.Total_Paid, 'f', 2, 64), "0", 0, "C", false, 0, "")
	pdf.Ln(6)
	// Total
	// Total
	pdf.SetX(59)
	pdf.CellFormat(25, 6, "VUELTO:     "+sale.Currency.Symbol+strconv.FormatFloat(sale.Change, 'f', 2, 64), "0", 0, "C", false, 0, "")
	pdf.Ln(20)
	// Gracias por su compra
	centerText(7, "ArialUnicode", "B", "Gracias por su compra", 12)

	// Guardar el PDF en un buffer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=document-"+*sale.Bill.String+".pdf")
	pdf.Output(w)
}

// Exportar datos a Excel
func (h *SaleHandler) ExportExcel(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud para obtener los IDs
	var requestBody struct {
		Ids []int `json:"ids"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Obtener las ventas según los IDs proporcionados
	var sales []domain.Sales
	if len(requestBody.Ids) > 0 {
		sales, err = h.SaleUC.GetSalesByIds(requestBody.Ids)
	} else {
		sales, err = h.SaleUC.GetAllSales()
	}
	if err != nil {
		http.Error(w, "Failed to fetch sales", http.StatusInternalServerError)
		return
	}

	// Crear un archivo Excel
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Ventas")
	if err != nil {
		http.Error(w, "Failed to create Excel file", http.StatusInternalServerError)
		return
	}

	// Agregar encabezados
	headers := []string{"Recibo", "Almacén", "Cliente", "Moneda", "Usuario", "Método de pago", "Orden de venta", "Fecha de asunto", "Descuento", "Subtotal", "Total", "Estado"}
	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		cell.SetString(header)
	}

	// Agregar datos de las ventas
	columnWidths := make([]int, len(headers)) // Para almacenar los anchos máximos por columna

	// Agregar datos de las ventas
	for _, sale := range sales {
		row := sheet.AddRow()
		
		var customerNameParts []string
		if sale.Customer.First_Name.Valid {
			customerNameParts = append(customerNameParts, *sale.Customer.First_Name.String)
		}
		if sale.Customer.Second_Name.Valid {
			customerNameParts = append(customerNameParts, *sale.Customer.Second_Name.String)
		}
		if sale.Customer.Surname.Valid {
			customerNameParts = append(customerNameParts, *sale.Customer.Surname.String)
		}
		if sale.Customer.Second_Surname.Valid {
			customerNameParts = append(customerNameParts, *sale.Customer.Second_Surname.String)
		}
		if sale.Customer.Company_Name.Valid {
			customerNameParts = append(customerNameParts, *sale.Customer.Company_Name.String)
		}
		customerName := strings.Join(customerNameParts, " ")

		var employeeNameParts []string
		employeeNameParts = append(employeeNameParts, sale.User.Employee.First_Name)
		if sale.User.Employee.Second_Name.Valid {
			employeeNameParts = append(employeeNameParts, *sale.User.Employee.Second_Name.String)
		}
		if sale.User.Employee.Surname.Valid {
			employeeNameParts = append(employeeNameParts, *sale.User.Employee.Surname.String)
		}
		if sale.User.Employee.Second_Surname.Valid {
			employeeNameParts = append(employeeNameParts, *sale.User.Employee.Second_Surname.String)
		}
		employeeName := strings.Join(employeeNameParts, " ")

		cells := []string{
			*sale.Bill.String,
			sale.Warehouse.Name,
			customerName,
			sale.Currency.Symbol + " - " + sale.Currency.Code,
			employeeName,
			sale.Payment_Method.Name,
			sale.Sale_Order.Reference,
			sale.Issue_Date.Format("2006-01-02"),
			fmt.Sprintf("%.2f", *sale.Discount.Float),
			fmt.Sprintf("%.2f", sale.Subtotal),
			fmt.Sprintf("%.2f", sale.Total),
			sale.Sale_Status,
		}

		for colIdx, value := range cells {
			cell := row.AddCell()
			cell.SetString(value)

			// Ajustar el ancho máximo de la columna
			if len(value) > columnWidths[colIdx] {
				columnWidths[colIdx] = len(value)
			}
		}
	}

	// Ajustar el ancho de las columnas
	for colIdx, width := range columnWidths {
		if width < 10 { // Asegurar un ancho mínimo razonable
			width = 10
		}
		sheet.Col(colIdx).Width = float64(width) + 2 // Agregar un pequeño margen
	}

	// Escribir el archivo Excel en la respuesta
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=sales-export.xlsx")
	err = file.Write(w)
	if err != nil {
		http.Error(w, "Failed to write Excel file", http.StatusInternalServerError)
		return
	}
}
