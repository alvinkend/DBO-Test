# DBO-Test
ER Diagram 
![Database ER diagram DBO test](https://github.com/alvinkend/DBO-Test/assets/36125209/1a6b78e6-b3e1-41bd-94dc-c9652bc5e110)

API Documentation
[Uploading {
	"info": {
		"_postman_id": "bd27e7d4-6f5c-4e8f-bfb7-5c1a1fc08df3",
		"name": "DBO TEST",
		"schema": "https://schema.getpostman.com/json/collection/v2.0.0/collection.json",
		"_exporter_id": "17899454"
	},
	"item": [
		{
			"name": "Register customer",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"customer_name\" : \"test lagi\",\n    \"email\" : \"testlagi@gmail.com\",\n    \"shipping_addres\" : \"jakarta\",\n    \"password\" : \"test123\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": "{{HOST}}/v1/register-customer"
			},
			"response": []
		},
		{
			"name": "Find Paged Customer",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{HOST}}/v1/customers?page=1&size=10",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"customers"
					],
					"query": [
						{
							"key": "page",
							"value": "1"
						},
						{
							"key": "size",
							"value": "10"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Find Customer",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{HOST}}/v1/customer/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"customer",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "2"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Update customer",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"customer_name\" : \"test lagi update\",\n    \"email\" : \"testlagi@gmail.com\",\n    \"shipping_addres\" : \"jakarta\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "{{HOST}}/v1/customer/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"customer",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "2"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Delete Customer",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "{{HOST}}/v1/customer/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"customer",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "2"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Login",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"email\" : \"test@gmail.com\",\n    \"password\" : \"test123\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": "{{HOST}}/v1/login"
			},
			"response": []
		},
		{
			"name": "Find Paged Order",
			"request": {
				"method": "GET",
				"header": [
					{
						"key": "Authorization",
						"value": "Bearer {{TOKEN}}",
						"type": "text"
					}
				],
				"url": {
					"raw": "{{HOST}}/v1/orders?page=1&size=10",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"orders"
					],
					"query": [
						{
							"key": "page",
							"value": "1"
						},
						{
							"key": "size",
							"value": "10"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Find Order",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{HOST}}/v1/order/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"order",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "1"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Create Order",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"customer_id\" : 1,\n    \"order_date\" : \"2006-01-02T00:00:00Z\",\n    \"total_amount\" : 1000\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": "{{HOST}}/v1/order"
			},
			"response": []
		},
		{
			"name": "Update Order",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"order_date\" : \"2006-01-02T00:00:00Z\",\n    \"total_amount\" : 2000\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "{{HOST}}/v1/order/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"order",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "6"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Delete Order",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "{{HOST}}/v1/order/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"order",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "6"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Find Paged Order Detail",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{HOST}}/v1/order-details?page=1&size=10",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"order-details"
					],
					"query": [
						{
							"key": "page",
							"value": "1"
						},
						{
							"key": "size",
							"value": "10"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Find Order Detail",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{HOST}}/v1/order-details/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"order-details",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "1"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Create Order detail",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"order_id\" : 1,\n    \"product_id\" : 1,\n    \"qty\" : 2,\n    \"order_price\": 120\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": "{{HOST}}/v1/order-detail"
			},
			"response": []
		},
		{
			"name": "Update Order Detail",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"product_id\" : 1,\n    \"qty\" : 12,\n    \"order_price\": 1210\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "{{HOST}}/v1/order-detail/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"order-detail",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "10"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Delete Order Detail",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "{{HOST}}/v1/order-detail/:id",
					"host": [
						"{{HOST}}"
					],
					"path": [
						"v1",
						"order-detail",
						":id"
					],
					"variable": [
						{
							"key": "id",
							"value": "10"
						}
					]
				}
			},
			"response": []
		}
	]
}DBO TEST.postman_collection.json…]()
