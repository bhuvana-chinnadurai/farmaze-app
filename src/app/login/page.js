"use client"
import { useState } from "react";
import { useRouter } from "next/navigation";
import withHeader from "@hoc/withHeader";
import { Input } from "antd";

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

  const handleLogin = () => {
    // Validate username and password
    localStorage.setItem("app-token", 1234);
    router.push("/");
  }

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
