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
	// Analizar el cuerpo de la solicitud con multipart/form-data
	err := r.ParseMultipartForm(10 << 20) // Límite de 10 MB
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Decodificar el producto desde los datos del formulario
	var product domain.Products
	productJSON := r.FormValue("product") // El producto enviado como JSON
	err = json.Unmarshal([]byte(productJSON), &product)
	if err != nil {
		http.Error(w, "Invalid product format", http.StatusBadRequest)
		return
	}

	// Manejar las imágenes subidas
	files := r.MultipartForm.File["images"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Leer el contenido del archivo como blob
		var buf bytes.Buffer
		_, err = io.Copy(&buf, file)
		if err != nil {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}

		// random
		min := 10
		max := 30

		// Obtener el nombre del archivo del encabezado
		filename := fileHeader.Filename

		// Generar un nombre de archivo único combinando time.Now y un número aleatorio
		if strings.Contains(filename, "undefined") {
			randomNumber := rand.Intn(max-min+1) + min                       // Generar un número aleatorio entre min y max
			cleanedFilename := strings.ReplaceAll(filename, "undefined", "") // Eliminar "undefined"
			filename = fmt.Sprintf("%d%d%s", time.Now().UnixNano(), randomNumber, cleanedFilename)
		} else {
			randomNumber := rand.Intn(max-min+1) + min // Generar un número aleatorio entre min y max
			filename = fmt.Sprintf("%d_%d_%s", time.Now().UnixNano(), randomNumber, filename)
		}

		// Guardar el blob como un archivo en el servidor
		imagePath := fmt.Sprintf("uploads/images/products/%s", filename)
		err = os.WriteFile(imagePath, buf.Bytes(), 0644)
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}

		// Crear la entrada en product_images
		// if i < len(product.Images) {
		// 	product.Images[i].URL = imagePath
		// } else {
		// 	product.Images = append(product.Images, domain.ProductImages{
		// 		URL:      imagePath,
		// 		Featured: "",
		// 	})
		// }
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

	err = r.ParseMultipartForm(10 << 20) // Límite de 10 MB
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	var product domain.Products
	productJSON := r.FormValue("product")
	err = json.Unmarshal([]byte(productJSON), &product)
	if err != nil {
		http.Error(w, "Invalid product format", http.StatusBadRequest)
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

	files := r.MultipartForm.File["images"]

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
	for _, fileHeader := range files {
		if fileHeader.Filename == "" {
			continue
		}

		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// imagePath, err := saveUploadedFile(fileHeader, file)
		// if err != nil {
		// 	http.Error(w, "Failed to save image", http.StatusInternalServerError)
		// 	return
		// }

		// Verificar si la imagen ya existe en la base de datos
		// if i < len(filteredImages) {
		// 	filteredImages[i].URL = imagePath
		// } else {
		// 	product.Images = append(filteredImages, domain.ProductImages{
		// 		URL:      imagePath,
		// 		Featured: "",
		// 	})
		// }
	}
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

func parseExcel(file multipart.File) ([]domain.Products, error) {
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
	requiredHeaders := []string{"Nombre", "Precio", "Costo", "Cantidad"}
	for _, reqHeader := range requiredHeaders {
		if _, exists := headerMap[reqHeader]; !exists {
			return nil, fmt.Errorf("missing required header: %s", reqHeader)
		}
	}

	// Iterar sobre las filas (desde la segunda fila en adelante)
	for _, row := range rows[1:] {
		if len(row) < len(headers) {
			continue // Saltar filas vacías o incompletas
		}

		// Mapear los datos a la estructura Products
		product := domain.Products{
			Name:     row[headerMap["Nombre"]],
			// Price:    parseFloat(row[headerMap["Precio"]]),
			// Cost:     parseFloat(row[headerMap["Costo"]]),
			Quantity: parseFloat(row[headerMap["Cantidad"]]),
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

		// Mapear los datos a la estructura Products
		product := domain.Products{
			Name:     record[headerMap["Nombre"]],
			// Price:    parseFloat(record[headerMap["Precio"]]),
			// Cost:     parseFloat(record[headerMap["Costo"]]),
			Quantity: parseFloat(record[headerMap["Cantidad"]]),
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
		products, err = parseExcel(file)
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
	headers := []string{"Nombre", "Marca", "Precio", "Costo", "Cantidad", "SKU", "Ancho (cm)", "Alto (cm)", "Profundidad (cm)", "Litros (l)", "Peso (kg)"}
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
		cells := []string{
			product.Name,
			product.Brand.Name,
			fmt.Sprintf("%.2f", product.Prices_cf),
			fmt.Sprintf("%.2f", product.Cost),
			fmt.Sprintf("%.2f", product.Quantity),
			*product.SKU.String,
			fmt.Sprintf("%.2f", product.Width),
			fmt.Sprintf("%.2f", product.Height),
			fmt.Sprintf("%.2f", product.Depth),
			fmt.Sprintf("%.2f", product.Liters),
			fmt.Sprintf("%.2f", product.Weight),
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
