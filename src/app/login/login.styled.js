import styled from "styled-components";
import { Button } from "antd";

export const Container = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: calc(100% - 80px);
  width: 100%;
  overflow: auto;
`;

export const LoginBox = styled.div`
  min-height: 400px;
  width: 360px;
  padding: 28px 20px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
`;

export const Title = styled.h2`
  display: flex;
  justify-content: center;
`;

export const Main = styled.div`
  display: flex;
  flex-direction: column;
  gap: 16px;
  width: 100%;
`;

export const InputContainer = styled.div`
  width: 100%;
  margin-bottom: 10px;
  text-align: left;
`;

export const InputLabel = styled.div`
  font-weight: bold;
  margin-bottom: 5px;
`;

export const LoginButton = styled(Button)`
  align-self: flex-end;
`;
