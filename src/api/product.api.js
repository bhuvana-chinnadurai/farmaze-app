import axios from "axios";

const BASE_URL = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/api/v1";
console.log("BASE_URL:", BASE_URL); // Log the BASE_URL


export const getProducts = () => {
  console.log("BASE_URL:", BASE_URL); // Log the BASE_URL
  axios.get(`${BASE_URL}/products`);
}

export const createOrder = payload => axios.post(`${BASE_URL}/orders`, payload);