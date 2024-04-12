import React from 'react';
import { Container, Title, Summary, SummaryTitle, SummaryItem, SummaryLabel, SummaryValue } from './thankyou.styled'; // Importing styled-components for styling

const ThankYouSummary = ({ orderSummary}) => {

  console.log("orderSummary",orderSummary)
// Extract client ID from order information
const clientId = orderSummary.client_id;
// Extract status and created at date from order information
const { status, created_at } = orderSummary;
// Extract products array from order information
const products = orderSummary.products || [];

  return (
    <Container>
      <Title>Thank You for Your Order!</Title>
      <Summary>
        <SummaryTitle>Client ID</SummaryTitle>
        <SummaryItem>
          <SummaryValue>{clientId}</SummaryValue>
        </SummaryItem>
      </Summary>
      <Summary>
        <SummaryTitle>Products</SummaryTitle>
        <ul>
          {products.map((product) => (
            <li key={product.product_id}>
              <SummaryItem>
                <SummaryLabel>{product.name}:</SummaryLabel> <SummaryValue>{product.quantity}</SummaryValue>
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