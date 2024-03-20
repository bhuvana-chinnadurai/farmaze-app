import axios from "axios";

export const getServerSideProps = async () => {
  const BASE_URL = process.env.REACT_APP_API_BASE_URL;
  
  // Fetch data using the BASE_URL if needed
  
  return {
    props: {
      BASE_URL
    }
  };
};

export const getProducts = async () => {
  const BASE_URL = process.env.REACT_APP_API_BASE_URL;
  console.log("BASE_URL:", BASE_URL);
  try {
    const response = await axios.get(`${BASE_URL}/products`);
    return response.data;
  } catch (error) {
    console.error("Error fetching products:", error);
    return null;
  }
}

export const createOrder = async (payload) => {
  const BASE_URL = process.env.REACT_APP_API_BASE_URL;
  console.log("BASE_URL:", BASE_URL);
  try {
    const response = await axios.post(`${BASE_URL}/orders`, payload);
    return response.data;
  } catch (error) {
    console.error("Error creating order:", error);
    return null;
  }
}