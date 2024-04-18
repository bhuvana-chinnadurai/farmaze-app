// components/UserDisplay.js
import { jwtDecode } from "jwt-decode";
import { useEffect, useState } from 'react';

const UserDisplay = () => {
  const [username, setUsername] = useState('');

  useEffect(() => {
        const token = sessionStorage.getItem('token'); // Access sessionStorage safely here
        if (token) {
            const decoded = jwtDecode(token);
            setUsername(decoded.username || 'User');
        }
    }, []);

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
          <p>Please log in.</p>
      )}
  </div>
    );
  
};

export default UserDisplay;
