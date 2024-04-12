import React from 'react';
import { Table } from 'antd';
import {
  StyledModal,
  ModalTitle,
  ModalBody,
  ModalRow,
  ModalLabel,
  ModalValue,
  CloseButton,
  StyledButton,
} from './preview.styled';

const PreviewOrderModal = ({ orderRequest, visible, onClose }) => {
  // Define columns for the order table
  const columns = [
    {
      title: 'Product Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Quantity',
      dataIndex: 'quantity',
      key: 'quantity',
    },
  ];

  const orderArray = Object.values(orderRequest).map(product => ({
    id: product.id,
    name: product.name,
    quantity: parseInt(product.qty), // Convert quantity to a number
  }));
  
  return (
    <StyledModal
      title={<ModalTitle>Order Preview</ModalTitle>}
      visible={visible}
      onCancel={onClose}
      footer={[
        <CloseButton key="close" onClick={onClose}>
          Close
        </CloseButton>,
      ]}
    >
      <ModalBody>

        {/* Render order table */}
        <Table
          dataSource={orderArray}
          columns={columns}
          rowKey="id"
        />
      </ModalBody>
    </StyledModal>
  );
};

export default PreviewOrderModal;