"use client"
import { useState } from "react";
import { useRouter } from "next/navigation";
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

  const handleInputChange = setter => ev => setter(ev.target.value)

  const handleLogin = async () => {
    try {
      const token = await postLogin(userName, password);
      sessionStorage.setItem('token', token); // Consider using HTTP-only cookies instead
      router.push('/'); // Navigate to home on successful login
    } catch (error) {
      alert(error.message); // Show error message from API
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
  )
};

export default withHeader(Login);