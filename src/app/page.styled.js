import styled from "styled-components";
import { Button, Input } from "antd";

export const Container = styled.div`
  display: flex;
  padding: 12px;
  flex-direction: column;
  gap: 12px;
  height: calc(100% - 140px);
  overflow: auto;
`;

export const LoadingContainer = styled(Container)`
  justify-content: center;
`;

export const Title = styled.h1`
  font-weight: bold;
`;

export const FilterContainer = styled.div`
  display: flex;
  justify-content: flex-end;
  gap: 12px;
`;

export const SearchInput = styled(Input)`
  width: 200px;
  height: 40px;
`;

export const TableContainer = styled.div`
  display: flex;
  flex-direction: column;
  gap: 24px;
`;

export const QuantityContainer = styled.div`
  display: flex;
  gap: 12px;
  align-items: center;
  cursor: pointer;
`;

export const QuantityInput = styled(Input)`
  width: 48px;
  & .ant-input-number-handler-wrap {
    display: none;
  }
`;

export const Footer = styled.div`
  display: flex;
  height: 60px;
  background-color: #f4f4f4; /* Light gray background color */
  padding: 16px;
  position: static;
  width: 100%;
  gap: 8px;
  justify-content: flex-end;
  align-items: center;
  box-shadow: 0 -4px 8px rgba(0, 0, 0, 0.1);
`;

export const FooterButton = styled(Button)`
  height: 40px;
`;
