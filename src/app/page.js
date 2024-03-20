"use client"
import { useCallback, useEffect, useMemo, useState } from "react";
import { Table, Pagination, Spin } from "antd";
import { PlusOutlined, MinusOutlined } from '@ant-design/icons';

import withHeader from '@hoc/withHeader';
import { getProducts, createOrder } from '@api/product.api';
import { COLUMNS, PAGE_SIZE } from "./constants";
import CategorySelect from "./components/CategorySelect";

import {
  Container,
  Title,
  FilterContainer,
  SearchInput,
  TableContainer,
  QuantityContainer,
  QuantityInput,
  LoadingContainer,
  Footer,
  FooterButton,
} from "./page.styled";

const Home = ({ isLoggedIn }) => {
  const [productList, setProductList] = useState([]);
  const [orderInfo, setOrderInfo] = useState({});
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(PAGE_SIZE);
  const [selectedCategory, setSelectedCategory] = useState([]);
  const [searchInput, setSearchInput] = useState("");

  useEffect(({}) => {
    if (!isLoggedIn) return;
    getProducts()
      .then(({ data }) => {
        setProductList(data);
      })
      .catch(() => {
        // alert("Error");
      })
  }, [isLoggedIn]);

  const handlePageChange = (page, CurrentPageSize) => {
    setCurrentPage(page || 1);
    setPageSize(CurrentPageSize);
  };

  const handleQuantityChange = productInfo => (ev) => {
    setOrderInfo({
      ...orderInfo,
      [productInfo.id]: {
        ...productInfo,
        qty: ev.target.value
      },
    });
  };

  const handleCategoryChange = (value) => {
    setSelectedCategory(value);
  };

  const handleInputChange = (ev) => {
    setSearchInput(ev.target.value);
  };

  const filterdProductData = useMemo(() => productList.filter((productInfo) => {
    const matchesSearch = !searchInput || `${productInfo.name} ${productInfo.description}`.toLowerCase().includes(searchInput.toLowerCase());
    const matchesCategory = !selectedCategory.length || selectedCategory.includes(productInfo.category);
    return matchesSearch && matchesCategory;
  }), [searchInput, selectedCategory, productList]);

  const paginatedData = useMemo(() => {
    const fromIdx = (currentPage - 1) * pageSize;
    const toLength = currentPage * pageSize;
    return filterdProductData.slice(fromIdx, toLength);
  }, [filterdProductData, pageSize, currentPage]);

  const updateQuantityWithHandle = (numberToAdd, productInfo) => () => {
    const quantityToUpdate = Number(orderInfo[productInfo.id]?.qty || 0) + numberToAdd;
    setOrderInfo({
      ...orderInfo,
      [productInfo.id]: {
        ...productInfo,
        qty: quantityToUpdate < 0 ? 0 : quantityToUpdate,
      },
    });
  };

  const getColumns = useCallback(() => [
    ...COLUMNS,
    {
      title: 'Quantity',
      dataIndex: 'quantity',
      render: (_, productInfo) => (
        <QuantityContainer>
          <MinusOutlined onClick={updateQuantityWithHandle(-1, productInfo)} />
          <QuantityInput
            type="number"
            min={0}
            value={orderInfo[productInfo.id]?.qty}
            onChange={handleQuantityChange(productInfo)}
          />
          <PlusOutlined onClick={updateQuantityWithHandle(1, productInfo)} />
        </QuantityContainer>
      ),
    }
  ], [orderInfo]);

  if (!isLoggedIn) return <LoadingContainer><Spin size="large" /></LoadingContainer>;

  return (
    <>
      <Container>
        <Title>Products</Title>
        <FilterContainer>
          <CategorySelect onChange={handleCategoryChange} selectedCategory={selectedCategory} />
          <SearchInput
            value={searchInput}
            onChange={handleInputChange}
            placeholder="Search Products"
          />
        </FilterContainer>
        <TableContainer>
          <Table
            columns={getColumns()}
            dataSource={paginatedData}
            pagination={false}
            rowKey="id"
          />
          <Pagination
            total={filterdProductData.length}
            onChange={handlePageChange}
            defaultCurrent={1}
            pageSizeOptions={[5, 10, 20, 50]}
            showSizeChanger
            current={currentPage}
            pageSize={pageSize}
          />
        </TableContainer>
      </Container>
      <Footer>
        <FooterButton type="tertiary">Reset</FooterButton>
        <FooterButton>Preview</FooterButton>
        <FooterButton type="primary">Place Order</FooterButton>
      </Footer>
    </>
  )
}

export default withHeader(Home);
