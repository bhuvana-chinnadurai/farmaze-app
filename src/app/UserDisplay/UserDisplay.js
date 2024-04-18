// components/UserDisplay.js
"use client"

import { jwtDecode } from "jwt-decode";
import { useEffect, useState } from 'react';
import { useRouter } from "next/navigation";
import withHeader from "@hoc/withHeader";

const UserDisplay = (username) => {
  const router = useRouter();
    const handleLogout = () => {
      sessionStorage.removeItem('token'); // Clear the token from sessionStorage
      sessionStorage.removeItem('userName'); // Also clear the username if stored
      router.push('/login'); // Redirect to the login page
  };

    return (
      <div>
      {username ? (
          <div>
              Welcome, {username}!   
              <button onClick={handleLogout}>Logout</button>
          </div>
      ) : (
          <p></p>
      )}
  </div>
    );
  
};

export default withHeader(UserDisplay);
