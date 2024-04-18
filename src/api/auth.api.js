import axios from "axios";
import { publicRuntimeConfig } from 'next.config.js';


export const postLogin = async (username, password) => {
  const BASE_URL = publicRuntimeConfig.REACT_APP_API_BASE_URL
  try {
    const response = await axios.post(`${BASE_URL}/auth/login`, {
      username,
      password,
    });
    return response.data; // Assuming the API returns token data on successful login
  } catch (error) {
    console.error("Error during login:", error.response || error);
    // Handle different types of errors appropriately
    if (error.response && error.response.status === 401) {
      throw new Error("Invalid credentials");
    } else {
      throw new Error("Login failed due to server error");
    }
  }
}
