import styled from "styled-components";

export const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 24px;
`;

export const Title = styled.h1`
  font-weight: bold;
  font-size: 24px;
  margin-bottom: 16px;
`;

export const Summary = styled.div`
  border: 1px solid #e8e8e8;
  padding: 16px;
  border-radius: 8px;
  width: 50%;
  max-width: 600px;
`;

export const SummaryTitle = styled.h2`
  font-weight: bold;
  font-size: 20px;
  margin-bottom: 16px;
`;

export const SummaryItem = styled.div`
  margin-bottom: 8px;
`;

export const SummaryLabel = styled.span`
  font-weight: bold;
`;

export const SummaryValue = styled.span`
  margin-left: 8px;
`;
