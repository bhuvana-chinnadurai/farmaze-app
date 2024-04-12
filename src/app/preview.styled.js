import styled from "styled-components";
import { Button, Input, Modal } from "antd";

export const StyledModal = styled(Modal)`
  .ant-modal-content {
    border-radius: 8px;
  }
`;

export const ModalTitle = styled.h1`
  font-weight: bold;
`;

export const ModalBody = styled.div`
  display: flex;
  flex-direction: column;
  gap: 12px;
`;

export const ModalRow = styled.div`
  display: flex;
  gap: 12px;
`;

export const ModalLabel = styled.label`
  font-weight: bold;
`;

export const ModalValue = styled.span`
  flex-grow: 1;
`;

export const CloseButton = styled(Button)`
  height: 40px;
`;

export const StyledButton = styled(Button)`
  height: 40px;
`;
