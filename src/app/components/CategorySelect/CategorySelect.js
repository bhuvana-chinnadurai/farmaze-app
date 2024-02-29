import {
  CategorySelectInput
} from "./CategorySelect.styled";

const options = [{
  label: "Fruits",
  value: "fruit"
}, {
  label: "Vegetables",
  value: "vegetable"
}, {
  label: "Oil",
  value: "oil"
}];

const CategorySelect = ({ onChange, selectedCategory }) => {
  return (
    <CategorySelectInput
      mode="multiple"
      allowClear
      placeholder="Please select category"
      options={options}
      maxTagCount="responsive"
      onChange={onChange}
      value={selectedCategory}
    />
  )
};

export default CategorySelect;
