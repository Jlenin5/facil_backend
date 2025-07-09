package handlers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/usecase"
	"github.com/Jlenin5/facil_backend/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
)

type ProductHandler struct {
	ProductUC      *usecase.ProductUseCase
	ProductImageUC *usecase.ProductImageUseCase
}

func NewProductHandler(productUC *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{ProductUC: productUC}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	// Decodificar el producto desde los datos del formulario
	var product domain.Products
	err := json.NewDecoder(r.Body).Decode(&product) // Decodificar el cuerpo de la solicitud JSON
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para crear un nuevo producto
	err = h.ProductUC.CreateProduct(&product)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Prooduct created successfully"})
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductUC.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.ProductUC.GetProductById(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product domain.Products
	err = json.NewDecoder(r.Body).Decode(&product) // Decodifica el cuerpo de la solicitud
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}
	product.Id = id

	// existingImages, err := h.ProductUC.GetAllProductImages(product.Id)
	// if err != nil {
	// 	http.Error(w, "Failed to retrieve existing images", http.StatusInternalServerError)
	// 	return
	// }

	// dbImagesMap := make(map[string]domain.ProductImages)
	// for _, img := range existingImages {
	// 	dbImagesMap[img.URL] = img
	// }

	// envImagesMap := make(map[string]domain.ProductImages)
	// for _, img := range product.Images {
	// 	envImagesMap[img.URL] = img
	// }

	// files := r.MultipartForm.File["images"]

	// Eliminar imágenes obsoletas
	// for url, img := range dbImagesMap {
	// 	if _, exists := envImagesMap[url]; !exists {
	// 		if err := h.deleteImage(img.Id, url); err != nil {
	// 			http.Error(w, "Failed to delete image", http.StatusInternalServerError)
	// 			return
	// 		}
	// 	}
	// }

	// Verificar si la URL de la imagen contiene "blob:"
	// filteredImages := filterImagesWithBlob(product.Images)

	// Procesar nuevas imágenes
	// for _, fileHeader := range files {
	// 	if fileHeader.Filename == "" {
	// 		continue
	// 	}

	// 	file, err := fileHeader.Open()
	// 	if err != nil {
	// 		http.Error(w, "Failed to open file", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	defer file.Close()

	// 	imagePath, err := saveUploadedFile(fileHeader, file)
	// 	if err != nil {
	// 		http.Error(w, "Failed to save image", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	Verificar si la imagen ya existe en la base de datos
	// 	if i < len(filteredImages) {
	// 		filteredImages[i].URL = imagePath
	// 	} else {
	// 		product.Images = append(filteredImages, domain.ProductImages{
	// 			URL:      imagePath,
	// 			Featured: "",
	// 		})
	// 	}
	// }
	// product.Images = filteredImages

	err = h.ProductUC.UpdateProduct(&product)
	if err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}

