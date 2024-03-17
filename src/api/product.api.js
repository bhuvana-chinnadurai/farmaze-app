import axios from "axios";

const BASE_URL = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/api/v1";

export const getProducts = () => {
  return axios.get(`${BASE_URL}/products`, {
    headers: {
      'Content-Type': 'application/json',
      'Origin': 'https://initial-code.d3n8p32hxdybas.amplifyapp.com',
      'Access-Control-Request-Method': 'GET',
      'Access-Control-Request-Headers': 'Authorization, Content-Type',
    },
    withCredentials: true, // Include cookies in the request
  });
};

export const createOrder = (payload) => {
  return axios.post(`${BASE_URL}/orders`, payload, {
    headers: {
      'Content-Type': 'application/json',
      'Origin': 'https://initial-code.d3n8p32hxdybas.amplifyapp.com',
      'Access-Control-Request-Method': 'POST',
      'Access-Control-Request-Headers': 'Authorization, Content-Type',
    },
    withCredentials: true, // Include cookies in the request
  });
};