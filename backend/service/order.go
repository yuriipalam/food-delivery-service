package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/model"
	"food_delivery/response"
	"github.com/lib/pq"
	"strconv"
)

type OrderService struct {
	cfg *config.Config
	db  *sql.DB
}

func NewOrderService(cfg *config.Config, db *sql.DB) *OrderService {
	return &OrderService{
		cfg: cfg,
		db:  db,
	}
}

func (os *OrderService) GetAllOrderResponses(customerID int, orders []model.Order) ([]response.OrderResponse, error) {
	orderSuppliersDict, err := os.getOrderSuppliersDictByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	var orderIDs []int
	for orderID := range orderSuppliersDict {
		orderIDs = append(orderIDs, orderID)
	}

	orderProductsDict, err := os.getOrderProductsDictByOrderIDs(orderIDs)
	if err != nil {
		return nil, err
	}

	var orderResponses []response.OrderResponse

	for _, order := range orders {
		var orderResponse response.OrderResponse

		orderMarshaled, err := json.Marshal(order)
		if err != nil {
			return nil, fmt.Errorf("cannot marshal order")
		}

		err = json.Unmarshal(orderMarshaled, &orderResponse)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal order")
		}

		orderResponse.Suppliers = orderSuppliersDict[order.ID]
		orderResponse.Products = orderProductsDict[order.ID]

		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}

func (os *OrderService) getOrderSuppliersDictByCustomerID(customerID int) (map[int][]response.OrderSupplierResponse, error) {
	query := `SELECT os.order_id, s.id, s.name, s.image FROM order_supplier os
			  JOIN supplier s ON os.supplier_id = s.id
		      JOIN "order" o ON o.id = os.order_id
			  WHERE o.customer_id = $1;`

	stmt, err := os.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare query")
	}

	rows, err := stmt.Query(customerID)
	if err != nil {
		return nil, fmt.Errorf("cannot run query")
	}

	orderSuppliersDict := make(map[int][]response.OrderSupplierResponse)

	for rows.Next() {
		var orderID int
		var image string

		var suppResp response.OrderSupplierResponse

		err := rows.Scan(
			&orderID,
			&suppResp.SupplierID,
			&suppResp.SupplierName,
			&image,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan supplier response")
		}

		suppResp.SupplierImageURL = fmt.Sprintf("%s/images/suppliers/%s/%s", config.Root, strconv.Itoa(suppResp.SupplierID), image)

		if orderSuppliersDict[orderID] == nil {
			orderSuppliersDict[orderID] = []response.OrderSupplierResponse{suppResp}
		} else {
			orderSuppliersDict[orderID] = append(orderSuppliersDict[orderID], suppResp)
		}
	}

	return orderSuppliersDict, nil
}

func (os *OrderService) getOrderProductsDictByOrderIDs(orderIDs []int) (map[int][]response.OrderProductResponse, error) {
	query := `SELECT op.order_id, p.id, p.supplier_id, p.name, op.product_quantity FROM order_product op
			  JOIN product p ON op.product_id = p.id
			  WHERE op.order_id = ANY($1)`

	stmt, err := os.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement")
	}

	rows, err := stmt.Query(pq.Array(orderIDs))
	if err != nil {
		return nil, fmt.Errorf("cannot run query")
	}

	var orderProductsDict = make(map[int][]response.OrderProductResponse)

	for rows.Next() {
		var orderID int
		var orderProductResp response.OrderProductResponse

		err := rows.Scan(
			&orderID,
			&orderProductResp.ProductID,
			&orderProductResp.ProductSupplierID,
			&orderProductResp.ProductName,
			&orderProductResp.ProductQuantity,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan rows")
		}

		if orderProductsDict[orderID] == nil {
			orderProductsDict[orderID] = []response.OrderProductResponse{orderProductResp}
		} else {
			orderProductsDict[orderID] = append(orderProductsDict[orderID], orderProductResp)
		}
	}

	return orderProductsDict, nil
}
