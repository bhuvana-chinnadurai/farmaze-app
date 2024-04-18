import React from 'react';
import { Container, Title, Summary, SummaryTitle, SummaryItem, SummaryLabel, SummaryValue } from './thankyou.styled';

const ThankYouSummary = ({  orderSummary,username, }) => {
  const {created_at, products = [] } = orderSummary;

  // Format the creation date
  const formattedDate = new Date(created_at).toLocaleString();

  const handlePrint = () => {
    window.print();
  };

  return (
    <Container>
      <Title>Hello {username}, Thank You for Your Order! Your order is placed successfully at {formattedDate}</Title>
      <button onClick={handlePrint} style={{ margin: '10px 0', padding: '8px 16px', cursor: 'pointer' }}>Print Invoice</button>
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
    </Container>
  );
};

export default ThankYouSummary;
