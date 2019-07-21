package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"math/rand"
	"net/http"
	"time"
)

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info,omitempty"`
	Price float64 `json:"price"`
}

var products []Product

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type:        productType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						for _, product := range products {
							if int(product.ID) == id {
								return product, nil
							}
						}
					}
					return nil, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get product list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return products, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        productType,
				Description: "Create new product",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					product := Product{
						ID:    int64(rand.Intn(100000)),
						Name:  p.Args["name"].(string),
						Info:  p.Args["info"].(string),
						Price: p.Args["price"].(float64),
					}
					products = append(products, product)
					return products, nil
				},
			},
			"update": &graphql.Field{
				Type:        productType,
				Description: "update product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(int)
					name, nameOk := p.Args["name"].(string)
					info, infoOk := p.Args["info"].(string)
					price, priceOk := p.Args["price"].(float64)
					product := Product{}
					for i, pro := range products {
						if int64(id) == pro.ID {
							if nameOk {
								products[i].Name = name
							}
							if infoOk {
								products[i].Info = info
							}
							if priceOk {
								products[i].Price = price
							}
							product = products[i]
							break
						}
					}
					return product, nil
				},
			},
			"delete": &graphql.Field{
				Type:        productType,
				Description: "Delete product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(int)
					product := Product{}
					for i, pro := range products {
						if int64(id) == pro.ID {
							product = products[i]
							products = append(products[:i], products[i+1:]...)
						}
					}
					return product, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func execQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(
		graphql.Params{
			Schema:        schema,
			RequestString: query,
		},
	)
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %+v", result.Errors)
	}
	return result
}

func initProductsData(p *[]Product) {
	product1 := Product{ID: 1, Name: "Chicha Morada", Info: "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)", Price: 7.99}
	product2 := Product{ID: 2, Name: "Chicha de jora", Info: "Chicha de jora is a corn beer chicha prepared by germinating maize, extracting the malt sugars, boiling the wort, and fermenting it in large vessels (traditionally huge earthenware vats) for several days (wiki)", Price: 5.95}
	product3 := Product{ID: 3, Name: "Pisco", Info: "Pisco is a colorless or yellowish-to-amber colored brandy produced in winemaking regions of Peru and Chile (wiki)", Price: 9.95}

	*p = append(*p, product1, product2, product3)
}

func main() {
	initProductsData(&products)

	http.HandleFunc("/product", func(w http.ResponseWriter, req *http.Request) {
		result := execQuery(req.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("listen on 8080")
	http.ListenAndServe(":8080", nil)
}
