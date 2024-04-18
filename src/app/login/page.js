import { useState } from "react";
import { useRouter } from "next/router"; // Corrected import for useRouter
import withHeader from "@hoc/withHeader";
import { Input } from "antd";
import { postLogin } from '@api/auth.api';

import {
  Container,
  LoginBox,
  Title,
  Main,
  InputContainer,
  InputLabel,
  LoginButton,
} from "./login.styled";

const Login = () => {
  const [userName, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const handleInputChange = setter => ev => setter(ev.target.value);

  const handleLogin = async () => {
    try {
      const { token } = await postLogin(userName, password);
      if (response.ok) {
        console.log("token is received ",token)
        sessionStorage.setItem('token', token);  // Use sessionStorage for more security compared to localStorage

        // Assume the HTTP-only cookie is automatically set by the server on successful login
        router.push('/'); // Navigate to home on successful login
      } else {
        throw new Error('Login failed. Please check your credentials and try again.');
      }
    } catch (error) {
      // It's better to not display raw error messages directly from the API in the UI
      console.error("Login error:", error);
      alert("Login failed. Please check your credentials and try again."); // More generic error message
    }
  };

  return (
    <Container>
      <LoginBox>
        <Title>Login</Title>
        <Main>
          <InputContainer>
            <InputLabel>Email</InputLabel>
            <Input type="email" onChange={handleInputChange(setUserName)} value={userName} />
          </InputContainer>
          <InputContainer>
            <InputLabel>Password</InputLabel>
            <Input type="password" onChange={handleInputChange(setPassword)} value={password} />
          </InputContainer>
        </Main>
        <LoginButton type="primary" onClick={handleLogin}>Login</LoginButton>
      </LoginBox>
    </Container>
  );
};

export default withHeader(Login);
