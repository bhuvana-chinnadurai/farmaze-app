import axios from "axios";
import getConfig from 'next/config';

const { publicRuntimeConfig } = getConfig();

export const getProducts = async () => {
  const BASE_URL = publicRuntimeConfig.REACT_APP_API_BASE_URL;
  console.log("BASE_URL in getProducts:", BASE_URL); // Log BASE_URL
  
  try {
    const response = await axios.get(`${BASE_URL}/products`);
    return response.data;
  } catch (error) {
    console.error("Error fetching products:", error);
    return null;
  }
}

export const createOrder = async (payload) => {
  const BASE_URL = publicRuntimeConfig.REACT_APP_API_BASE_URL;
  console.log("BASE_URL:", BASE_URL);
  try {
    const response = await axios.post(`${BASE_URL}/orders`, payload);
    return response.data;
  } catch (error) {
    console.error("Error creating order:", error);
    return null;
  }
}