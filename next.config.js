require('dotenv').config();

/** @type {import('next').NextConfig} */
const nextConfig = {
  // Define public runtime configuration (accessible on the client side)
  publicRuntimeConfig: {
    // Define your environment variables here
    REACT_APP_API_BASE_URL: process.env.REACT_APP_API_BASE_URL,
    // Add other environment variables as needed
  },
};

module.exports = nextConfig;