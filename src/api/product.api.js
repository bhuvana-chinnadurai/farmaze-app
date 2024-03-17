import axios from "axios";
import getConfig from "next/config";

const { publicRuntimeConfig } = getConfig();
const BASE_URL = publicRuntimeConfig.REACT_APP_API_BASE_URL;

export const getProducts = () => {
  console.log("BASE_URL:", BASE_URL);
  return axios.get(`${BASE_URL}/products`);
}

export const createOrder = payload => axios.post(`${BASE_URL}/orders`, payload);