import axios from "axios";

const BASE_URL = "http://ec2-3-71-1-37.eu-central-1.compute.amazonaws.com:8080/api/v1";

export const getProducts = () => {
  console.log("BASE_URL:", BASE_URL);
  return axios.get(`${BASE_URL}/products`);
}

export const createOrder = payload => axios.post(`${BASE_URL}/orders`, payload);