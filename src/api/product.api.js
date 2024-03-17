import axios from "axios";

const BASE_URL = process.env.REACT_APP_API_BASE_URL;

export const getProducts = () => {
  console.log("BASE_URL:", BASE_URL);
  return axios.get(`${BASE_URL}/products`);
}

export const createOrder = payload => axios.post(`${BASE_URL}/orders`, payload);