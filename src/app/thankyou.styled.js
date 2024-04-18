import styled from 'styled-components';

// Existing styles...
export const Container = styled.div`
  padding: 20px;
  @media print {
    color: black;
    background: white;
    border: none;
    box-shadow: none;
  }
`;

export const Title = styled.h1`
  @media print {
    font-size: 20px;
  }
`;

export const Summary = styled.div`
  margin-bottom: 20px;
  @media print {
    margin-bottom: 10px;
  }
`;

export const SummaryTitle = styled.h2`
  @media print {
    font-size: 16px;
  }
`;

export const SummaryItem = styled.div`
  @media print {
    display: flex;
    justify-content: space-between;
    margin: 5px 0;
  }
`;

export const SummaryLabel = styled.span`
  font-weight: bold;
  @media print {
    font-weight: normal;
  }
`;

export const SummaryValue = styled.span`
  @media print {
    font-weight: normal;
  }
`;