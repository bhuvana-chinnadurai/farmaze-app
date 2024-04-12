// src/mirage/server.js
import { createServer, Model } from 'miragejs';
import productData from '../data/products.json';  // Adjust the path as necessary

export function makeServer({ environment = "development" } = {}) {
  return createServer({
    environment,

    models: {
      product: Model,
      order: Model,
    },

    seeds(server) {
      productData.forEach(product => {
        server.create("product", product);
      });
      // Continue adding mock data as needed
    },

    routes() {
      this.namespace = 'api';
      this.get("/products", (schema) => {
        return schema.products.all().models.map(product => {
          return {
            id: product.id,
            name: product.name,
            price: product.price,
            description: product.description,
            available_quantity: product.available_quantity,
            category: product.category,
            unit: product.unit
          };
        });
      });

      const orderData = {
        "id": "95c0ee63-a481-48ea-aaf1-fad6b63aa58f",
        "client_id": "42ae47d6-3a90-4652-a28b-6cd9f3a139fc",
        "products": [
            {
                "product_id": "19f83d06-ae67-4a7c-82e1-f89e84a57f00",
                "quantity": 4,
                "name": "Carrot"
            },
            {
                "product_id": "d85682c3-3cda-408e-9f16-392139cacc33",
                "quantity": 7,
                "name": "Apple",
            }
        ],
        "total_price": 0,
        "created_at": "2024-04-12T11:11:12.794541125Z",
        "status": "ordered"
      };

      this.post("/orders", async (schema, request) => {
        return orderData;
      });
      // Define other API routes
    }
  });
}