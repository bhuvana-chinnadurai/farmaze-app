/** @type {import('next').NextConfig} */
const nextConfig = {
  // Define public runtime configuration (accessible on the client side)
  publicRuntimeConfig: {
    // Define your environment variables here
    REACT_APP_API_BASE_URL: process.env.REACT_APP_API_BASE_URL,
    // Add other environment variables as needed
  },
  // Define server runtime configuration (accessible on the server side)
  serverRuntimeConfig: {
    // Define your server-side environment variables here
    // For example:
    // SERVER_SIDE_ENV_VARIABLE: process.env.SERVER_SIDE_ENV_VARIABLE,
    // Add other server-side environment variables as needed
  },
};

module.exports = nextConfig;