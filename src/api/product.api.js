import axios from "axios";
import { publicRuntimeConfig } from 'next.config.js';


export const getProducts = async () => {
  const BASE_URL = publicRuntimeConfig.REACT_APP_API_BASE_URL
  try {
    const response = await axios.get(`${BASE_URL}/products`);
    console.log("response",response.data);
    return response.data
  } catch (error) {
    console.error("Error fetching products:", error);
    return null;
  }
}

export const createOrder = async (payload) => {
  try {
    const response = await axios.post(`${BASE_URL}/orders`, payload);
    return response.data;
  } catch (error) {
    console.error("Error creating order:", error);
    return null;
  }
}