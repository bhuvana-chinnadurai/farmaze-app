"use client"
import { useCallback, useEffect, useMemo, useState } from "react";
import { Table, Pagination, Spin } from "antd";
import { PlusOutlined, MinusOutlined } from '@ant-design/icons';

import withHeader from '@hoc/withHeader';
import { getProducts } from '@api/product.api';
import { createOrder } from '@api/product.api';
import { COLUMNS, PAGE_SIZE } from "./constants";
import CategorySelect from "./components/CategorySelect";
import PreviewOrderModal from './preview'; // Import your PreviewOrderModal component
import ThankYouSummary from './thankyou'
import { makeServer } from '@mirage/server';

if (process.env.NODE_ENV === 'development') {
  makeServer();
}

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

const Home = ({ isLoggedIn, userName }) => {
  const [productList, setProductList] = useState([]);
  const [orderRequest, setOrderRequest] = useState({});
  const [orderSummary, setOrderSummary] = useState({});
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(PAGE_SIZE);
  const [selectedCategory, setSelectedCategory] = useState([]);
  const [searchInput, setSearchInput] = useState("");
  const [isPreviewVisible, setPreviewVisible] = useState(false);
  const [isOrderPlaced, setOrderPlaced] = useState(false);

  useEffect(() => {
    if (!isLoggedIn) return;  // Checks if the user is logged in before fetching data
  
    const fetchProducts = async () => {
      const products = await getProducts();  // Waits for the asynchronous fetch to complete
      if (Array.isArray(products)) {
        setProductList(products);
      } else {
        console.error("Fetched data is not an array:", products);
        setProductList([]); // Fallback to an empty array
      }
    };
  
    fetchProducts();  // Calls the async function to fetch products
  }, [isLoggedIn]);  // Dependence on isLoggedIn ensures this runs only when its value changes
  

  const handlePageChange = (page, CurrentPageSize) => {
    setCurrentPage(page || 1);
    setPageSize(CurrentPageSize);
  };

  const handleQuantityChange = productInfo => (ev) => {
    setOrderRequest({
      ...orderRequest,
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

  function convertOrderRequest(orderRequest, clientId) {
    const products = [];

    // Iterate over each item in the original order
    Object.entries(orderRequest).forEach(([productId, productDetails]) => {
        const quantity = parseInt(productDetails.qty, 10);
        const pricePerUnit = parseFloat(productDetails.price);
        const subtotal = quantity * pricePerUnit;

        // Add product to the products list with new format
        products.push({
            id: productId,
            quantity: quantity
        });

        // Accumulate total price
    });

    // Return the new order format
    return {
        client_id: clientId,
        products: products,
    };
}


  const handlePlaceOrder = async () => {
    try {
      // Assuming 'orderRequest' contains all necessary order data
      console.log("orderRequest",orderRequest)
      const preparedRequest= convertOrderRequest(orderRequest,"42ae47d6-3a90-4652-a28b-6cd9f3a139fc")
      const order = await createOrder(preparedRequest);
      setOrderSummary(order);

      console.log("orderRequest after changing",preparedRequest)

      setOrderPlaced(true); 
      // Set order placed state to true
      // Add additional logic for success feedback or redirection
    } catch (error) {
      console.error('Error placing order:', error);
      // Handle errors, perhaps show an error message to the user
    }
  };
  
  
   // Function to handle preview button click
  const handlePreview = () => {
    setPreviewVisible(true); // Set the preview modal to be visible
  };

  // Function to handle closing of the preview modal
  const handleClosePreview = () => {
    setPreviewVisible(false); // Set the preview modal to be hidden
  };

  const matchesSearch = (product) => {
    const searchLower = searchInput.toLowerCase();
    return !searchInput || product.name.toLowerCase().includes(searchLower) || product.description.toLowerCase().includes(searchLower);
  };

  const matchesCategory = (product) => {
    return !selectedCategory.length || selectedCategory.includes(product.category);
  };
  
  const filteredProductData = useMemo(() => {
    return productList.filter(product => matchesSearch(product) && matchesCategory(product));
  }, [productList, searchInput, selectedCategory]);
  

  const paginatedData = useMemo(() => {
    const fromIdx = (currentPage - 1) * pageSize;
    const toLength = currentPage * pageSize;
    return filteredProductData.slice(fromIdx, toLength);
  }, [filteredProductData, pageSize, currentPage]);

  const updateQuantityWithHandle = (numberToAdd, productInfo) => () => {
    const quantityToUpdate = Number(orderRequest[productInfo.id]?.qty || 0) + numberToAdd;
    setOrderRequest({
      ...orderRequest,
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
            value={orderRequest[productInfo.id]?.qty}
            onChange={handleQuantityChange(productInfo)}
          />
          <PlusOutlined onClick={updateQuantityWithHandle(1, productInfo)} />
        </QuantityContainer>
      ),
    }
  ], [orderRequest]);

  if (!isLoggedIn) return <LoadingContainer><Spin size="large" /></LoadingContainer>;

  return (
    <>
    {!isLoggedIn && <LoadingContainer><Spin size="large" /></LoadingContainer>}
    {isLoggedIn && !isOrderPlaced && (
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
            total={filteredProductData.length}
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
        <FooterButton onClick={handlePreview}>Preview</FooterButton>
        <FooterButton type="primary" onClick={handlePlaceOrder}>Place Order</FooterButton>
      </Footer>
      <PreviewOrderModal
      orderRequest={orderRequest}  // Pass the order info to the modal
      visible={isPreviewVisible}  // Control the visibility of the modal
      onClose={handleClosePreview}  // Pass the handleClosePreview function to close the modal
    />
    </>
    )}
    {isOrderPlaced && (
        <ThankYouSummary userName={userName} orderSummary={orderSummary} />
      )}
    </>
  );
};

export default withHeader(Home);
