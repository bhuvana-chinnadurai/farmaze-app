import React from 'react';
import { Container, Title, Summary, SummaryTitle, SummaryItem, SummaryLabel, SummaryValue } from './thankyou.styled';

const ThankYouSummary = ({ orderSummary }) => {
  const { client_id, status, created_at, products = [] } = orderSummary;

  const handlePrint = () => {
    window.print();
  };

  return (
    <Container>
      <Title>Thank You for Your Order!</Title>
      <button onClick={handlePrint}>Print Invoice</button>
      <Summary>
        <SummaryTitle>Client ID</SummaryTitle>
        <SummaryItem>
          <SummaryValue>{client_id}</SummaryValue>
        </SummaryItem>
      </Summary>
      <Summary>
        <SummaryTitle>Products</SummaryTitle>
        <ul>
          {products.map((product) => (
            <li key={product.product_id}>
              <SummaryItem>
                <SummaryLabel>{product.name}:</SummaryLabel>
                <SummaryValue>{product.quantity}</SummaryValue>
              </SummaryItem>
            </li>
          ))}
        </ul>
      </Summary>
      <Summary>
        <SummaryTitle>Status</SummaryTitle>
        <SummaryItem>
          <SummaryValue>{status}</SummaryValue>
        </SummaryItem>
      </Summary>
      <Summary>
        <SummaryTitle>Created At</SummaryTitle>
        <SummaryItem>
          <SummaryValue>{new Date(created_at).toLocaleString()}</SummaryValue>
        </SummaryItem>
      </Summary>
    </Container>
  );
};

export default ThankYouSummary;