func (h *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Obtén el ID desde los parámetros de la ruta
	if err != nil {
		http.Error(w, "Invalid product Id", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar el producto por ID
	err = h.ProductUC.DeleteProductById(id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

func (h *ProductHandler) DeleteProductsByIds(w http.ResponseWriter, r *http.Request) {
	var ids []int

	// Decodifica el array de IDs del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para eliminar los productos por sus IDs
	err = h.ProductUC.DeleteProductsByIds(ids)
	if err != nil {
		http.Error(w, "Failed to delete products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Products deleted successfully"})
}

func saveUploadedFile(fileHeader *multipart.FileHeader, file multipart.File) (string, error) {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return "", err
	}

	filename := sanitizeFilename(fileHeader.Filename)
	imagePath := fmt.Sprintf("uploads/images/products/%s", filename)

	if err := os.WriteFile(imagePath, buf.Bytes(), 0644); err != nil {
		return "", err
	}

	return imagePath, nil
}

func sanitizeFilename(filename string) string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(21) + 10 // Número aleatorio entre 10 y 30

	if strings.Contains(filename, "undefined") {
		filename = strings.ReplaceAll(filename, "undefined", "")
	}

	return fmt.Sprintf("%d%d%s", time.Now().UnixNano(), randomNumber, filename)
}

func (h *ProductHandler) deleteImage(imageID int, imagePath string) error {
	if err := h.ProductUC.DeleteProductImageById(imageID); err != nil {
		return err
	}
	if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// func filterImagesWithBlob(images []domain.ProductImages) []domain.ProductImages {
// 	var filtered []domain.ProductImages
// 	for _, img := range images {
// 		if strings.HasPrefix(img.URL, "blob:") {
// 			filtered = append(filtered, img)
// 		}
// 	}
// 	return filtered
// }

const (
	PricesCF   domain.PriceType = "prices_cf"
	PricesSF   domain.PriceType = "prices_sf"
	PricesBox  domain.PriceType = "prices_box"
)

// Margen de precios por tipo y nombre
var PriceMargins = map[string]map[string]float64{
	"CF":  {},
	"SF":  {},
	"BOX": {},
}

func init() {
	for i := 1; i <= 20; i++ {
		key := fmt.Sprintf("price_%d", i)
		value := 1.00 + float64(i)/100.0
		PriceMargins["CF"][key] = value
		PriceMargins["SF"][key] = value
		PriceMargins["BOX"][key] = value
	}
}

// Función para calcular precios según tipo y costo
func CalculatePrices(prices []domain.PriceProducts, priceType domain.PriceType, cost float64) []domain.PriceProducts {
	var result []domain.PriceProducts
	var marginType string

	switch priceType {
	case PricesCF:
		marginType = "CF"
	case PricesSF:
		marginType = "SF"
	case PricesBox:
		marginType = "BOX"
	default:
		return result
	}

	for _, p := range prices {
		margin, ok := PriceMargins[marginType][p.Name]
		if !ok {
			// margen por defecto si no se encuentra
			switch priceType {
				case PricesCF:
					margin = 1.01
				case PricesSF:
					margin = 1.01
				case PricesBox:
					margin = 1.01
			}
		}

		var calculatedPrice float64
		if priceType == PricesCF {
			calculatedPrice = (((cost / 1.18) * 1.05 * margin) * 1.18)
		} else {
			calculatedPrice = cost * margin
		}

		result = append(result, domain.PriceProducts{
			Id:    p.Id,
			Name:  p.Name,
			Price: calculatedPrice,
		})
	}

	return result
}

func FindPriceByName(prices []domain.PriceProducts, name string) (domain.PriceProducts, bool) {
	for _, price := range prices {
		if price.Name == name {
			return price, true
		}
	}
	return domain.PriceProducts{}, false
}

func parseExcel(file multipart.File, userIDFloat int) ([]domain.Products, error) {
	var products []domain.Products

	// Leer el archivo Excel
	excelFile, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer excelFile.Close()

	// Obtener la primera hoja
	sheetName := excelFile.GetSheetName(0)
	if sheetName == "" {
		return nil, fmt.Errorf("no sheets found in the Excel file")
	}

	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read rows from sheet: %w", err)
	}

	// Validar que haya datos y leer encabezados
	if len(rows) < 2 {
		return nil, fmt.Errorf("empty file or missing headers")
	}

	headers := rows[0]               // Primera fila (encabezados)
	headerMap := mapHeaders(headers) // Mapear encabezados con índices

	// Validar que los encabezados requeridos estén presentes
	requiredHeaders := []string{"Codigo", "Nombre", "Precio(CF)", "Precio(SF)", "Precio(Caja)", "Costo"}
	for _, reqHeader := range requiredHeaders {
		if _, exists := headerMap[reqHeader]; !exists {
			return nil, fmt.Errorf("missing required header: %s", reqHeader)
		}
	}

	// Iterar sobre las filas (desde la segunda fila en adelante)
	for _, record := range rows[1:] {

		// Simulación de precios originales
		prices := []domain.PriceProducts{
			{Id: 1, Name: "price_1"},
			{Id: 2, Name: "price_2"},
			{Id: 3, Name: "price_3"},
			{Id: 4, Name: "price_4"},
			{Id: 5, Name: "price_5"},
			{Id: 6, Name: "price_6"},
			{Id: 7, Name: "price_7"},
			{Id: 8, Name: "price_8"},
			{Id: 9, Name: "price_9"},
			{Id: 10, Name: "price_10"},
			{Id: 11, Name: "price_11"},
			{Id: 12, Name: "price_12"},
			{Id: 13, Name: "price_13"},
			{Id: 14, Name: "price_14"},
			{Id: 15, Name: "price_15"},
			{Id: 16, Name: "price_16"},
			{Id: 17, Name: "price_17"},
			{Id: 18, Name: "price_18"},
			{Id: 19, Name: "price_19"},
			{Id: 20, Name: "price_20"},
		}

		var cost = parseFloat(record[headerMap["Costo"]])

		// Ejemplo con CF
		calculatedCF := CalculatePrices(prices, PricesCF, cost)
		// Ejemplo con SF
		calculatedSF := CalculatePrices(prices, PricesSF, cost)
		// Ejemplo con BOX
		calculatedBox := CalculatePrices(prices, PricesBox, cost)
		
		var priceCF, priceSF, priceBox float64

		priceNumberCF, err := strconv.Atoi(record[headerMap["Precio(CF)"]])
		if err != nil {
			fmt.Printf("Error convirtiendo Precio(CF) a entero: %v\n", err)
			priceNumberCF = 0 // Valor por defecto
		}
		priceNameCF := fmt.Sprintf("price_%d", priceNumberCF)
		if foundPrice, exists := FindPriceByName(calculatedCF, priceNameCF); exists {
			priceCF = foundPrice.Price
    } else {
			fmt.Printf("Precio %s no encontrado en CF\n", priceNameCF)
			priceCF = 0
    }

		priceNumberSF, err := strconv.Atoi(record[headerMap["Precio(SF)"]])
		if err != nil {
			fmt.Printf("Error convirtiendo Precio(SF) a entero: %v\n", err)
			priceNumberSF = 0 // Valor por defecto
		}
		priceNameSF := fmt.Sprintf("price_%d", priceNumberSF)
    if foundPrice, exists := FindPriceByName(calculatedSF, priceNameSF); exists {
			priceSF = foundPrice.Price
    } else {
			fmt.Printf("Precio %s no encontrado en SF\n", priceNameSF)
			priceSF = 0
    }

		priceNumberBox, err := strconv.Atoi(record[headerMap["Precio(Caja)"]])
		if err != nil {
			fmt.Printf("Error convirtiendo Precio(Caja) a entero: %v\n", err)
			priceNumberBox = 0 // Valor por defecto
		}
		priceNameBox := fmt.Sprintf("price_%d", priceNumberBox)
    if foundPrice, exists := FindPriceByName(calculatedBox, priceNameBox); exists {
			priceBox = foundPrice.Price
    } else {
			fmt.Printf("Precio %s no encontrado en Box\n", priceNameBox)
			priceBox = 0
    }

		cfJson, _ := json.Marshal(calculatedCF)
		sfJson, _ := json.Marshal(calculatedSF)
		boxJson, _ := json.Marshal(calculatedBox)

		if len(record) < len(headers) {
			continue // Saltar filas vacías o incompletas
		}

		// Crear variables temporales para los campos que necesitan punteros
		codigo := record[headerMap["Codigo"]]
    var codigoPtr *string
    if codigo != "" {
			codigoPtr = &codigo
    }

    var precioCFPtr *float64
    if record[headerMap["Precio(CF)"]] != "" {
			precioCFPtr = &priceCF
    }

    var precioSFPtr *float64
    if record[headerMap["Precio(SF)"]] != "" {
			precioSFPtr = &priceSF
    }

    var precioCajaPtr *float64
    if record[headerMap["Precio(Caja)"]] != "" {
			precioCajaPtr = &priceBox
    }

		// Mapear los datos a la estructura Products
		product := domain.Products{
			SKU: domain.NullString{
				String: codigoPtr, 
				Valid:  codigo != "",
			},
			Name: record[headerMap["Nombre"]],
			Prices_cf: cfJson,
			Prices_sf: sfJson,
			Prices_box: boxJson,
			Featured_Pcf: domain.NullFloat{
				Float: precioCFPtr, 
				Valid: record[headerMap["Precio(CF)"]] != "",
			},
			Featured_Psf: domain.NullFloat{
				Float: precioSFPtr, 
				Valid: record[headerMap["Precio(SF)"]] != "",
			},
			Featured_Pbox: domain.NullFloat{
				Float: precioCajaPtr, 
				Valid: record[headerMap["Precio(Caja)"]] != "",
			},
			Cost:       cost,
			Created_By: userIDFloat,
    }

		products = append(products, product)
	}

	return products, nil
}

func parseCSV(file multipart.File) ([]domain.Products, error) {
	var products []domain.Products

	// Leer el CSV
	reader := csv.NewReader(file)
	reader.Comma = ','          // Establecer el delimitador (puede ser ';' o '\t' según el archivo)
	reader.FieldsPerRecord = -1 // Permitir un número variable de campos por fila

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	// Validar que haya datos y leer encabezados
	if len(records) < 2 {
		return nil, fmt.Errorf("empty file or missing headers")
	}

	headers := records[0]
	headerMap := mapHeaders(headers)

	// Validar que los encabezados requeridos estén presentes
	requiredHeaders := []string{"Nombre", "Precio", "Costo", "Cantidad"}
	for _, reqHeader := range requiredHeaders {
		if _, exists := headerMap[reqHeader]; !exists {
			return nil, fmt.Errorf("missing required header: %s", reqHeader)
		}
	}

	// Iterar sobre las filas (desde la segunda fila en adelante)
	for _, record := range records[1:] {
		if len(record) < len(headers) {
			continue // Saltar filas vacías o incompletas
		}

		// Crear variables temporales para los campos que necesitan punteros
		codigo := record[headerMap["Codigo"]]
		var codigoPtr *string
		if codigo != "" {
			codigoPtr = &codigo
		}

		precioCF := parseFloat(record[headerMap["Precio(CF)"]])
		var precioCFPtr *float64
		if record[headerMap["Precio(CF)"]] != "" {
			precioCFPtr = &precioCF
		}

		precioSF := parseFloat(record[headerMap["Precio(SF)"]])
		var precioSFPtr *float64
		if record[headerMap["Precio(SF)"]] != "" {
			precioSFPtr = &precioSF
		}

		precioCaja := parseFloat(record[headerMap["Precio(Caja)"]])
		var precioCajaPtr *float64
		if record[headerMap["Precio(Caja)"]] != "" {
				precioCajaPtr = &precioCaja
		}

		// Mapear los datos a la estructura Products
		product := domain.Products{
			SKU: domain.NullString{
				String: codigoPtr, 
				Valid:  codigo != "",
			},
			Name: record[headerMap["Nombre"]],
			Featured_Pcf: domain.NullFloat{
				Float: precioCFPtr, 
				Valid:   record[headerMap["Precio(CF)"]] != "",
			},
			Featured_Psf: domain.NullFloat{
				Float: precioSFPtr, 
				Valid:   record[headerMap["Precio(SF)"]] != "",
			},
			Featured_Pbox: domain.NullFloat{
				Float: precioCajaPtr, 
				Valid:   record[headerMap["Precio(Caja)"]] != "",
			},
			Cost: parseFloat(record[headerMap["Costo"]]),
		}

		products = append(products, product)
	}

	return products, nil
}

func mapHeaders(headers []string) map[string]int {
	m := make(map[string]int)
	for i, h := range headers {
		m[h] = i
	}
	return m
}

func parseFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return f
}

func (h *ProductHandler) ImportExcel(w http.ResponseWriter, r *http.Request) {
	// Obtener los datos del usuario del contexto
	userData, ok := r.Context().Value(middleware.UserContextKey).(map[string]interface{})
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
	
	// Parsear el formulario con el archivo
	err := r.ParseMultipartForm(10 << 20) // 10 MB de límite
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Obtener el archivo desde la petición
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Determinar el tipo de archivo
	filename := header.Filename
	var products []domain.Products

	switch {
	case strings.HasSuffix(filename, ".xlsx"):
		// Procesar archivo Excel
		products, err = parseExcel(file, int(userIDFloat))
	case strings.HasSuffix(filename, ".csv"):
		// Procesar archivo CSV
		products, err = parseCSV(file)
	default:
		http.Error(w, "Unsupported file format. Only .xlsx and .csv are supported", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Error processing file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Guardar los productos en la base de datos
	err = h.ProductUC.SaveProducts(products)
	if err != nil {
		http.Error(w, "Error saving products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Products imported successfully"})
}

// Exportar datos a Excel
func (h *ProductHandler) ExportExcel(w http.ResponseWriter, r *http.Request) {
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
	var products []domain.Products
	if len(requestBody.Ids) > 0 {
		products, err = h.ProductUC.GetProductsByIds(requestBody.Ids)
	} else {
		products, err = h.ProductUC.GetAllProducts()
	}
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
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
	headers := []string{"Codigo", "Nombre", "Precio(CF)", "Precio(SF)", "Precio(Caja)", "Costo"}
	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		cell.SetString(header)
	}

	// Agregar datos de las ventas
	columnWidths := make([]int, len(headers)) // Para almacenar los anchos máximos por columna

	// Agregar datos de las ventas
	for _, product := range products {
		row := sheet.AddRow()
		
		// Manejar campos posibles nil
		sku := ""
		if product.SKU.String != nil {
			sku = *product.SKU.String
		}
		
		pcf := 0.0
		if product.Featured_Pcf.Float != nil {
			pcf = *product.Featured_Pcf.Float
		}
		
		psf := 0.0
		if product.Featured_Psf.Float != nil {
			psf = *product.Featured_Psf.Float
		}
		
		pbox := 0.0
		if product.Featured_Pbox.Float != nil {
			pbox = *product.Featured_Pbox.Float
		}

		cells := []string{
			sku,
			product.Name,
			fmt.Sprintf("%.2f", pcf),
			fmt.Sprintf("%.2f", psf),
			fmt.Sprintf("%.2f", pbox),
			fmt.Sprintf("%.2f", product.Cost),
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
	w.Header().Set("Content-Disposition", "attachment; filename=products-export.xlsx")
	err = file.Write(w)
	if err != nil {
		http.Error(w, "Failed to write Excel file", http.StatusInternalServerError)
		return
	}
}