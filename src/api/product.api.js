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
  const BASE_URL = publicRuntimeConfig.REACT_APP_API_BASE_URL
  // Retrieve the token from sessionStorage or localStorage
  const token = sessionStorage.getItem('token'); // or localStorage.getItem('token');

  try {
    console.log("token ",token)
    const response = await axios.post(`${BASE_URL}/orders`, payload,{ 
      headers: {
      'Authorization': `Bearer ${token}`, // Set the Authorization header
      'Content-Type': 'application/json'
    }});
    return response.data;
  } catch (error) {
    console.error("Error creating order:", error);
    return null;
  }
}